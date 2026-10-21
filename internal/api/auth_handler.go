package api

import (
	"CRUDUSERS/internal/auth"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// LoginHandler autentica o usuário e devolve o par de tokens
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Payload inválido", http.StatusBadRequest)
		return
	}

	// Validação fictícia para testes (Substitua pela busca no banco de dados)
	if req.Email != "admin@email.com" || req.Password != "123456" {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	tokens, err := auth.GenareteTokenPair("user-123", req.Email)
	if err != nil {
		http.Error(w, "Erro ao gerar tokens", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)
}

// RefreshHandler recebe um Refresh Token e gera um novo Access Token
func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	var req RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Payload inválido", http.StatusBadRequest)
		return
	}

	claims, err := auth.ValidateToken(req.RefreshToken)
	if err != nil {
		http.Error(w, "Refresh token inválido ou expirado", http.StatusUnauthorized)
		return
	}

	tokens, err := auth.GenareteTokenPair(claims.UserID, claims.Email)
	if err != nil {
		http.Error(w, "Erro ao renovar token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)
}
