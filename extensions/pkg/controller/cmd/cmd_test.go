// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package cmd

import (
	"errors"

	"github.com/onsi/gomega/gbytes"
	logzap "sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/gardener/gardener/pkg/utils/test"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cmd", func() {
	var (
		ctrl *gomock.Controller
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
	})
	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("#LogErrAndExit", func() {
		It("should log the error and exit", func() {
			called := false
			defer test.WithVar(&Exit, func(code int) {
				called = true
				Expect(code).To(Equal(1))
			})()

			logBuffer := gbytes.NewBuffer()
			defer test.WithVar(&Log, logzap.New(logzap.JSONEncoder(), logzap.WriteTo(logBuffer)))()

			LogErrAndExit(errors.New("error"), "msg")
			Expect(called).To(BeTrue())
			Eventually(logBuffer).Should(gbytes.Say(`"msg":"msg","error":"error"`))
		})
	})
})
