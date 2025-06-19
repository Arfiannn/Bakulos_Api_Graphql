package mutation

import (
	"bakulos_grapghql/auth"
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"bakulos_grapghql/routes/types"
	"fmt"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

var Login = &graphql.Field{
	Type: types.LoginResponseType,
	Args: graphql.FieldConfigArgument{
		"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		email := getString(p, "email")
		password := getString(p, "password")

		// Cek ke tabel users dulu
		var user models.User
		if err := db.DB.Where("email = ?", email).First(&user).Error; err == nil {
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {
				token, err := auth.GenerateJWT(user.IDUser, "user")
				if err != nil {
					return nil, fmt.Errorf("gagal generate token")
				}
				return map[string]interface{}{
					"message": "Login berhasil",
					"role":    "user",
					"id_user": user.IDUser,
					"token":   token,
				}, nil
			}
		}

		// Cek ke tabel penjual
		var penjual models.Penjual
		if err := db.DB.Where("email = ?", email).First(&penjual).Error; err == nil {
			if err := bcrypt.CompareHashAndPassword([]byte(penjual.Password), []byte(password)); err == nil {
				token, err := auth.GenerateJWT(penjual.IDPenjual, "penjual")
				if err != nil {
					return nil, fmt.Errorf("gagal generate token")
				}
				return map[string]interface{}{
					"message":    "Login berhasil",
					"role":       "penjual",
					"id_penjual": penjual.IDPenjual,
					"token":      token,
				}, nil
			}
		}

		return nil, fmt.Errorf("email atau password salah")
	},
}
