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

// func (r *AuthService) Login(ctx context.Context, email string, password string) (*contracts.AuthResponse, error) {

// 	user, err := r.UserService.GetUserByEmail(ctx, email)

// 	if err != nil {
// 		return nil, err
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

// 	if err != nil {
// 		return nil, err
// 	}

// 	claimsRefreshToken := &Claims{
// 		user.UserID,
// 		jwt.RegisteredClaims{
// 			ID:        uuid.UUIDv4(),
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 7)),
// 			NotBefore: jwt.NewNumericDate(time.Now()),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 			Issuer:    "lvdigitalpro",
// 			Subject:   "loginRT",
// 		},
// 	}

// 	claimsAccessToken := &Claims{
// 		user.UserID,
// 		jwt.RegisteredClaims{
// 			ID:        uuid.UUIDv4(),
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
// 			NotBefore: jwt.NewNumericDate(time.Now()),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 			Issuer:    "lvdigitalpro",
// 			Subject:   "loginAT",
// 		},
// 	}

// 	acTk := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claimsAccessToken)

// 	accessToken, err := acTk.SignedString(os.Getenv("ACCESS_TOKEN_SECRET"))

// 	if err != nil {
// 		return nil, err
// 	}

// 	rfTk := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claimsRefreshToken)

// 	refreshToken, err := rfTk.SignedString(os.Getenv("REFRESH_TOKEN_SECRET"))

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &contracts.AuthResponse{
// 		AccessToken:  accessToken,
// 		RefreshToken: refreshToken,
// 	}, nil
// }

func Validate(ctx context.Context, token string) (*jwt.Token, error) {

	return jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		if _, mt := token.Method.(*jwt.SigningMethodHMAC); !mt {
			return nil, errors.New("unexpected signing method")
		}
		secret := []byte(os.Getenv("SECRET_AT"))
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
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
