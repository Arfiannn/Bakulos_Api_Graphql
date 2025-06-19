package mutation

import (
	"github.com/graphql-go/graphql"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		// === USER ===
		"createUser":       CreateUser,
		"updateUser":       UpdateUser,
		"updateUserProfil": UpdateUserProfil,

		// === PENJUAL ===
		"createPenjual":       CreatePenjual,
		"updatePenjual":       UpdatePenjual,
		"updatePenjualProfil": UpdatePenjualProfil,

		// === ALAMAT ===
		"createAlamat": CreateAlamat,
		"updateAlamat": UpdateAlamat,
		"deleteAlamat": DeleteAlamat,
		"alamatUtama":  AlamatUtama,
		
	},
})
