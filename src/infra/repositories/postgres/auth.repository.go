package postgres

import (
	"context"
	"database/sql"

	"github.com/lvdigitalpro/back/src/data/contracts"
	"github.com/lvdigitalpro/back/src/data/services"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) contracts.IAuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

func (r *AuthRepository) Login(ctx context.Context, email string, password string) (*contracts.AuthResponse, error) {
	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	prepare, err := tx.PrepareContext(ctx, "SELECT * FROM platform.users WHERE email = $1")

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	row := tx.StmtContext(ctx, prepare).QueryRowContext(ctx, email)

	var user contracts.UserContract

	err = row.Scan(&user.UserID, &user.Name, &user.LastName, &user.BirthDate, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Type, &user.Ir, &user.Nrle, &user.EnterpriseName, &user.CodProject)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, err
	}

	at, err := services.JwtGenerateAT(ctx, user.UserID)
	if err != nil {
		return nil, err
	}

	rt, err := services.JwtGenerateRT(ctx, user.UserID)

	if err != nil {
		return nil, err
	}
	user.Password = ""

	return &contracts.AuthResponse{
		AccessToken:  at,
		RefreshToken: rt,
		User:         user,
	}, nil
}

func (r *AuthRepository) RefreshToken(ctx context.Context, user_id string) (*contracts.AuthResponse, error) {
	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	prepare, err := tx.PrepareContext(ctx, "SELECT * FROM platform.users WHERE user_id = $1")

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	row := tx.StmtContext(ctx, prepare).QueryRowContext(ctx, user_id)

	var user contracts.UserContract

	err = row.Scan(&user.UserID, &user.Name, &user.LastName, &user.BirthDate, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Type, &user.Ir, &user.Nrle, &user.EnterpriseName, &user.CodProject)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	at, err := services.JwtGenerateAT(ctx, user.UserID)
	if err != nil {
		return nil, err
	}

	rt, err := services.JwtGenerateRT(ctx, user.UserID)

	if err != nil {
		return nil, err
	}
	user.Password = ""

	return &contracts.AuthResponse{
		AccessToken:  at,
		RefreshToken: rt,
		User:         user,
	}, nil

}
