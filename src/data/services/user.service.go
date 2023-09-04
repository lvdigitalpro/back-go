package services

import (
	"context"
	"errors"
	"strings"

	"github.com/lvdigitalpro/back/src/data/contracts"
	"github.com/lvdigitalpro/back/src/domain/utils"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	NewUser(
		ctx context.Context, input contracts.InputNewUserContract,
	) (*string, error)
	UpdateUser(
		ctx context.Context, input contracts.InputUpdateUserContract,
	) (*string, error)
	GetUsers(ctx context.Context) ([]*contracts.UserContract, error)
	GetUser(ctx context.Context, user_id string) (*contracts.UserContract, error)
	GetUserByEmail(ctx context.Context, email string) (*contracts.UserContract, error)
	GetUserByIr(ctx context.Context, ir string) (*contracts.UserContract, error)
	GetUserByProject(ctx context.Context, cod_project int) (*contracts.UserContract, error)
}

type UserService struct {
	Repo contracts.IUsersRepository
	IUserService
}

func NewUserService(repo contracts.IUsersRepository) IUserService {
	return &UserService{Repo: repo}
}

func (s *UserService) NewUser(
	ctx context.Context, input contracts.InputNewUserContract,
) (*string, error) {

	input.Name = strings.TrimSpace(input.Name)
	input.LastName = strings.TrimSpace(input.LastName)
	input.Email = strings.TrimSpace(input.Email)
	input.Password = strings.TrimSpace(input.Password)
	input.PasswordConfirmation = strings.TrimSpace(input.PasswordConfirmation)
	input.Ir = strings.TrimSpace(input.Ir)
	if input.Nrle != nil {
		*input.Nrle = strings.TrimSpace(*input.Nrle)
	}

	if input.Password != input.PasswordConfirmation {
		return nil, errors.New("passwords do not match")
	}

	isIrValid := utils.IsCPF(input.Ir)
	if !isIrValid {
		return nil, errors.New("invalid IR")
	}

	if input.Nrle != nil {
		isNrleValid := utils.IsCNPJ(*input.Nrle)
		if !isNrleValid {
			return nil, errors.New("invalid NRLE")
		}
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	input.Password = string(hashPassword)
	exec, err := s.Repo.NewUser(
		ctx, input,
	)
	if err != nil {
		return nil, err
	}
	return exec, nil
}

func (s *UserService) UpdateUser(
	ctx context.Context, input contracts.InputUpdateUserContract,
) (*string, error) {

	input.Name = strings.TrimSpace(input.Name)
	input.LastName = strings.TrimSpace(input.LastName)
	input.Email = strings.TrimSpace(input.Email)
	input.Password = strings.TrimSpace(input.Password)
	input.PasswordConfirmation = strings.TrimSpace(input.PasswordConfirmation)
	input.Ir = strings.TrimSpace(input.Ir)

	if input.Nrle != nil {
		*input.Nrle = strings.TrimSpace(*input.Nrle)
	}

	if input.Password != input.PasswordConfirmation {
		return nil, errors.New("passwords do not match")
	}

	isIrValid := utils.IsCPF(input.Ir)
	if !isIrValid {
		return nil, errors.New("invalid IR")
	}

	if input.Nrle != nil {
		isNrleValid := utils.IsCNPJ(*input.Nrle)
		if !isNrleValid {
			return nil, errors.New("invalid NRLE")
		}
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	input.Password = string(hashPassword)
	exec, err := s.Repo.UpdateUser(
		ctx, input,
	)
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

func (s *UserService) GetUserByProject(ctx context.Context, cod_project int) (*contracts.UserContract, error) {

	exec, err := s.Repo.GetUserByProject(ctx, cod_project)
	if err != nil {
		return nil, err
	}
	return exec, nil
}
