// Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bastion

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	operationsv1alpha1 "github.com/gardener/gardener/pkg/apis/operations/v1alpha1"
	kutil "github.com/gardener/gardener/pkg/utils/kubernetes"
)

func (c *Controller) bastionAdd(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		c.log.Error(err, "Couldn't get key for object", "object", obj)
		return
	}
	c.bastionQueue.Add(key)
}

func (c *Controller) bastionUpdate(_, newObj interface{}) {
	c.bastionAdd(newObj)
}

func (c *Controller) bastionDelete(obj interface{}) {
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		c.log.Error(err, "Couldn't get key for object", "object", obj)
		return
	}
	c.bastionQueue.Add(key)
}

func (c *Controller) shootAdd(ctx context.Context, obj interface{}) {
	shoot, ok := obj.(*gardencorev1beta1.Shoot)
	if !ok {
		return
	}

	// only shoot deletions should trigger this, so we can cleanup Bastions
	if shoot.DeletionTimestamp == nil {
		return
	}

	// list all bastions that reference this shoot
	// TODO: this should be done via a field-selector
	bastionList := operationsv1alpha1.BastionList{}
	listOptions := client.ListOptions{Namespace: shoot.Namespace, Limit: 1}

	if err := c.gardenClient.List(ctx, &bastionList, &listOptions); err != nil {
		c.log.Error(err, "Failed to list Bastions")
		return
	}

	for _, bastion := range bastionList.Items {
		if bastion.Spec.ShootRef.Name == shoot.Name {
			c.bastionAdd(bastion)
		}
	}
}

func (c *Controller) shootUpdate(ctx context.Context, _, newObj interface{}) {
	newShoot := newObj.(*gardencorev1beta1.Shoot)

	if newShoot.Status.ObservedGeneration != newShoot.Generation {
		c.shootAdd(ctx, newObj)
	}
}

func (c *Controller) shootDelete(ctx context.Context, obj interface{}) {
	c.shootAdd(ctx, obj)
}

// NewBastionReconciler creates a new instance of a reconciler which reconciles Bastions.
func NewBastionReconciler(gardenClient client.Client, maxLifetime time.Duration) reconcile.Reconciler {
	return &reconciler{
		gardenClient: gardenClient,
		maxLifetime:  maxLifetime,
	}
}

type reconciler struct {
	gardenClient client.Client
	maxLifetime  time.Duration
}

// Reconcile reacts to updates on Bastion resources and cleans up expired Bastions.
func (r *reconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	log := logf.FromContext(ctx)

	bastion := &operationsv1alpha1.Bastion{}
	if err := r.gardenClient.Get(ctx, request.NamespacedName, bastion); err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("Object is gone, stop reconciling")
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, fmt.Errorf("error retrieving object from store: %w", err)
	}

	// do not reconcile anymore once the object is marked for deletion
	if bastion.DeletionTimestamp != nil {
		return reconcile.Result{}, nil
	}

	log = log.WithValues("shootName", bastion.Spec.ShootRef.Name)

	// fetch associated Shoot
	shoot := gardencorev1beta1.Shoot{}
	shootKey := kutil.Key(bastion.Namespace, bastion.Spec.ShootRef.Name)
	if err := r.gardenClient.Get(ctx, shootKey, &shoot); err != nil {
		// This should never happen, as the shoot deletion is stopped unless all Bastions
		// are removed. This is required because once a Shoot is gone, the Cluster resource
		// is gone as well and without that, cleanly destroying a Bastion is not possible.
		if apierrors.IsNotFound(err) {
			log.Info("Deleting bastion because target shoot is gone")
			return reconcile.Result{}, client.IgnoreNotFound(r.gardenClient.Delete(ctx, bastion))
		}

		return reconcile.Result{}, fmt.Errorf("could not get shoot %v: %w", shootKey, err)
	}

	// delete the bastion if the shoot is marked for deletion
	if shoot.DeletionTimestamp != nil {
		log.Info("Deleting bastion because target shoot is in deletion")
		return reconcile.Result{}, client.IgnoreNotFound(r.gardenClient.Delete(ctx, bastion))
	}

	// the Shoot for this bastion has been migrated to another Seed, we have to garbage-collect
	// the old bastion (bastions are not migrated, users are required to create new bastions);
	// equality is the correct check here, as the admission plugin already prevents Bastions
	// from existing without a spec.SeedName being set. So it cannot happen that we accidentally
	// delete a Bastion without seed (i.e. an unreconciled, new Bastion);
	// under normal operations, shoots cannot be migrated to another seed while there are still
	// bastions for it, so this check here is just a safety measure.
	if !equality.Semantic.DeepEqual(shoot.Spec.SeedName, bastion.Spec.SeedName) {
		log.Info("Deleting bastion because the referenced Shoot has been migrated to another Seed", "newSeed", shoot.Spec.SeedName)
		return reconcile.Result{}, client.IgnoreNotFound(r.gardenClient.Delete(ctx, bastion))
	}

	now := time.Now()

	// delete the bastion once it has expired
	if bastion.Status.ExpirationTimestamp != nil && now.After(bastion.Status.ExpirationTimestamp.Time) {
		log.Info("Deleting expired bastion", "expirationTimestamp", bastion.Status.ExpirationTimestamp.Time)
		return reconcile.Result{}, client.IgnoreNotFound(r.gardenClient.Delete(ctx, bastion))
	}

	// delete the bastion once it has reached its maximum lifetime
	if time.Since(bastion.CreationTimestamp.Time) > r.maxLifetime {
		log.Info("Deleting bastion because it reached its maximum lifetime", "creationTimestamp", bastion.CreationTimestamp.Time, "maxLifetime", r.maxLifetime)
		return reconcile.Result{}, client.IgnoreNotFound(r.gardenClient.Delete(ctx, bastion))
	}

	// requeue when the Bastion expires or reaches its lifetime, whichever is sooner
	requeueAfter := time.Until(bastion.CreationTimestamp.Time.Add(r.maxLifetime))
	if bastion.Status.ExpirationTimestamp != nil {
		expiresIn := time.Until(bastion.Status.ExpirationTimestamp.Time)
		if expiresIn < requeueAfter {
			requeueAfter = expiresIn
		}
	}

	return reconcile.Result{
		RequeueAfter: requeueAfter,
	}, nil
}
