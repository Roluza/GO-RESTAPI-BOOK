package utils

import (
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "SecretKey"
// buat token
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

// Middleware untuk memvalidasi token JWT
func AuthMiddleware(c *fiber.Ctx) error {
	// Ambil token dari header Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing or invalid token",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse dan validasi token JWT
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Verifikasi algoritma token
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		log.Println("Error parsing token:", err) // Log error
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	// Token valid, lanjutkan ke handler berikutnya
	return c.Next()
}