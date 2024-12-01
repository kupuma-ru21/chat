package resolvers

import (
	"backend/ent"
	"backend/gqlgen"

	"github.com/99designs/gqlgen/graphql"
)

// Resolver is the resolver root.
type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return gqlgen.NewExecutableSchema(gqlgen.Config{
		Resolvers: &Resolver{client},
	})
}
