package types

import (
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id_user":  &graphql.Field{Type: graphql.Int},
		"nama":     &graphql.Field{Type: graphql.String},
		"email":    &graphql.Field{Type: graphql.String},
		"telepon":  &graphql.Field{Type: graphql.String},
		"profil":   &graphql.Field{Type: graphql.String},
		"password": &graphql.Field{Type: graphql.String},
	},
})
