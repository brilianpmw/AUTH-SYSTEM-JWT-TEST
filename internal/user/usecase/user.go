package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/brilianpmw/linknau/presentation"
)

func (uc *Usecase) DoLogin(ctx context.Context, req presentation.LoginRequest) (string, error) {
	user, err := uc.repository.User.GetUserDataByUserName(ctx, req.Username)
	if err != nil || user.Password != req.Password {
		return "", errors.New("unauthorized")
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &presentation.Claims{
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(presentation.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
