package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lvdigitalpro/back/src/data/services"
)

type authString string

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]

		validate, err := services.Validate(context.Background(), auth)

		if err != nil || !validate.Valid {
			fmt.Println(validate.Valid)
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		customClaim, _ := validate.Claims.(*services.Claims)

		ctx := context.WithValue(r.Context(), authString("auth"), customClaim)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func CtxValue(ctx context.Context) *services.Claims {
	raw, _ := ctx.Value(authString("auth")).(*services.Claims)
	return raw
}