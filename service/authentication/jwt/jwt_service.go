package jwt

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secret_key = []byte("34C427392FE57CFCDC1B2FC395627")

type UserClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

const ContextClaimsKey = "claims"

func CreateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secret_key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := extractBearer(c.GetHeader("Authorization"))
		if err != nil {
			abortUnauthorized(c, err.Error())
			return
		}

		claims := &UserClaims{}
		token, err := jwt.ParseWithClaims(
			tokenStr,
			claims,
			func(t *jwt.Token) (any, error) {
				// Batasi algoritma agar tidak terkena alg-swap
				if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
					return nil, errors.New("algoritma JWT tidak valid")
				}
				return []byte(secret_key), nil
			},
			jwt.WithLeeway(30*time.Second), // toleransi clock skew
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
			// (opsional) jwt.WithIssuer("my-service"), jwt.WithAudience("my-web"),
		)
		if err != nil {
			log.Println("parse token error:", err)
			abortUnauthorized(c, "token tidak valid")
			return
		}
		if !token.Valid {
			abortUnauthorized(c, "token tidak valid")
			return
		}

		if claims.ExpiresAt == nil || time.Until(claims.ExpiresAt.Time) <= 0 {
			abortUnauthorized(c, "token expired")
			return
		}

		// simpan claims ke context agar handler lain bisa akses
		c.Set(ContextClaimsKey, claims)

		// lanjut ke handler berikutnya
		c.Next()
	}
}

// --- Helpers ---

func extractBearer(authorization string) (string, error) {
	if authorization == "" {
		return "", errors.New("authorization header kosong")
	}
	parts := strings.SplitN(authorization, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || parts[1] == "" {
		return "", errors.New("format Authorization harus 'Bearer <token>'")
	}
	return parts[1], nil
}

func abortUnauthorized(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error":   "unauthorized",
		"message": message,
	})
}
