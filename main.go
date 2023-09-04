package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/lvdigitalpro/back/src/data/services"
	"github.com/lvdigitalpro/back/src/graph"
	"github.com/lvdigitalpro/back/src/infra/adapters"
	"github.com/lvdigitalpro/back/src/infra/repositories/postgres"
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

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &resolvers.Resolver{
					UserService: services.NewUserService(postgres.NewUserRepository(pgAdapter)),
				},
			},
		),
	)
	srv.AddTransport(&transport.Websocket{})
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
