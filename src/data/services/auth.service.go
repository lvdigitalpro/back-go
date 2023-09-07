package services

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lvdigitalpro/back/src/data/contracts"
)

type IAuthService interface {
	Login(ctx context.Context, email string, password string) (*contracts.AuthResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*contracts.AuthResponse, error)
}

type AuthService struct {
	AuthRepository contracts.IAuthRepository
}

func NewAuthService(authRepository contracts.IAuthRepository) IAuthService {
	return &AuthService{
		AuthRepository: authRepository,
	}
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func (r *AuthService) Login(ctx context.Context, email string, password string) (*contracts.AuthResponse, error) {
	exec, err := r.AuthRepository.Login(ctx, email, password)

	if err != nil {
		return nil, err
	}

	return exec, nil
}

func (r *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*contracts.AuthResponse, error) {

	bearer := "Bearer "
	refreshToken = refreshToken[len(bearer):]

	validate, err := ValidateRT(ctx, refreshToken)

	if err != nil || !validate.Valid {
		return nil, errors.New("invalid token")
	}

	customClaim, _ := validate.Claims.(*Claims)

	exec, err := r.AuthRepository.RefreshToken(ctx, customClaim.UserID)

	if err != nil {
		return nil, err
	}

	return exec, nil

}

func ValidateAT(ctx context.Context, token string) (*jwt.Token, error) {

	return jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		if _, mt := token.Method.(*jwt.SigningMethodHMAC); !mt {
			return nil, errors.New("unexpected signing method")
		}
		secret := []byte(os.Getenv("SECRET_AT"))
		return secret, nil

	})

}

func ValidateRT(ctx context.Context, token string) (*jwt.Token, error) {

	return jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		if _, mt := token.Method.(*jwt.SigningMethodHMAC); !mt {
			return nil, errors.New("unexpected signing method")
		}
		secret := []byte(os.Getenv("SECRET_RT"))
		return secret, nil

	})

}

func JwtGenerateRT(ctx context.Context, userID string) (string, error) {
	claimsRefreshToken := &Claims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 7)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		}}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefreshToken)

	token, err := t.SignedString([]byte(os.Getenv("SECRET_RT")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func JwtGenerateAT(ctx context.Context, userID string) (string, error) {
	claimsRefreshToken := &Claims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		}}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefreshToken)

	token, err := t.SignedString([]byte(os.Getenv("SECRET_AT")))
	if err != nil {
		return "", err
	}

	return token, nil
}
