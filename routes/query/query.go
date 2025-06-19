package query

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"bakulos_grapghql/routes/types"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type: graphql.NewList(types.UserType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userClaims := p.Context.Value("user")
				if userClaims == nil {
					return nil, fmt.Errorf("Unauthorized")
				}

				claims := userClaims.(jwt.MapClaims)
				idUser := uint(claims["id_user"].(float64))

				var user models.User
				if err := db.DB.First(&user, idUser).Error; err != nil {
					return nil, err
				}

				return []models.User{user}, nil
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

		"idpenjuals": &graphql.Field{
			Type: graphql.NewList(types.PenjualType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Penjual
				return data, db.DB.Find(&data).Error
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

		"products": &graphql.Field{
			Type: graphql.NewList(types.ProductType),
			Args: graphql.FieldConfigArgument{
				"id_user": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Product

				err := db.DB.Preload("Penjual").Find(&data).Error
				if err != nil {
					return nil, err
				}
				var idUser uint = 0
				if val, ok := p.Args["id_user"]; ok && val != nil {
					idUser = uint(val.(int))
				}

				for i := range data {
					var fav models.Favorite
					err := db.DB.Where("id_product = ? AND id_user = ?", data[i].IDProduct, idUser).First(&fav).Error
					if err == nil {
						data[i].IDFavorite = &fav.IDFavorite
					} else {
						data[i].IDFavorite = nil
					}
				}
				return data, nil
			},
		},

		"keranjangs": &graphql.Field{
			Type: graphql.NewList(types.KeranjangType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Keranjang
				err := db.DB.Preload("Product").Preload("User").Find(&data).Error
				if err != nil {
					return nil, err
				}
				return data, nil
			},
		},

		"historys": &graphql.Field{
			Type: graphql.NewList(types.HistoryType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.History
				err := db.DB.Preload("Product").Preload("User").Find(&data).Error
				return data, err
			},
		},

		"checkouts": &graphql.Field{
			Type: graphql.NewList(types.CheckoutType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Checkout
				err := db.DB.Preload("Product").Preload("User").Preload("Alamat").Preload("Keranjang").Find(&data).Error
				if err != nil {
					return nil, err
				}
				return data, nil
			},
		},

		"favorites": &graphql.Field{
			Type: graphql.NewList(types.FavoriteType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Favorite
				err := db.DB.Preload("Product").Preload("User").Find(&data).Error
				if err != nil {
					return nil, err
				}
				return data, nil
			},
		},

		"chats": &graphql.Field{
			Type: graphql.NewList(types.ChatType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var chats []models.Chat
				err := db.DB.
					Preload("User").
					Preload("Penjual").
					Order("created_at ASC").
					Find(&chats).Error
				return chats, err
			},
		},
		"chatsByUserPenjual": &graphql.Field{
			Type: graphql.NewList(types.ChatType),
			Args: graphql.FieldConfigArgument{
				"id_user":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"id_penjual": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var chats []models.Chat
				idUser := p.Args["id_user"].(int)
				idPenjual := p.Args["id_penjual"].(int)
				err := db.DB.
					Where("id_user = ? AND id_penjual = ?", idUser, idPenjual).
					Order("created_at ASC"). // ✅ tambahkan sorting by waktu juga di sini
					Preload("User").
					Preload("Penjual").
					Find(&chats).Error
				return chats, err
			},
		},
		"countUnreadChatByPenjual": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"id_penjual": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idPenjual := p.Args["id_penjual"].(int)
				var count int64
				err := db.DB.Model(&models.Chat{}).
					Where("id_penjual = ? AND sender = ? AND is_read = ?", idPenjual, "user", false).
					Count(&count).Error
				return int(count), err
			},
		},
		"countUnreadChatByUser": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"id_user": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idUser := p.Args["id_user"].(int)
				var count int64
				err := db.DB.Model(&models.Chat{}).
					Where("id_user = ? AND sender = ? AND is_read = ?", idUser, "penjual", false).
					Count(&count).Error
				return int(count), err
			},
		},
	},
})
