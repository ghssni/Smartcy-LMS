package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type JWTCustomClaims struct {
	UserID   uint32 `json:"userID"`
	Name     string `json:"name"`
	UserRole string `json:"userRole"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint32, name, userRole string) (string, time.Time, error) {
	expireDuration := viper.GetInt("JWT_EXPIRE")
	expiresAt := time.Now().Add(time.Hour * time.Duration(expireDuration))

	claims := &JWTCustomClaims{
		UserID:   userID,
		Name:     name,
		UserRole: userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as a response
	signedString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", time.Time{}, err
	}

	return signedString, expiresAt, nil
}
