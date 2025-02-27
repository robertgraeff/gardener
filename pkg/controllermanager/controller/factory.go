// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	gardencore "github.com/gardener/gardener/pkg/apis/core"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	"github.com/gardener/gardener/pkg/apis/seedmanagement"
	seedmanagementv1alpha1 "github.com/gardener/gardener/pkg/apis/seedmanagement/v1alpha1"
	"github.com/gardener/gardener/pkg/client/kubernetes/clientmap"
	"github.com/gardener/gardener/pkg/client/kubernetes/clientmap/keys"
	"github.com/gardener/gardener/pkg/controllermanager/apis/config"
	bastioncontroller "github.com/gardener/gardener/pkg/controllermanager/controller/bastion"
	csrcontroller "github.com/gardener/gardener/pkg/controllermanager/controller/certificatesigningrequest"
	cloudprofilecontroller "github.com/gardener/gardener/pkg/controllermanager/controller/cloudprofile"
	controllerdeploymentcontroller "github.com/gardener/gardener/pkg/controllermanager/controller/controllerdeployment"
	controllerregistrationcontroller "github.com/gardener/gardener/pkg/controllermanager/controller/controllerregistration"
	eventcontroller "github.com/gardener/gardener/pkg/controllermanager/controller/event"
	exposureclasscontroller "github.com/gardener/gardener/pkg/controllermanager/controller/exposureclass"
	managedseedsetcontroller "github.com/gardener/gardener/pkg/controllermanager/controller/managedseedset"
	plantcontroller "github.com/gardener/gardener/pkg/controllermanager/controller/plant"
	projectcontroller "github.com/gardener/gardener/pkg/controllermanager/controller/project"
	quotacontroller "github.com/gardener/gardener/pkg/controllermanager/controller/quota"
	secretbindingcontroller "github.com/gardener/gardener/pkg/controllermanager/controller/secretbinding"
	seedcontroller "github.com/gardener/gardener/pkg/controllermanager/controller/seed"
	shootcontroller "github.com/gardener/gardener/pkg/controllermanager/controller/shoot"
	"github.com/gardener/gardener/pkg/logger"
	"github.com/gardener/gardener/pkg/operation/garden"
)

// GardenControllerFactory contains information relevant to controllers for the Garden API group.
type GardenControllerFactory struct {
	cfg       *config.ControllerManagerConfiguration
	clientMap clientmap.ClientMap
	recorder  record.EventRecorder
}

// NewGardenControllerFactory creates a new factory for controllers for the Garden API group.
func NewGardenControllerFactory(clientMap clientmap.ClientMap, cfg *config.ControllerManagerConfiguration, recorder record.EventRecorder) *GardenControllerFactory {
	return &GardenControllerFactory{
		cfg:       cfg,
		clientMap: clientMap,
		recorder:  recorder,
	}
}

