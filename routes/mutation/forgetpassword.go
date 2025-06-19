package mutation

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"bakulos_grapghql/routes/types"
	"fmt"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

var ForgetPassword = &graphql.Field{
	Type: types.ForgetPasswordResponseType,
	Args: graphql.FieldConfigArgument{
		"email":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"new_password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		email := getString(p, "email")
		newPassword := getString(p, "new_password")

		// ✅ Cek apakah email ada di tabel User dulu
		var user models.User
		if err := db.DB.Where("email = ?", email).First(&user).Error; err == nil {
			// ✅ Hash password baru
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
			if err != nil {
				return nil, fmt.Errorf("gagal hash password baru: %v", err)
			}

			if err := db.DB.Model(&user).Update("password", string(hashedPassword)).Error; err != nil {
				return nil, fmt.Errorf("gagal update password: %v", err)
			}
			return map[string]interface{}{
				"message": "Password berhasil diganti",
				"role":    "user",
			}, nil
		}

		// ✅ Jika tidak ada di User, cek di Penjual
		var penjual models.Penjual
		if err := db.DB.Where("email = ?", email).First(&penjual).Error; err == nil {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
			if err != nil {
				return nil, fmt.Errorf("gagal hash password baru: %v", err)
			}

			if err := db.DB.Model(&penjual).Update("password", string(hashedPassword)).Error; err != nil {
				return nil, fmt.Errorf("gagal update password: %v", err)
			}
			return map[string]interface{}{
				"message": "Password berhasil diganti",
				"role":    "penjual",
			}, nil
		}

		return nil, fmt.Errorf("email tidak ditemukan di user maupun penjual")
	},
}
