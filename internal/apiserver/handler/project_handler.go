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

package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	projectv1alpha1 "github.com/matrixhub-ai/matrixhub/api/go/v1alpha1"
	"github.com/matrixhub-ai/matrixhub/internal/domain/project"
	"github.com/matrixhub-ai/matrixhub/internal/infra/log"
)

type ProjectHandler struct {
	ProjectService project.IProjectService
}

func (ph *ProjectHandler) RemoveProjectMembers(ctx context.Context, request *projectv1alpha1.RemoveProjectMembersRequest) (*projectv1alpha1.RemoveProjectMembersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (ph *ProjectHandler) UpdateProject(ctx context.Context, request *projectv1alpha1.UpdateProjectRequest) (*projectv1alpha1.UpdateProjectResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (ph *ProjectHandler) ListProjects(ctx context.Context, request *projectv1alpha1.ListProjectsRequest) (*projectv1alpha1.ListProjectsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (ph *ProjectHandler) DeleteProject(ctx context.Context, request *projectv1alpha1.DeleteProjectRequest) (*projectv1alpha1.DeleteProjectResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (ph *ProjectHandler) AddProjectMemberWithRole(ctx context.Context, request *projectv1alpha1.AddProjectMemberWithRoleRequest) (*projectv1alpha1.AddProjectMemberWithRoleResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (ph *ProjectHandler) UpdateProjectMemberRole(ctx context.Context, request *projectv1alpha1.UpdateProjectMemberRoleRequest) (*projectv1alpha1.UpdateProjectMemberRoleResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func NewProjectHandler(ps project.IProjectService) *ProjectHandler {
	return &ProjectHandler{
		ProjectService: ps,
	}
}
func (ph *ProjectHandler) ListProjectMembers(ctx context.Context, request *projectv1alpha1.ListProjectMembersRequest) (*projectv1alpha1.ListProjectMembersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (ph *ProjectHandler) RegisterToServer(opt *ServerOptions) {
	// Register GRPC Handler
	projectv1alpha1.RegisterProjectsServer(opt.GRPCServer, ph)
	if err := projectv1alpha1.RegisterProjectsHandlerServer(context.Background(), opt.GatewayMux, ph); err != nil {
		log.Errorf("register handler error: %s", err.Error())
	}
}

func (ph *ProjectHandler) GetProject(ctx context.Context, request *projectv1alpha1.GetProjectRequest) (*projectv1alpha1.GetProjectResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Not implemented")
}

func (ph *ProjectHandler) CreateProject(ctx context.Context, request *projectv1alpha1.CreateProjectRequest) (*projectv1alpha1.CreateProjectResponse, error) {
	if err := request.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	param := &project.Project{
		Name: request.Name,
	}
	err := ph.ProjectService.CreateProject(ctx, param)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &projectv1alpha1.CreateProjectResponse{}, nil
}