// Run starts all the controllers for the Garden API group. It also performs bootstrapping tasks.
func (f *GardenControllerFactory) Run(ctx context.Context) error {
	log := logf.Log.WithName("controller")

	gardenClientSet, err := f.clientMap.GetClient(ctx, keys.ForGarden())
	if err != nil {
		return fmt.Errorf("failed to get garden client: %+v", err)
	}

	if err := addAllFieldIndexes(ctx, gardenClientSet.Cache()); err != nil {
		return err
	}

	if err := f.clientMap.Start(ctx.Done()); err != nil {
		return fmt.Errorf("failed to start ClientMap: %+v", err)
	}

	runtime.Must(garden.BootstrapCluster(ctx, gardenClientSet))
	log.Info("Successfully bootstrapped Garden cluster")

	// Create controllers.
	bastionController, err := bastioncontroller.NewBastionController(ctx, log, f.clientMap, f.cfg.Controllers.Bastion.MaxLifetime.Duration)
	if err != nil {
		return fmt.Errorf("failed initializing Bastion controller: %w", err)
	}

	cloudProfileController, err := cloudprofilecontroller.NewCloudProfileController(ctx, log, f.clientMap, f.recorder)
	if err != nil {
		return fmt.Errorf("failed initializing CloudProfile controller: %w", err)
	}

	controllerRegistrationController, err := controllerregistrationcontroller.NewController(ctx, f.clientMap)
	if err != nil {
		return fmt.Errorf("failed initializing ControllerRegistration controller: %w", err)
	}

	csrController, err := csrcontroller.NewCSRController(ctx, log, f.clientMap)
	if err != nil {
		return fmt.Errorf("failed initializing CSR controller: %w", err)
	}

	exposureClassController, err := exposureclasscontroller.NewExposureClassController(ctx, f.clientMap, f.recorder)
	if err != nil {
		return fmt.Errorf("failed initializing ExposureClass controller: %w", err)
	}

	plantController, err := plantcontroller.NewController(ctx, f.clientMap, f.cfg)
	if err != nil {
		return fmt.Errorf("failed initializing Plant controller: %w", err)
	}

	projectController, err := projectcontroller.NewProjectController(ctx, f.clientMap, f.cfg, f.recorder)
	if err != nil {
		return fmt.Errorf("failed initializing Project controller: %w", err)
	}

	quotaController, err := quotacontroller.NewQuotaController(ctx, f.clientMap, f.recorder)
	if err != nil {
		return fmt.Errorf("failed initializing Quota controller: %w", err)
	}

	secretBindingController, err := secretbindingcontroller.NewSecretBindingController(ctx, f.clientMap, f.recorder)
	if err != nil {
		return fmt.Errorf("failed initializing SecretBinding controller: %w", err)
	}

	seedController, err := seedcontroller.NewSeedController(ctx, f.clientMap, f.cfg)
	if err != nil {
		return fmt.Errorf("failed initializing Seed controller: %w", err)
	}

	controllerDeploymentController, err := controllerdeploymentcontroller.New(ctx, f.clientMap, logger.Logger)
	if err != nil {
		return fmt.Errorf("failed initializing ControllerDeployment controller: %w", err)
	}

	shootController, err := shootcontroller.NewShootController(ctx, f.clientMap, f.cfg, f.recorder)
	if err != nil {
		return fmt.Errorf("failed initializing Shoot controller: %w", err)
	}

	managedSeedSetController, err := managedseedsetcontroller.NewManagedSeedSetController(ctx, f.clientMap, f.cfg, f.recorder, logger.Logger)
	if err != nil {
		return fmt.Errorf("failed initializing ManagedSeedSet controller: %w", err)
	}

	go bastionController.Run(ctx, f.cfg.Controllers.Bastion.ConcurrentSyncs)
	go cloudProfileController.Run(ctx, f.cfg.Controllers.CloudProfile.ConcurrentSyncs)
	go controllerDeploymentController.Run(ctx, f.cfg.Controllers.ControllerDeployment.ConcurrentSyncs)
	go controllerRegistrationController.Run(ctx, f.cfg.Controllers.ControllerRegistration.ConcurrentSyncs)
	go csrController.Run(ctx, 1)
	go plantController.Run(ctx, f.cfg.Controllers.Plant.ConcurrentSyncs)
	go projectController.Run(ctx, f.cfg.Controllers.Project.ConcurrentSyncs)
	go quotaController.Run(ctx, f.cfg.Controllers.Quota.ConcurrentSyncs)
	go secretBindingController.Run(ctx, f.cfg.Controllers.SecretBinding.ConcurrentSyncs, f.cfg.Controllers.SecretBindingProvider.ConcurrentSyncs)
	go seedController.Run(ctx, f.cfg.Controllers.Seed.ConcurrentSyncs)
	go shootController.Run(ctx, f.cfg.Controllers.ShootMaintenance.ConcurrentSyncs, f.cfg.Controllers.ShootQuota.ConcurrentSyncs, f.cfg.Controllers.ShootHibernation.ConcurrentSyncs, f.cfg.Controllers.ShootReference.ConcurrentSyncs, f.cfg.Controllers.ShootRetry.ConcurrentSyncs, f.cfg.Controllers.ShootConditions.ConcurrentSyncs, f.cfg.Controllers.ShootStatusLabel.ConcurrentSyncs)
	go exposureClassController.Run(ctx, f.cfg.Controllers.ExposureClass.ConcurrentSyncs)
	go managedSeedSetController.Run(ctx, f.cfg.Controllers.ManagedSeedSet.ConcurrentSyncs)

	if eventControllerConfig := f.cfg.Controllers.Event; eventControllerConfig != nil {
		eventController, err := eventcontroller.NewController(ctx, f.clientMap, eventControllerConfig)
		if err != nil {
			return fmt.Errorf("failed initializing Event controller: %w", err)
		}

		go eventController.Run(ctx)
	}

	log.Info("gardener-controller-manager initialized")

	// Shutdown handling
	<-ctx.Done()

	log.Info("I have received a stop signal and will no longer watch resources")
	log.Info("Bye Bye!")

	return nil
}

// addAllFieldIndexes adds all field indexes used by gardener-controller-manager to the given FieldIndexer (i.e. cache).
// field indexes have to be added before the cache is started (i.e. before the clientmap is started)
func addAllFieldIndexes(ctx context.Context, indexer client.FieldIndexer) error {
	if err := indexer.IndexField(ctx, &gardencorev1beta1.Project{}, gardencore.ProjectNamespace, func(obj client.Object) []string {
		project, ok := obj.(*gardencorev1beta1.Project)
		if !ok {
			return []string{""}
		}
		if project.Spec.Namespace == nil {
			return []string{""}
		}
		return []string{*project.Spec.Namespace}
	}); err != nil {
		return fmt.Errorf("failed to add indexer to Project Informer: %w", err)
	}

	if err := indexer.IndexField(ctx, &gardencorev1beta1.Shoot{}, gardencore.ShootSeedName, func(obj client.Object) []string {
		shoot, ok := obj.(*gardencorev1beta1.Shoot)
		if !ok {
			return []string{""}
		}
		if shoot.Spec.SeedName == nil {
			return []string{""}
		}
		return []string{*shoot.Spec.SeedName}
	}); err != nil {
		return fmt.Errorf("failed to add indexer to Shoot Informer: %w", err)
	}

	if err := indexer.IndexField(ctx, &seedmanagementv1alpha1.ManagedSeed{}, seedmanagement.ManagedSeedShootName, func(obj client.Object) []string {
		ms, ok := obj.(*seedmanagementv1alpha1.ManagedSeed)
		if !ok {
			return []string{""}
		}
		if ms.Spec.Shoot == nil {
			return []string{""}
		}
		return []string{ms.Spec.Shoot.Name}
	}); err != nil {
		return fmt.Errorf("failed to add indexer to ManagedSeed Informer: %w", err)
	}

	return nil
}
