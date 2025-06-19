package auth

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = []byte("RAHASIA_JWT")

func GenerateJWT(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"id_user": userID,
		"role":    role,
		"exp":     time.Now().Add(10 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

func ValidateJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&payload)

	var user models.User
	if err := db.DB.Where("email = ?", payload.Email).First(&user).Error; err == nil {
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)) == nil {
			token, _ := GenerateJWT(user.IDUser, "user")
			json.NewEncoder(w).Encode(map[string]string{"token": token})
			return
		}
	}

	var penjual models.Penjual
	if err := db.DB.Where("email = ?", payload.Email).First(&penjual).Error; err == nil {
		if bcrypt.CompareHashAndPassword([]byte(penjual.Password), []byte(payload.Password)) == nil {
			token, _ := GenerateJWT(penjual.IDPenjual, "penjual")
			json.NewEncoder(w).Encode(map[string]string{"token": token})
			return
		}
	}

	http.Error(w, "Email atau password salah", http.StatusUnauthorized)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow GraphiQL GET (UI in browser)
		if r.Method == http.MethodGet && r.URL.Path == "/graphql" {
			next.ServeHTTP(w, r)
			return
		}

		if r.URL.Path == "/login" || r.URL.Path == "/ws" {
			next.ServeHTTP(w, r)
			return
		}

		if isPublicGraphQLRequest(r) {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := ValidateJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func isPublicGraphQLRequest(r *http.Request) bool {
	if r.Method != http.MethodPost || r.Header.Get("Content-Type") != "application/json" {
		return false
	}

	var bodyCopy struct {
		Query string `json:"query"`
	}

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(r.Body)

	r.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))
	_ = json.Unmarshal(buf.Bytes(), &bodyCopy)

	return strings.Contains(bodyCopy.Query, "login") ||
		strings.Contains(bodyCopy.Query, "createUser") ||
		strings.Contains(bodyCopy.Query, "createPenjual")
}

func GetUserClaims(p graphql.ResolveParams) (jwt.MapClaims, error) {
	claims := p.Context.Value("user")
	if claims == nil {
		return nil, fmt.Errorf("Unauthorized")
	}
	return claims.(jwt.MapClaims), nil
}
