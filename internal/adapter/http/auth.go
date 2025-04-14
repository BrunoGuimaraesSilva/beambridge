package http

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/BrunoGuimaraesSilva/beambridge/internal/domain"
	"github.com/BrunoGuimaraesSilva/beambridge/internal/usecase/auth"
)

type AuthHandler struct {
	usecase *auth.Auth
}

func NewAuthHandler(usecase *auth.Auth) *AuthHandler {
	return &AuthHandler{usecase: usecase}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		domain.BadRequestError("Invalid request body", "Must be a valid JSON")
		return
	}

	if !isValidPassword(req.Password) {
		domain.BadRequestError("Invalid password", "Password must be at least 8 characters long and contain at least one number and one special character")
		return
	}

	err := h.usecase.Register(r.Context(), req.Email, req.Password)
	if err != nil {
		domain.BadRequestError("Registration failed", "An unexpected error occurred")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered"})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		domain.BadRequestError("Invalid request body", "Must be a valid JSON")
		return
	}

	token, err := h.usecase.Login(r.Context(), req.Email, req.Password)

	if err != nil {
		domain.InternalError("Login failed", "An unexpected error occurred")
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		domain.UnauthorizedError("Missing authorization header", "Must be a valid token")
		return
	}

	token := authHeader
	if len(authHeader) > 7 && authHeader[0:7] == "Bearer " {
		token = authHeader[7:]
	}

	newToken, err := h.usecase.RefreshToken(r.Context(), token)
	if err != nil {
		domain.UnauthorizedError("Invalid or expired token", "Must be a valid token")
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": newToken})
}

func isValidPassword(password string) bool {
	if len(password) < 8 {
		domain.UnauthorizedError("Password too short", "Must be at least 8 characters")
		return false
	}

	hasNumber, _ := regexp.MatchString(`[0-9]`, password)
	if !hasNumber {
		domain.UnauthorizedError("Password must contain at least one number", "Must contain at least one number")
		return false
	}

	hasSpecialChar, _ := regexp.MatchString(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};':"\\|,.<>\/?]+`, password)
	if !hasSpecialChar {
		domain.UnauthorizedError("Password must contain at least one special character", "Must contain at least one special character")
		return false
	}
	return true
}
