package mutation

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"fmt"

	"github.com/graphql-go/graphql"
)

var MarkChatAsRead = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"id_user": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"id_penjual": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"role": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idUser := p.Args["id_user"].(int)
		idPenjual := p.Args["id_penjual"].(int)
		role := p.Args["role"].(string)

		var sender string
		if role == "user" {
			sender = "penjual"
		} else if role == "penjual" {
			sender = "user"
		} else {
			return nil, fmt.Errorf("role tidak valid")
		}

		err := db.DB.Model(&models.Chat{}).
			Where("id_user = ? AND id_penjual = ? AND sender = ? AND is_read = ?", idUser, idPenjual, sender, false).
			Update("is_read", true).Error

		if err != nil {
			return nil, fmt.Errorf("gagal update status pesan: %v", err)
		}

		return "Pesan ditandai telah dibaca", nil
	},
}
