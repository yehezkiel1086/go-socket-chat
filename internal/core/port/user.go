package port

import (
	"context"
	"go-socket/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
}
