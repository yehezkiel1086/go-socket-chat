package repository

import (
	"context"
	"go-socket/internal/adapter/storage/postgres"
	"go-socket/internal/core/domain"
)

type UserRepository struct {
	db *postgres.DB
}

func InitUserRepository(db *postgres.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	db := ur.db.GetDB()
	if err := db.Create(&user).Error; err != nil {
		return &domain.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user *domain.User

	db := ur.db.GetDB()
	if err := db.Model(&domain.User{}).First(&user, "username = ?", username).Error; err != nil {
		return &domain.User{}, err
	}

	return user, nil
}
