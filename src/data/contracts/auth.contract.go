package contracts

import "context"

type AuthResponse struct {
	RefreshToken string       `json:"refreshToken"`
	AccessToken  string       `json:"accessToken"`
	User         UserContract `json:"user"`
}

type IAuthRepository interface {
	Login(ctx context.Context, email string, password string) (*AuthResponse, error)
}
