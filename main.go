package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Name:              "Welcome Message",
			Type:              graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "World", nil
			},
			Description:       "Get the welcome message",
		},
	}
	rootQuery := graphql.ObjectConfig{
		Name:        "RootQuery",
		Fields:      fields,
		Description: "Root Query description",
	}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create schema %v", err)
	}

	query := `{
		hello
	}`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %v", r.Errors)
	}
	rJson, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJson)
}
