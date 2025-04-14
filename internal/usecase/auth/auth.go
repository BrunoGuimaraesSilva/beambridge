package auth

import (
	"context"
	"time"

	"github.com/BrunoGuimaraesSilva/beambridge/internal/domain"
	"github.com/BrunoGuimaraesSilva/beambridge/internal/port"
	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	userRepo port.UserRepository
	secret   string
}

func New(userRepo port.UserRepository, secret string) *Auth {
	return &Auth{userRepo: userRepo, secret: secret}
}

func (a *Auth) Register(ctx context.Context, email, password string) error {
	user, err := domain.NewUser(email, password)
	if err != nil {
		return err
	}
	return a.userRepo.SaveUser(ctx, user)
}

func (a *Auth) Login(ctx context.Context, email, password string) (string, error) {
	user, err := a.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if user == nil || !user.VerifyPassword(password) {
		return "", domain.BadRequestError("Invalid credentials", "Invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.String(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(a.secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a *Auth) RefreshToken(ctx context.Context, tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.BadRequestError("Invalid token", "Invalid token")
		}
		return []byte(a.secret), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	userID := token.Claims.(jwt.MapClaims)["sub"].(string)
	user, err := a.userRepo.FindUserByEmail(ctx, userID)
	if err != nil || user == nil {
		return "", domain.BadRequestError("Invalid credentials", "Invalid email or password")
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.String(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err = newToken.SignedString([]byte(a.secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
