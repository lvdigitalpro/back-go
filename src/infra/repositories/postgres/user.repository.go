package postgres

import (
	"context"
	"database/sql"

	"github.com/dgryski/trifles/uuid"
	"github.com/golang-module/carbon/v2"
	"github.com/lvdigitalpro/back/src/data/contracts"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
	contracts.IUsersRepository
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) NewUser(
	ctx context.Context, input contracts.InputNewUserContract,
) (*string, error) {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	prepare, err := tx.PrepareContext(ctx, "insert into platform.users (user_id, name, last_name, birth_date, email, password, created_at, type, ir, enterprise_name, nrle) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id := uuid.UUIDv4()

	_, err = tx.StmtContext(ctx, prepare).Exec(id, input.Name, input.LastName, input.BirthDate, input.Email, input.Password, carbon.Now().ToIso8601String(), input.Type, input.Ir, input.EnterpriseName, input.Nrle)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	response := "User created successfully"
	return &response, nil
}
func (r *UserRepository) UpdateUser(
	ctx context.Context, input contracts.InputUpdateUserContract,
) (*string, error) {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	get := tx.QueryRowContext(ctx, "select * from platform.users where user_id = $1", input.UserID)

	var oldUser contracts.UserContract
	err = get.Scan(&oldUser.UserID, &oldUser.Name, &oldUser.LastName, &oldUser.BirthDate, &oldUser.Email, &oldUser.Password, &oldUser.CreatedAt, &oldUser.UpdatedAt, &oldUser.Type, &oldUser.Ir, &oldUser.Nrle, &oldUser.EnterpriseName, &oldUser.CodProject)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if input.OldPassword != nil {

		err = bcrypt.CompareHashAndPassword([]byte(oldUser.Password), []byte(*input.OldPassword))

		if err != nil {
			tx.Rollback()
			return nil, err
		}
	} else {

		err = bcrypt.CompareHashAndPassword([]byte(oldUser.Password), []byte(input.PasswordConfirmation))

		if err != nil {
			tx.Rollback()
			return nil, err
		}

	}

	prepare, err := tx.PrepareContext(
		ctx,
		"update platform.users set name = $1, last_name = $2, birth_date = $3, email = $4, password = $5, created_at = $6, type = $7, ir = $8, enterprise_name = $9, nrle = $10, updated_at = $11 where user_id = $12;",
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.StmtContext(ctx, prepare).Exec(
		input.Name, input.LastName, input.BirthDate, input.Email, input.Password, input.CreatedAt, input.Type, input.Ir, input.EnterpriseName, input.Nrle, carbon.Now().ToIso8601String(), input.UserID,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	response := "User updated successfully"

	return &response, nil
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]*contracts.UserContract, error) {
	// tx, err := r.DB.BeginTx(ctx, nil)
	// if err != nil {
	// 	return nil, err
	// }

	// prepare, err := tx.PrepareContext(ctx, "select * from back.users")

	// if err != nil {
	// 	return nil, err
	// }

	// rows, err := tx.StmtContext(ctx, prepare).Query()

	// if err != nil {
	// 	return nil, err
	// }

	var users []*contracts.UserContract

	// for rows.Next() {
	// 	var user contracts.UserContract
	// 	err = rows.Scan(
	// 		&user.ID, &user.Role, &user.Name, &user.Lastname, &user.Ir, &user.Email, &user.Password,
	// 		&user.EnterpriseName, &user.Nrle, pq.Array(user.Projects), &user.CreatedAt, &user.UpdatedAt,
	// 	)

	// 	user.Password = ""

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	users = append(users, &user)
	// }

	// err = rows.Close()

	// if err != nil {
	// 	return nil, err
	// }

	// err = tx.Commit()

	// if err != nil {
	// 	return nil, err
	// }

	return users, nil
}
func (r *UserRepository) GetUser(ctx context.Context, user_id string) (*contracts.UserContract, error) {
	// tx, err := r.DB.BeginTx(ctx, nil)

	// if err != nil {
	// 	return nil, err
	// }

	// prepare, err := tx.PrepareContext(ctx, "select * from back.users where id = $1")

	// if err != nil {
	// 	return nil, err
	// }

	// row := tx.StmtContext(ctx, prepare).QueryRow(id)

	var user contracts.UserContract

	// err = row.Scan(
	// 	&user.ID, &user.Role, &user.Name, &user.Lastname, &user.Ir, &user.Email, &user.Password, &user.EnterpriseName,
	// 	&user.Nrle, pq.Array(user.Projects), &user.CreatedAt, &user.UpdatedAt,
	// )

	// user.Password = ""

	// if err != nil {
	// 	return nil, err
	// }

	// err = tx.Commit()

	// if err != nil {
	// 	return nil, err
	// }

	return &user, nil

}
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*contracts.UserContract, error) {
	// tx, err := r.DB.BeginTx(ctx, nil)

	// if err != nil {
	// 	return nil, err
	// }

	// prepare, err := tx.PrepareContext(ctx, "select * from back.users where email = $1")

	// if err != nil {
	// 	return nil, err
	// }

	// row := tx.StmtContext(ctx, prepare).QueryRow(email)

	user := contracts.UserContract{}

	// err = row.Scan(
	// 	&user.ID, &user.Role, &user.Name, &user.Lastname, &user.Ir, &user.Email, &user.Password,
	// 	&user.EnterpriseName, &user.Nrle, pq.Array(user.Projects), &user.CreatedAt, &user.UpdatedAt,
	// )

	// user.Password = ""

	// if err != nil {
	// 	return nil, err
	// }

	// err = tx.Commit()

	// if err != nil {
	// 	return nil, err
	// }

	return &user, nil

}
func (r *UserRepository) GetUserByIr(ctx context.Context, ir string) (*contracts.UserContract, error) {
	// tx, err := r.DB.BeginTx(ctx, nil)

	// if err != nil {
	// 	return nil, err
	// }

	// prepare, err := tx.PrepareContext(ctx, "select * from back.users where ir = $1")

	// if err != nil {
	// 	return nil, err
	// }

	// row := tx.StmtContext(ctx, prepare).QueryRow(ir)

	user := contracts.UserContract{}

	// err = row.Scan(
	// 	&user.ID, &user.Role, &user.Name, &user.Lastname, &user.Ir, &user.Email, &user.Password, &user.EnterpriseName,
	// 	&user.Nrle, pq.Array(user.Projects), &user.CreatedAt, &user.UpdatedAt,
	// )

	// user.Password = ""

	// if err != nil {
	// 	return nil, err

	// }

	// err = tx.Commit()

	// if err != nil {
	// 	return nil, err
	// }

	return &user, nil
}
func (r *UserRepository) GetUserByProject(ctx context.Context, cod_project int) (*contracts.UserContract, error) {

	// tx, err := r.DB.BeginTx(ctx, nil)

	// if err != nil {
	// 	return nil, err
	// }

	// prepare, err := tx.PrepareContext(ctx, "select * from back.users where arraycontains(projects,$1) = $1")

	// if err != nil {
	// 	return nil, err
	// }

	// row := tx.StmtContext(ctx, prepare).QueryRow(project)

	user := contracts.UserContract{}

	// err = row.Scan(
	// 	&user.ID, &user.Role, &user.Name, &user.Lastname, &user.Ir, &user.Email, &user.Password,
	// 	&user.EnterpriseName, &user.Nrle, pq.Array(user.Projects), &user.CreatedAt, &user.UpdatedAt,
	// )
	// user.Password = ""

	// if err != nil {
	// 	return nil, err
	// }

	// err = tx.Commit()

	// if err != nil {
	// 	return nil, err
	// }

	return &user, nil
}
