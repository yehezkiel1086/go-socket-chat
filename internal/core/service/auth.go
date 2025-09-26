package service

import (
	"context"
	"errors"
	"go-socket/internal/core/domain"
	"go-socket/internal/core/port"
	"go-socket/internal/core/util"
)

type AuthService struct {
	repo port.UserRepository
}

func InitAuthService(repo port.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (as *AuthService) Login(ctx context.Context, req *domain.AuthRequest)(*domain.AuthResponse, error) {
	// check username
	user, err := as.repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return &domain.AuthResponse{}, errors.New("invalid credentials")
	}

	// check password
	if err := util.ComparePassword(user.Password, req.Password); err != nil {
		return &domain.AuthResponse{}, errors.New("invalid credentials")
	}

	// return response
	return &domain.AuthResponse{
		Username: user.Username,
	}, nil
}
