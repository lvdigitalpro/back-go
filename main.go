package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/lvdigitalpro/back/src/graph"
	"github.com/lvdigitalpro/back/src/presentation/resolvers"
	"log"
	"net/http"
)

func main() {

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}}))
	srv.AddTransport(&transport.Websocket{})
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Fatal(http.ListenAndServe(":8080", srv))
}
