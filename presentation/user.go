package presentation

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

// should be stored at GSM/env file
var JwtKey = []byte("my_secret_key")

type IUser interface {
	GetUserDataByUserName(ctx context.Context, username string) (User, error)
}

type IUserUC interface {
	DoLogin(ctx context.Context, req LoginRequest) (string, error)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
