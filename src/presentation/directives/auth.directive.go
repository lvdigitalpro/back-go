package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/lvdigitalpro/back/src/presentation/middlewares"
	"github.com/vektah/gqlparser/gqlerror"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	tokenData := middlewares.CtxValue(ctx)
	if tokenData == nil {
		return nil, &gqlerror.Error{
			Message: "Access Denied",
		}
	}

	return next(ctx)
}
