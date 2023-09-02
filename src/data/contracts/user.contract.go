package contracts

import (
	"context"
	"github.com/lvdigitalpro/back/src/domain/entities"
)

type UserContract = entities.User // This is a contract

type RoleContract = entities.Role // This is a contract

type UserContractInterface interface {
	Validate() bool
	AssignUser(user UserContract) UserContract
	NewUser(role RoleContract, name string, lastname string, ir string, email string, password string) (
		*UserContract, error,
	)
	AssignEnterprise(enterpriseName string, nrle string)
}

type IUsersRepository interface {
	NewUser(ctx context.Context, user UserContract) (*string, error)
	UpdateUser(
		ctx context.Context, id string, name string, lastname string, ir string, email string, password string,
		enterpriseName *string, nrle *string,
	) (*string, error)
	DeleteUser(
		ctx context.Context, id string, ir string, nrle *string, password string,
	) (*string, error)
	GetUsers(ctx context.Context) ([]*UserContract, error)
	GetUser(ctx context.Context, id string) (*UserContract, error)
	GetUserByEmail(ctx context.Context, email string) (*UserContract, error)
	GetUserByIr(ctx context.Context, ir string) (*UserContract, error)
	GetUserByProject(ctx context.Context, project string) (*UserContract, error)
}
