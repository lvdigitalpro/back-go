package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/lvdigitalpro/back/src/domain/entities"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*entities.AuthPayload, error) {
	exec, err := r.AuthService.Login(ctx, email, password)

	if err != nil {
		return nil, err
	}

	return &entities.AuthPayload{
		AccessToken:  exec.AccessToken,
		RefreshToken: exec.RefreshToken,
		User:         &exec.User,
	}, nil
}

// RefreshToken is the resolver for the refresh_token field.
func (r *mutationResolver) RefreshToken(ctx context.Context, refreshToken string) (*entities.AuthPayload, error) {
	exec, err := r.AuthService.RefreshToken(ctx, refreshToken)

	if err != nil {
		return nil, err
	}

	return &entities.AuthPayload{
		AccessToken:  exec.AccessToken,
		RefreshToken: exec.RefreshToken,
		User:         &exec.User,
	}, nil
}
