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

package botanist_test

import (
	"context"
	"fmt"

	"github.com/gardener/gardener/pkg/operation"
	. "github.com/gardener/gardener/pkg/operation/botanist"
	mockdwdaccess "github.com/gardener/gardener/pkg/operation/botanist/component/dependencywatchdog/mock"
	shootpkg "github.com/gardener/gardener/pkg/operation/shoot"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
)

var _ = Describe("DependencyWatchdogAccess", func() {
	var (
		ctrl     *gomock.Controller
		botanist *Botanist
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		botanist = &Botanist{Operation: &operation.Operation{}}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("#DeployDependencyWatchdogAccess", func() {
		var (
			dwdAccess *mockdwdaccess.MockAccessInterface

			ctx      = context.TODO()
			fakeErr  = fmt.Errorf("fake err")
			caCert   = []byte("cert")
			caSecret = &corev1.Secret{Data: map[string][]byte{"ca.crt": caCert}}
		)

		BeforeEach(func() {
			dwdAccess = mockdwdaccess.NewMockAccessInterface(ctrl)

			botanist.StoreSecret("ca", caSecret)

			botanist.Shoot = &shootpkg.Shoot{
				Components: &shootpkg.Components{
					DependencyWatchdogAccess: dwdAccess,
				},
			}

			dwdAccess.EXPECT().SetCACertificate(caCert)
		})

		It("should set the secrets and deploy", func() {
			dwdAccess.EXPECT().Deploy(ctx)
			Expect(botanist.DeployDependencyWatchdogAccess(ctx)).To(Succeed())
		})

		It("should fail when the deploy function fails", func() {
			dwdAccess.EXPECT().Deploy(ctx).Return(fakeErr)
			Expect(botanist.DeployDependencyWatchdogAccess(ctx)).To(MatchError(fakeErr))
		})
	})
})
