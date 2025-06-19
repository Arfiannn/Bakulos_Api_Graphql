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

		// === PRODUCT ===
		"createProduct":     CreateProduct,
		"updateProduct":     UpdateProduct,
		"deleteProduct":     DeleteProduct,
		"updateProductStok": UpdateProductStok,

		// === FAVORITE ===
		"createFavorite": CreateFavorite,
		"updateFavorite": UpdateFavorite,
		"deleteFavorite": DeleteFavorite,

		// === CHECKOUT ===
		"createCheckout":  CreateCheckout,
		"updateCheckout":  UpdateCheckout,
		"deleteCheckout":  DeleteCheckout,
		"confirmCheckout": ConfirmCheckout,

		// === KERANJANG ===
		"createKeranjang": CreateKeranjang,
		"updateKeranjang": UpdateKeranjang,
		"deleteKeranjang": DeleteKeranjang,

		// === CHAT ===
		"markChatAsRead": MarkChatAsRead,

		// === LOGIN ===
		"login":          Login,
		"forgetPassword": ForgetPassword,
	},
})
