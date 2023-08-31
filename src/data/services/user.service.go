package services

import (
	"context"
	"github.com/lvdigitalpro/back/src/data/contracts"
)

type IUserService interface {
	NewUser(ctx context.Context, role string, name string, lastname string, ir string, email string, password string, passwordConfirmation string, enterpriseName *string, nrle *string) (*string, error)
	UpdateUser(ctx context.Context, id string, name string, lastname string, ir string, email string, password string, passwordConfirmation string, enterpriseName *string, nrle *string) (*string, error)
	DeleteUser(ctx context.Context, id string, ir string, nrle *string, password string, passwordConfirmation string) (*string, error)
	GetUsers(ctx context.Context) ([]*contracts.UserContract, error)
	GetUser(ctx context.Context, id string) (*contracts.UserContract, error)
	GetUserByEmail(ctx context.Context, email string) (*contracts.UserContract, error)
	GetUserByIr(ctx context.Context, ir string) (*contracts.UserContract, error)
	GetUserByProject(ctx context.Context, project string) (*contracts.UserContract, error)
}

type UserService struct {
	Repo contracts.IUsersRepository
	IUserService
}

func NewUserService(repo contracts.IUsersRepository) IUserService {
	return &UserService{Repo: repo}
}

func (s *UserService) NewUser(ctx context.Context, role string, name string, lastname string, ir string, email string, password string, passwordConfirmation string, enterpriseName *string, nrle *string) (*string, error) {

	exec, err := s.Repo.NewUser(ctx, role, name, lastname, ir, email, password, passwordConfirmation, enterpriseName, nrle)
	if err != nil {
		return nil, err
	}
	return exec, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id string, name string, lastname string, ir string, email string, password string, passwordConfirmation string, enterpriseName *string, nrle *string) (*string, error) {

	exec, err := s.Repo.UpdateUser(ctx, id, name, lastname, ir, email, password, passwordConfirmation, enterpriseName, nrle)
	if err != nil {
		return nil, err
	}

	return exec, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string, ir string, nrle *string, password string, passwordConfirmation string) (*string, error) {

	exec, err := s.Repo.DeleteUser(ctx, id, ir, nrle, password, passwordConfirmation)
	if err != nil {
		return nil, err
	}
	return exec, nil
}

func (s *UserService) GetUsers(ctx context.Context) ([]*contracts.UserContract, error) {

	exec, err := s.Repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return exec, nil
}

func (s *UserService) GetUser(ctx context.Context, id string) (*contracts.UserContract, error) {

	exec, err := s.Repo.GetUser(ctx, id)

	if err != nil {
		return nil, err
	}
	return exec, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*contracts.UserContract, error) {

	exec, err := s.Repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return exec, nil
}

func (s *UserService) GetUserByIr(ctx context.Context, ir string) (*contracts.UserContract, error) {

	exec, err := s.Repo.GetUserByIr(ctx, ir)
	if err != nil {
		return nil, err
	}
	return exec, nil
}

func (s *UserService) GetUserByProject(ctx context.Context, project string) (*contracts.UserContract, error) {

	exec, err := s.Repo.GetUserByProject(ctx, project)
	if err != nil {
		return nil, err
	}
	return exec, nil
}
