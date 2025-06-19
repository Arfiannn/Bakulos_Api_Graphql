package websocket

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type ChatMessage struct {
	IDUser    *uint  `json:"id_user"`
	IDPenjual *uint  `json:"id_penjual"`
	Chat      string `json:"chat"`
	Sender    string `json:"sender"`
	CreatedAt string `json:"created_at"` // kirim string (ISO8601), bukan time.Time langsung
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("❌ Gagal upgrade websocket:", err)
		return
	}
	defer conn.Close()

	log.Println("✅ WebSocket client connected!")

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("❌ Error membaca message:", err)
			break
		}

		log.Println("📥 Pesan diterima:", string(message))

		var chat ChatMessage
		if err := json.Unmarshal(message, &chat); err != nil {
			log.Println("❌ Error parsing JSON:", err)
			conn.WriteJSON(map[string]string{"error": "Format data JSON salah"})
			continue
		}

		if chat.IDUser == nil || chat.IDPenjual == nil || chat.Sender == "" || chat.Chat == "" {
			conn.WriteJSON(map[string]string{"error": "Data chat tidak lengkap"})
			continue
		}

		var user models.User
		if err := db.DB.First(&user, *chat.IDUser).Error; err != nil {
			log.Printf("❌ User dengan ID %d tidak ditemukan\n", *chat.IDUser)
			conn.WriteJSON(map[string]string{"error": "User tidak ditemukan"})
			continue
		}

		var penjual models.Penjual
		if err := db.DB.First(&penjual, *chat.IDPenjual).Error; err != nil {
			log.Printf("❌ Penjual dengan ID %d tidak ditemukan\n", *chat.IDPenjual)
			conn.WriteJSON(map[string]string{"error": "Penjual tidak ditemukan"})
			continue
		}

		now := time.Now()
		newChat := models.Chat{
			IDUser:    chat.IDUser,
			IDPenjual: chat.IDPenjual,
			Chat:      chat.Chat,
			Sender:    chat.Sender,
			IsRead:    false,
			CreatedAt: now,
		}

		if err := db.DB.Create(&newChat).Error; err != nil {
			log.Println("❌ Gagal simpan chat:", err)
			conn.WriteJSON(map[string]string{"error": "Gagal simpan chat"})
			continue
		}

		log.Printf("✅ Chat berhasil disimpan: %+v\n", newChat)

		conn.WriteJSON(map[string]interface{}{
			"message": "Chat berhasil disimpan",
			"data": ChatMessage{
				IDUser:    chat.IDUser,
				IDPenjual: chat.IDPenjual,
				Chat:      chat.Chat,
				Sender:    chat.Sender,
				CreatedAt: now.Format(time.RFC3339),
			},
		})
	}
}
