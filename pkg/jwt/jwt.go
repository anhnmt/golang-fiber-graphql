package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/xdorro/golang-fiber-base-project/graph/model"
	"github.com/xdorro/golang-fiber-base-project/internal/config"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateToken(email string) (*model.Token, error) {
	jwtConfig := config.Jwt()

	if jwtConfig.SecretKey == "" {
		return nil, errors.New("SecretKey is invalid")
	}

	tokenExpired := time.Now().Add(time.Minute * time.Duration(jwtConfig.AccessToken.ExpiredAt)).Unix()
	accessToken, err := GenerateJwtToken(email, jwtConfig.SecretKey, tokenExpired)
	if err != nil {
		return nil, err
	}

	refreshExpired := time.Now().Add(time.Minute * time.Duration(jwtConfig.RefreshToken.ExpiredAt)).Unix()
	refreshToken, err := GenerateJwtToken(email, jwtConfig.SecretKey, refreshExpired)
	if err != nil {
		return nil, err
	}

	return &model.Token{
		TokenType:      "Bearer",
		AccessToken:    accessToken,
		TokenExpired:   int(tokenExpired),
		RefreshToken:   refreshToken,
		RefreshExpired: int(refreshExpired),
	}, nil
}

// VerifyPassword compares raw password with it's hashed values
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJwtToken(email string, secretKey string, expiredAt int64) (token string, err error) {
	// Create token
	newToken := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := newToken.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["expired"] = expiredAt

	token, err = newToken.SignedString([]byte(secretKey))
	return
}
