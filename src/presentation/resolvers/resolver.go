//go:generate go run github.com/99designs/gqlgen generate

package resolvers

import (
	"github.com/99designs/gqlgen/plugin/federation/testdata/entityresolver/generated"
	"github.com/lvdigitalpro/back/src/data/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService    services.IUserService
	ProjectService services.IProjectService
}

func (r *Resolver) Entity() generated.EntityResolver {
	//TODO implement me
	panic("implement me")
}
