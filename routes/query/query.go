package query

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"bakulos_grapghql/routes/types"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type: graphql.NewList(types.UserType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.User
				return data, db.DB.Find(&data).Error
			},
		},
		"usersbyid": &graphql.Field{
			Type: types.UserType,
			Args: graphql.FieldConfigArgument{
				"id_user": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data models.User
				id := p.Args["id_user"].(int)
				db.DB.First(&data, id)
				return data, nil
			},
		},
	},
})
