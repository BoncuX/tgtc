package gqlserver

import (
	"github.com/graphql-go/graphql"
	"github.com/radityaqb/tgtc/backend/dictionary"
)

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) GetProduct() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["product_id"].(int)

		// return r.p.GetProduct(ctx, int64(id))   	// update to use Usecase from previous session
		return dictionary.Product{
			ID:   int64(id),
			Name: "test",
		}, nil
	}
}

func (r *Resolver) AddProduct() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {

		name, _ := p.Args["product_name"].(string)

		/*
			// update to use Usecase from previous session
			req := dictionary.Product{
				Name: name,
			}

			return r.p.AddProduct(ctx, req)
		*/

		return dictionary.Product{
			ID:   0,
			Name: name,
		}, nil
	}
}
