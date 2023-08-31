package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/lvdigitalpro/back/src/domain/entities"
	"github.com/lvdigitalpro/back/src/graph"
)

// NewUser is the resolver for the newUser field.
func (r *mutationResolver) NewUser(ctx context.Context, role entities.Role, name string, lastname string, ir string, email string, password string, passwordConfirmation string, enterpriseName *string, nrle *string) (string, error) {
	panic(fmt.Errorf("not implemented: NewUser - newUser"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, name *string, lastname *string, ir *string, email *string, password *string, passwordConfirmation *string, enterpriseName *string, nrle *string) (string, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string, ir string, nrle *string, password string, passwordConfirmation string) (string, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context) ([]*entities.User, error) {
	panic(fmt.Errorf("not implemented: GetUsers - getUsers"))
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*entities.User, error) {
	panic(fmt.Errorf("not implemented: GetUser - getUser"))
}

// GetUserByEmail is the resolver for the getUserByEmail field.
func (r *queryResolver) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	panic(fmt.Errorf("not implemented: GetUserByEmail - getUserByEmail"))
}

// GetUserByIr is the resolver for the getUserByIR field.
func (r *queryResolver) GetUserByIr(ctx context.Context, ir string) (*entities.User, error) {
	panic(fmt.Errorf("not implemented: GetUserByIr - getUserByIR"))
}

// GetUserByProject is the resolver for the getUserByProject field.
func (r *queryResolver) GetUserByProject(ctx context.Context, project string) (*entities.User, error) {
	panic(fmt.Errorf("not implemented: GetUserByProject - getUserByProject"))
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) NewUserEnterprise(ctx context.Context, role entities.Role, name string, lastname string, ir string, email string, password string, passwordConfirmation string, enterpriseName *string, nrle *string) (string, error) {
	panic(fmt.Errorf("not implemented: NewUserEnterprise - newUserEnterprise"))
}
func (r *mutationResolver) Hello(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: Hello - hello"))
}
func (r *queryResolver) Hello(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: Hello - hello"))
}