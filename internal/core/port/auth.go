package port

import (
	"context"
	"go-socket/internal/core/domain"
)

type AuthService interface {
	Login(ctx context.Context, req *domain.AuthRequest) (*domain.AuthResponse, error)
}
