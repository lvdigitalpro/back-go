package postgres

import (
	"github.com/lvdigitalpro/back/src/data/contracts"
	"github.com/lvdigitalpro/back/src/infra/database"
)

type UserRepository struct {
	DB *database.DBContract[contracts.UserContract]
	contracts.IUsersRepository
}

func NewUserRepository(db *database.DBContract[contracts.UserContract]) *UserRepository {
	return &UserRepository{DB: db}
}
