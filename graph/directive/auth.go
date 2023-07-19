package directive

import (
	"context"
	"gqlgen-playground/graph/gql"

	"github.com/99designs/gqlgen/graphql"
)

var Directive gql.DirectiveRoot = gql.DirectiveRoot{
	IsAuthenticated: IsAuthenticated,
}

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	// TODO:
	return next(ctx)
}
