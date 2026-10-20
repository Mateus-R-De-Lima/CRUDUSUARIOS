package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenPair struct {
	AccesToken   string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type CustomClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// getSecretKey busca a chave secreta das variáveis de ambiente
func getSecretKey() []byte {

	secret := os.Getenv("GOCRUD_JWT_SECRET")

	return []byte(secret)

}

func GenareteTokenPair(userID, email string) (*TokenPair, error) {
	secret := getSecretKey()

	//1 - Criar Acess token

	accessClaims := CustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "API-CRUD-USER",
		},
	}

	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	acessToken, err := accessTokenObj.SignedString(secret)

	if err != nil {
		return nil, err
	}

	// 2. Criar Refresh Token

	refreshClaims := accessClaims

	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	refreshToken, err := refreshTokenObj.SignedString(secret)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccesToken:   acessToken,
		RefreshToken: refreshToken,
	}, nil

}
