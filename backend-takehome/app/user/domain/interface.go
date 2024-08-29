package domain

import (
	"app/common/snap/snapauth"
	"context"
)

type Service interface {
	Register(ctx context.Context, req RegisterRequest) (GeneralResponse, error)
	Login(ctx context.Context, req LoginRequest) (GeneralResponse, error)
}

type Repository interface {
	StoreUser(model User) error

	FindUserByEmail(email string) (User, error)
	//FindUserByUsernameAndPassword(ctx context.Context, username, password string) (*User, error)
}

type CacheRepository interface {
	StoreAccessToken(ctx context.Context, userID string, dt snapauth.AccessTokenResponse) error
	GetAccessToken(ctx context.Context, accessToken string) (bool, error)
}
