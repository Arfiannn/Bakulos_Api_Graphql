package mutation

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"bakulos_grapghql/routes/types"
	"fmt"

	"github.com/graphql-go/graphql"
)

var CreateKeranjang = &graphql.Field{
	Type: types.KeranjangType,
	Args: graphql.FieldConfigArgument{
		"id_product": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"id_user":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"jumlah":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"sizeK":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idProduct := getInt(p, "id_product")
		idUser := getInt(p, "id_user")
		jumlah := getInt(p, "jumlah")
		sizeK := getString(p, "sizeK")

		var existingKeranjang models.Keranjang
		err := db.DB.Where("id_user = ? AND id_product = ? AND sizek = ?", idUser, idProduct, sizeK).First(&existingKeranjang).Error

		if err == nil {
			existingKeranjang.Jumlah += jumlah
			if err := db.DB.Save(&existingKeranjang).Error; err != nil {
				return nil, fmt.Errorf("gagal update jumlah: %v", err)
			}
			db.DB.Preload("User").Preload("Product").First(&existingKeranjang, existingKeranjang.IDKeranjang)
			return existingKeranjang, nil
		}

		if err.Error() != "record not found" {
			return nil, fmt.Errorf("gagal cek keranjang: %v", err)
		}

		keranjang := models.Keranjang{
			IDProduct: uint(idProduct),
			IDUser:    uint(idUser),
			Jumlah:    jumlah,
			SizeK:     sizeK,
		}
		if err := db.DB.Create(&keranjang).Error; err != nil {
			return nil, err
		}
		db.DB.Preload("User").Preload("Product").First(&keranjang, keranjang.IDKeranjang)
		return keranjang, nil
	},
}

var UpdateKeranjang = &graphql.Field{
	Type: types.KeranjangType,
	Args: graphql.FieldConfigArgument{
		"id_keranjang": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"id_product":   &graphql.ArgumentConfig{Type: graphql.Int},
		"jumlah":       &graphql.ArgumentConfig{Type: graphql.Int},
		"sizeK":        &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idKeranjang := getInt(p, "id_keranjang")
		var keranjang models.Keranjang
		if err := db.DB.First(&keranjang, idKeranjang).Error; err != nil {
			return nil, fmt.Errorf("keranjang dengan id %d tidak ditemukan", idKeranjang)
		}
		updates := map[string]interface{}{}
		if v := getInt(p, "id_product"); v != 0 {
			var product models.Product
			if err := db.DB.First(&product, v).Error; err != nil {
				return nil, fmt.Errorf("produk dengan id %d tidak ditemukan", v)
			}
			updates["id_product"] = v
		}
		if v := getInt(p, "jumlah"); v != 0 {
			updates["jumlah"] = v
		}
		if v, ok := p.Args["sizeK"].(string); ok && v != "" {
			updates["sizeK"] = v
		}
		if len(updates) == 0 {
			return nil, fmt.Errorf("tidak ada field yang diperbarui")
		}
		if err := db.DB.Model(&keranjang).Updates(updates).Error; err != nil {
			return nil, fmt.Errorf("gagal mengupdate keranjang: %v", err)
		}
		if err := db.DB.Preload("User").Preload("Product").First(&keranjang, idKeranjang).Error; err != nil {
			return nil, fmt.Errorf("gagal mengambil ulang keranjang")
		}
		return keranjang, nil
	},
}

var DeleteKeranjang = &graphql.Field{
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name: "DeleteKeranjangResponse",
		Fields: graphql.Fields{
			"message": &graphql.Field{Type: graphql.String},
		},
	}),
	Args: graphql.FieldConfigArgument{
		"id_keranjang": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := getInt(p, "id_keranjang")
		var keranjang models.Keranjang
		if err := db.DB.First(&keranjang, id).Error; err != nil {
			return map[string]interface{}{"message": "Keranjang tidak ditemukan"}, nil
		}
		if err := db.DB.Delete(&keranjang).Error; err != nil {
			return map[string]interface{}{"message": "Gagal menghapus keranjang"}, nil
		}
		return map[string]interface{}{"message": "Keranjang berhasil dihapus"}, nil
	},
}
