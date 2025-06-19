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

		"penjuals": &graphql.Field{
			Type: graphql.NewList(types.PenjualType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Penjual
				return data, db.DB.Find(&data).Error
			},
		},
		"penjualsbyid": &graphql.Field{
			Type: types.PenjualType,
			Args: graphql.FieldConfigArgument{
				"id_penjual": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data models.Penjual
				id := p.Args["id_penjual"].(int)
				db.DB.First(&data, id)
				return data, nil
			},
		},

		"alamats": &graphql.Field{
			Type: graphql.NewList(types.AlamatType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Alamat
				if err := db.DB.Preload("User").Find(&data).Error; err != nil {
					return nil, err
				}
				return data, nil
			},
		},
	},
})
