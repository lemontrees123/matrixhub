// Copyright The MatrixHub Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package project_test

import (
	mhe2e "github.com/matrixhub-ai/matrixhub/test/e2e"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("test Project", Label("project"), func() {

	Context("test CreateProject API", func() {

		It("should create a project successfully", Label("smoke", "L00001"), func() {
			projectName := mhe2e.GenerateTestProjectName("project")
			GinkgoWriter.Printf("Creating project: %v\n", projectName)

			resp, err := apiClient.CreateProject(ctx, projectName)
			Expect(err).NotTo(HaveOccurred(), "should not have error")
			Expect(resp).NotTo(BeNil(), "response should not be nil")
			Expect(resp.Success).To(BeTrue(), "should return success")
			Expect(resp.HTTPStatusCode).To(BeNumerically(">=", 200), "should return 2xx status")
			Expect(resp.HTTPStatusCode).To(BeNumerically("<", 300), "should return 2xx status")

			GinkgoWriter.Printf("Create response: success=%v, status=%d\n", resp.Success, resp.HTTPStatusCode)

			_, _ = apiClient.DeleteProject(ctx, projectName)
		})

		It("should fail to create a project with empty name", Label("L00002"), func() {
			resp, err := apiClient.CreateProject(ctx, "")
			Expect(err).NotTo(HaveOccurred(), "should not have error")
			Expect(resp).NotTo(BeNil(), "response should not be nil")
			Expect(resp.Success).To(BeFalse(), "should return failure")
			Expect(resp.Error).NotTo(BeNil(), "error should not be nil")
			Expect(resp.Error.Code).To(Equal(3), "should return invalid argument error code")

			GinkgoWriter.Printf("Error response: code=%v, message=%v\n", resp.Error.Code, resp.Error.Message)
		})
	})

	Context("test project operations", func() {

		It("should create and get same project", Label("L00005"), func() {
			projectName := mhe2e.GenerateTestProjectName("project")
			GinkgoWriter.Printf("Create and get same project: %v\n", projectName)

			createResp, err := apiClient.CreateProject(ctx, projectName)
			Expect(err).NotTo(HaveOccurred(), "create should succeed")
			Expect(createResp.Success).To(BeTrue(), "create should return success")
			defer func() {
				_, _ = apiClient.DeleteProject(ctx, projectName)
			}()
		})
	})
})
