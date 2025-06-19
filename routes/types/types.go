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

var PenjualType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Penjual",
	Fields: graphql.Fields{
		"id_penjual": &graphql.Field{Type: graphql.Int},
		"nama":       &graphql.Field{Type: graphql.String},
		"email":      &graphql.Field{Type: graphql.String},
		"telepon":    &graphql.Field{Type: graphql.String},
		"password":   &graphql.Field{Type: graphql.String},
		"profil":     &graphql.Field{Type: graphql.String},
	},
})

var LoginResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "LoginResponse",
	Fields: graphql.Fields{
		"id_user":    &graphql.Field{Type: graphql.Int},
		"id_penjual": &graphql.Field{Type: graphql.Int},
		"message":    &graphql.Field{Type: graphql.String},
		"role":       &graphql.Field{Type: graphql.String},
		"token":      &graphql.Field{Type: graphql.String},
	},
})
