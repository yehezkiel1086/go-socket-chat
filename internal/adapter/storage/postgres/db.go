package postgres

import (
	"context"
	"fmt"
	"go-socket/internal/adapter/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func InitDB(ctx context.Context, conf *config.DB) (*DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", conf.Host, conf.User, conf.Password, conf.Name, conf.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &DB{}, err
	}

	return &DB{db: db}, nil
}

func (d *DB) MigrateDB(dbs ...any) error {
	return d.db.AutoMigrate(dbs...)
}

func (d *DB) GetDB() *gorm.DB {
	return d.db
}
