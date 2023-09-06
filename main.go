package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lvdigitalpro/back/src/data/services"
	"github.com/lvdigitalpro/back/src/graph"
	"github.com/lvdigitalpro/back/src/infra/adapters"
	"github.com/lvdigitalpro/back/src/infra/repositories/postgres"
	"github.com/lvdigitalpro/back/src/presentation/directives"
	"github.com/lvdigitalpro/back/src/presentation/middlewares"
	"github.com/lvdigitalpro/back/src/presentation/resolvers"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	pgAdapter := adapters.PgAdapter()
	defer pgAdapter.Close()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &resolvers.Resolver{
			UserService:    services.NewUserService(postgres.NewUserRepository(pgAdapter)),
			ProjectService: services.NewProjectService(postgres.NewProjectsRepository(pgAdapter)),
			AuthService:    services.NewAuthService(postgres.NewAuthRepository(pgAdapter)),
		},
		Directives: graph.DirectiveRoot{Auth: directives.Auth},
	}))

	r := mux.NewRouter()

	r.Use(middlewares.AuthMiddleware)

	srv.AddTransport(&transport.Websocket{})
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
