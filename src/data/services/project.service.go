package services

import (
	"context"
	"errors"

	"github.com/lvdigitalpro/back/src/data/contracts"
)

type IProjectService interface {
	NewProject(ctx context.Context, input contracts.InputNewProjectContract) (*string, error)
	UpdateProject(ctx context.Context, input contracts.InputUpdateProjectContract) (*string, error)
	DeleteProject(ctx context.Context, input contracts.InputDeleteProjectContract) (*string, error)
	GetProjects(ctx context.Context) ([]*contracts.ProjectContract, error)
	GetProject(ctx context.Context, cod_project int) (*contracts.ProjectContract, error)
}

type ProjectService struct {
	Repo contracts.IProjectsRepository
}

func NewProjectService(repo contracts.IProjectsRepository) IProjectService {
	return &ProjectService{Repo: repo}
}

func (s *ProjectService) NewProject(ctx context.Context, input contracts.InputNewProjectContract) (*string, error) {
	return s.Repo.NewProject(ctx, input)
}

func (s *ProjectService) UpdateProject(ctx context.Context, input contracts.InputUpdateProjectContract) (*string, error) {
	return s.Repo.UpdateProject(ctx, input)
}

func (s *ProjectService) DeleteProject(ctx context.Context, input contracts.InputDeleteProjectContract) (*string, error) {

	if input.Password != input.PasswordConfirmation {
		return nil, errors.New("passwords do not match")
	}
	return s.Repo.DeleteProject(ctx, input)
}

func (s *ProjectService) GetProjects(ctx context.Context) ([]*contracts.ProjectContract, error) {
	return s.Repo.GetProjects(ctx)
}

func (s *ProjectService) GetProject(ctx context.Context, cod_project int) (*contracts.ProjectContract, error) {

	return s.Repo.GetProject(ctx, cod_project)
}
