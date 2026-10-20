package auth

import (
	"errors"
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

func ValidateToken(tokenString string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (any, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Erro na assinatura do Token:")
		}
		return getSecretKey(), nil

	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok {
		return nil, errors.New("token inválido ou expirado")
	}

	return claims, nil
}
