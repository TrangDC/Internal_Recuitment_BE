//go:build ignore

package main

import (
	"errors"
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension()
	if !errors.Is(err, nil) {
		log.Fatalf("Error: failed creating entgql extension: %v", err)
	}
	if err := entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureExecQuery,
		},
	}, entc.Extensions(ex)); !errors.Is(err, nil) {
		log.Fatalf("Error: failed running ent codegen: %v", err)
	}
}
