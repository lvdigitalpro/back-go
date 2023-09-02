package postgres

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
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
	ctx context.Context, user contracts.UserContract,
) (*string, error) {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	var prepare *sql.Stmt

	if user.Role == "USER" {
		prepar, err := tx.PrepareContext(
			ctx,
			"insert into back.users (id, role, name, lastname, ir, email, password, created_at,updated_at) values ($1,$2, $3, $4, $5, $6, $7, $8, $9)",
		)
		if err != nil {
			return nil, err
		}

		prepare = prepar
	}

	if err != nil {
		return nil, err
	}

	if user.Role == "USER" {
		_, err = tx.StmtContext(ctx, prepare).Exec(
			user.ID, user.Role, user.Name, user.Lastname, user.Ir, user.Email, user.Password, user.CreatedAt,
			user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	response := "User created successfully"
	return &response, nil
}
func (r *UserRepository) UpdateUser(
	ctx context.Context, id string, name string, lastname string, ir string, email string, password string,
	enterpriseName *string, nrle *string,
) (*string, error) {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	prepare, err := tx.PrepareContext(
		ctx,
		"update back.users set name = $1, lastname = $2, ir = $3, email = $4, password = $5, enterpriseName = $6,nrle = $7 where id = $8",
	)

	if err != nil {
		return nil, err
	}

	_, err = tx.StmtContext(ctx, prepare).Exec(
		name, lastname, ir, email, password, enterpriseName, nrle, id,
	)

	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	response := "User updated successfully"

	return &response, nil
}
func (r *UserRepository) DeleteUser(
	ctx context.Context, id string, ir string, nrle *string, password string,
) (*string, error) {

	get := r.DB.QueryRow("select password from back.users where id = $1", id)

	var passwordDB string
	err := get.Scan(&passwordDB)

	err = bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(password))

	if err != nil {
		return nil, err
	}

	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	prepare, err := tx.PrepareContext(
		ctx, "delete from back.users where id = $1",
	)

	if err != nil {
		return nil, err
	}

	_, err = tx.StmtContext(ctx, prepare).Exec(id)

	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	response := "User deleted successfully"

	return &response, nil
}
func (r *UserRepository) GetUsers(ctx context.Context) ([]*contracts.UserContract, error) {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	prepare, err := tx.PrepareContext(ctx, "select * from back.users")

	if err != nil {
		return nil, err
	}

	rows, err := tx.StmtContext(ctx, prepare).Query()

	if err != nil {
		return nil, err
	}

	var users []*contracts.UserContract

	for rows.Next() {
		var user contracts.UserContract
		err = rows.Scan(
			&user.ID, &user.Role, &user.Name, &user.Lastname, &user.Ir, &user.Email, &user.Password,
			&user.EnterpriseName, &user.Nrle, pq.Array(user.Projects), &user.CreatedAt, &user.UpdatedAt,
		)

		user.Password = ""

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	err = rows.Close()

	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return users, nil
}
func (r *UserRepository) GetUser(ctx context.Context, id string) (*contracts.UserContract, error) {
	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	prepare, err := tx.PrepareContext(ctx, "select * from back.users where id = $1")

	if err != nil {
		return nil, err
	}

	row := tx.StmtContext(ctx, prepare).QueryRow(id)

	var user contracts.UserContract

	err = row.Scan(
		&user.ID, &user.Role, &user.Name, &user.Lastname, &user.Ir, &user.Email, &user.Password, &user.EnterpriseName,
		&user.Nrle, pq.Array(user.Projects), &user.CreatedAt, &user.UpdatedAt,
	)

	user.Password = ""

	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return &user, nil

}
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*contracts.UserContract, error) {
	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	prepare, err := tx.PrepareContext(ctx, "select * from back.users where email = $1")

	if err != nil {
		return nil, err
	}

	row := tx.StmtContext(ctx, prepare).QueryRow(email)

	user := contracts.UserContract{}

	err = row.Scan(
		&user.ID, &user.Role, &user.Name, &user.Lastname, &user.Ir, &user.Email, &user.Password,
		&user.EnterpriseName, &user.Nrle, pq.Array(user.Projects), &user.CreatedAt, &user.UpdatedAt,
	)

	user.Password = ""

	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return &user, nil

}
func (r *UserRepository) GetUserByIr(ctx context.Context, ir string) (*contracts.UserContract, error) {
	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	prepare, err := tx.PrepareContext(ctx, "select * from back.users where ir = $1")

	if err != nil {
		return nil, err
	}

	row := tx.StmtContext(ctx, prepare).QueryRow(ir)

	user := contracts.UserContract{}

	err = row.Scan(
		&user.ID, &user.Role, &user.Name, &user.Lastname, &user.Ir, &user.Email, &user.Password, &user.EnterpriseName,
		&user.Nrle, pq.Array(user.Projects), &user.CreatedAt, &user.UpdatedAt,
	)

	user.Password = ""

	if err != nil {
		return nil, err

	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *UserRepository) GetUserByProject(ctx context.Context, project string) (*contracts.UserContract, error) {

	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	prepare, err := tx.PrepareContext(ctx, "select * from back.users where arraycontains(projects,$1) = $1")

	if err != nil {
		return nil, err
	}

	row := tx.StmtContext(ctx, prepare).QueryRow(project)

	user := contracts.UserContract{}

	err = row.Scan(
		&user.ID, &user.Role, &user.Name, &user.Lastname, &user.Ir, &user.Email, &user.Password,
		&user.EnterpriseName, &user.Nrle, pq.Array(user.Projects), &user.CreatedAt, &user.UpdatedAt,
	)
	user.Password = ""

	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return &user, nil
}
