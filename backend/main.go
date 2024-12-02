package main

import (
	"backend/ent"
	"backend/ent/migrate"
	"backend/gqlgen/resolvers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	env := os.Getenv("ENV")

	if env == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port is not set")
	}

	postgresUser := os.Getenv("POSTGRES_USER")
	if postgresUser == "" {
		log.Fatal("postgresUser is not set")
	}

	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if postgresPassword == "" {
		log.Fatal("postgresPassword is not set")
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost == "" {
		log.Fatal("postgresHost is not set")
	}

	postgresPort := os.Getenv("POSTGRES_PORT")
	if postgresHost == "" {
		log.Fatal("postgresPort is not set")
	}

	postgresDb := os.Getenv("POSTGRES_DB")
	if postgresDb == "" {
		log.Fatal("postgresDb is not set")
	}

	postgresqlUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		postgresUser, postgresPassword, postgresHost,
		postgresPort, postgresDb,
	)

	// Create ent.Client and run the schema migration.
	client, err := ent.Open(dialect.Postgres, postgresqlUrl)
	if err != nil {
		log.Fatal("opening ent client: ", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client: ", err)
	}

	srv := handler.NewDefaultServer(resolvers.NewSchema(client))
	http.Handle("/",
		playground.Handler("Todo", "/query"),
	)
	http.Handle("/query", srv)

	log.Printf("listening on http://localhost:%s/", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
