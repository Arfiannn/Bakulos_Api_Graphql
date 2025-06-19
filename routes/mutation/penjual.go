package mutation

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"bakulos_grapghql/routes/types"
	"fmt"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

var CreatePenjual = &graphql.Field{
	Type: types.PenjualType,
	Args: graphql.FieldConfigArgument{
		"nama":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"telepon":  &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		email := p.Args["email"].(string)

		var existingPenjual models.Penjual
		if err := db.DB.Where("email = ?", email).First(&existingPenjual).Error; err == nil {
			return nil, fmt.Errorf("email sudah terdaftar")
		}

		var existingUser models.User
		if err := db.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
			return nil, fmt.Errorf("email sudah terdaftar di user")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Args["password"].(string)), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		penjual := models.Penjual{
			Nama:     p.Args["nama"].(string),
			Email:    email,
			Telepon:  getString(p, "telepon"),
			Password: string(hashedPassword),
		}

		if err := db.DB.Create(&penjual).Error; err != nil {
			return nil, err
		}

		return penjual, nil
	},
}

var UpdatePenjual = &graphql.Field{
	Type: types.PenjualType,
	Args: graphql.FieldConfigArgument{
		"id_penjual":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"nama":         &graphql.ArgumentConfig{Type: graphql.String},
		"telepon":      &graphql.ArgumentConfig{Type: graphql.String},
		"password":     &graphql.ArgumentConfig{Type: graphql.String},
		"old_password": &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := getInt(p, "id_penjual")
		var penjual models.Penjual

		if err := db.DB.First(&penjual, id).Error; err != nil {
			return nil, fmt.Errorf("penjual tidak ditemukan")
		}

		updates := map[string]interface{}{}

		if v := getString(p, "nama"); v != "" {
			updates["nama"] = v
		}
		if v := getString(p, "telepon"); v != "" {
			updates["telepon"] = v
		}

		if v := getString(p, "password"); v != "" {
			oldPassword := getString(p, "old_password")
			if oldPassword == "" {
				return nil, fmt.Errorf("password lama wajib diisi")
			}
			err := bcrypt.CompareHashAndPassword([]byte(penjual.Password), []byte(oldPassword))
			if err != nil {
				return nil, fmt.Errorf("password lama salah")
			}

			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
			if err != nil {
				return nil, err
			}
			updates["password"] = string(hashedPassword)
		}

		if err := db.DB.Model(&penjual).Updates(updates).Error; err != nil {
			return nil, err
		}

		db.DB.First(&penjual, id)
		return penjual, nil
	},
}

var UpdatePenjualProfil = &graphql.Field{
	Type: types.PenjualType,
	Args: graphql.FieldConfigArgument{
		"id_penjual": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"profil":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		IDPenjual := p.Args["id_penjual"].(int)
		profil := p.Args["profil"].(string)
		var penjual models.Penjual
		if err := db.DB.First(&penjual, IDPenjual).Error; err != nil {
			return nil, fmt.Errorf("penjual tidak ditemukan")
		}
		penjual.Profil = profil
		if err := db.DB.Save(&penjual).Error; err != nil {
			return nil, err
		}
		return penjual, nil
	},
}
