package contracts

import (
	"context"

	"github.com/lvdigitalpro/back/src/domain/entities"
)

type UserContract = entities.User // This is a contract

// type UserContractInterface interface {
// 	Validate() bool
// 	AssignUser(user UserContract) UserContract
// 	NewUser(role RoleContract, name string, lastname string, ir string, email string, password string) (
// 		*UserContract, error,
// 	)
// 	AssignEnterprise(enterpriseName string, nrle string)
// }

type InputNewUserContract = entities.InputNewUser

type InputUpdateUserContract = entities.InputUpdateUser

type IUsersRepository interface {
	NewUser(ctx context.Context, input InputNewUserContract) (*string, error)
	UpdateUser(
		ctx context.Context, input InputUpdateUserContract,
	) (*string, error)
	GetUsers(ctx context.Context) ([]*UserContract, error)
	GetUser(ctx context.Context, user_id string) (*UserContract, error)
	GetUserByEmail(ctx context.Context, email string) (*UserContract, error)
	GetUserByIr(ctx context.Context, ir string) (*UserContract, error)
	GetUserByProject(ctx context.Context, cod_project int) (*UserContract, error)
}
