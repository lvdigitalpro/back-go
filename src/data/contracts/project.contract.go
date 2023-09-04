package contracts

import (
	"context"

	"github.com/lvdigitalpro/back/src/domain/entities"
)

type ProjectContract = entities.Project
type InputNewProjectContract = entities.InputNewProject
type InputUpdateProjectContract = entities.InputUpdateProject
type InputDeleteProjectContract = entities.InputDeleteProject
type StatusProjectContract = entities.Status

const (
	QUEUE       = entities.StatusQueue
	IN_REVIEW   = entities.StatusInReview
	IN_PROGRESS = entities.StatusInProgress
	DONE        = entities.StatusDone
)

type IProjectsRepository interface {
	NewProject(ctx context.Context, input InputNewProjectContract) (*string, error)
	UpdateProject(ctx context.Context, input InputUpdateProjectContract) (*string, error)
	DeleteProject(ctx context.Context, input InputDeleteProjectContract) (*string, error)
	GetProjects(ctx context.Context) ([]*ProjectContract, error)
	GetProject(ctx context.Context, cod_project int) (*ProjectContract, error)
}
