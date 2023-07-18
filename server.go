package main

import (
	"database/sql"
	"fmt"
	"gqlgen-playground/graph"
	"gqlgen-playground/graph/gql"
	"gqlgen-playground/service"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	_ "github.com/lib/pq"
)

const (
	defaultPort = "8080"

	// DB
	host     = "localhost"
	port     = 5432
	user     = "pguser"
	password = "pgpass"
	dbname   = "example"
)

func main() {
	endpointPort := os.Getenv("PORT")
	if endpointPort == "" {
		endpointPort = defaultPort
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// health check
	if err := db.Ping(); err != nil {
		panic(err)
	}

	service := service.New(db)

	srv := handler.NewDefaultServer(gql.NewExecutableSchema(gql.Config{Resolvers: &graph.Resolver{
		Service: service,
	}}))

	http.Handle("/", playground.ApolloSandboxHandler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", endpointPort)
	log.Fatal(http.ListenAndServe(":"+endpointPort, nil))
}
