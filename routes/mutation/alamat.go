package mutation

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"bakulos_grapghql/routes/types"
	"fmt"

	"github.com/graphql-go/graphql"
)

var CreateAlamat = &graphql.Field{
	Type: types.AlamatType,
	Args: graphql.FieldConfigArgument{
		"id_user":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"alamat":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"namaA":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"teleponA": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idUser := getInt(p, "id_user")

		// Cek user valid
		var user models.User
		if err := db.DB.First(&user, idUser).Error; err != nil {
			return nil, fmt.Errorf("user dengan id %d tidak ditemukan", idUser)
		}

		// Cek apakah user ini sudah punya alamat
		var count int64
		if err := db.DB.Model(&models.Alamat{}).Where("id_user = ?", idUser).Count(&count).Error; err != nil {
			return nil, fmt.Errorf("gagal menghitung jumlah alamat")
		}

		// Jika belum ada alamat, set default true
		isDefault := false
		if count == 0 {
			isDefault = true
		}

		alamat := models.Alamat{
			IDUser:      uint(idUser),
			Alamat:      getString(p, "alamat"),
			NamaA:       getString(p, "namaA"),
			TeleponA:    getString(p, "teleponA"),
			AlamatUtama: isDefault, // ✅ auto set default
		}

		if err := db.DB.Create(&alamat).Error; err != nil {
			return nil, fmt.Errorf("gagal membuat alamat: %v", err)
		}

		if err := db.DB.Preload("User").First(&alamat, alamat.IDAlamat).Error; err != nil {
			return nil, fmt.Errorf("gagal mengambil data alamat dengan relasi user")
		}

		return alamat, nil
	},
}

var UpdateAlamat = &graphql.Field{
	Type: types.AlamatType,
	Args: graphql.FieldConfigArgument{
		"id_alamat": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"alamat":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"namaA":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"teleponA":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := getInt(p, "id_alamat")
		var alamat models.Alamat
		if err := db.DB.First(&alamat, id).Error; err != nil {
			return nil, fmt.Errorf("alamat dengan id %d tidak ditemukan", id)
		}
		alamat.Alamat = getString(p, "alamat")
		alamat.NamaA = getString(p, "namaA")
		alamat.TeleponA = getString(p, "teleponA")
		if err := db.DB.Save(&alamat).Error; err != nil {
			return nil, fmt.Errorf("gagal mengupdate alamat: %v", err)
		}
		return alamat, nil
	},
}

var DeleteAlamat = &graphql.Field{
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name: "DeleteAlamatResponse",
		Fields: graphql.Fields{
			"message": &graphql.Field{Type: graphql.String},
		},
	}),
	Args: graphql.FieldConfigArgument{
		"id_alamat": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := getInt(p, "id_alamat")
		var alamat models.Alamat
		if err := db.DB.First(&alamat, id).Error; err != nil {
			return map[string]interface{}{"message": "Alamat tidak ditemukan"}, nil
		}
		if err := db.DB.Delete(&alamat).Error; err != nil {
			return map[string]interface{}{"message": "Gagal menghapus alamat"}, nil
		}
		return map[string]interface{}{"message": "Alamat berhasil dihapus"}, nil
	},
}

var AlamatUtama = &graphql.Field{
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name: "AlamatUtamaResponse",
		Fields: graphql.Fields{
			"message": &graphql.Field{Type: graphql.String},
		},
	}),
	Args: graphql.FieldConfigArgument{
		"id_alamat": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"id_user":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idAlamat := p.Args["id_alamat"].(int)
		idUser := p.Args["id_user"].(int)

		// reset semua default
		if err := db.DB.Model(&models.Alamat{}).Where("id_user = ?", idUser).Update("alamat_utama", false).Error; err != nil {
			return nil, err
		}

		// set default baru
		if err := db.DB.Model(&models.Alamat{}).Where("id_alamat = ?", idAlamat).Update("alamat_utama", true).Error; err != nil {
			return nil, err
		}

		return map[string]interface{}{"message": "Alamat default berhasil diubah"}, nil
	},
}
