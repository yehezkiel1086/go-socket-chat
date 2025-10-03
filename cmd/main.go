package main

import (
	"context"
	"fmt"
	"go-socket/internal/adapter/config"
	"go-socket/internal/adapter/handler"
	"go-socket/internal/adapter/storage/postgres"
	"go-socket/internal/adapter/storage/postgres/repository"
	"go-socket/internal/core/domain"
	"go-socket/internal/core/service"
)

func main() {
	// init config
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Config imported successfully ✅")

	// init db
	ctx := context.Background()
	db, err := postgres.InitDB(ctx, conf.DB)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connected successfully ✅")	

	// migrate database
	if err := db.MigrateDB(&domain.User{}); err != nil {
		panic(err)
	}
	fmt.Println("DBs migrated successfully ✅")

	// dependency injection
	userRepo := repository.InitUserRepository(db)
	userSvc := service.InitUserService(userRepo)
	userHandler := handler.InitUserHandler(userSvc)

	authSvc := service.InitAuthService(userRepo)
	authHandler := handler.InitAuthHandler(authSvc)

	hubRepo := repository.InitHubRepository()
	hubSvc := service.InitHubService(hubRepo)
	hubHandler := handler.InitHubHandler(hubSvc)

	clientSvc := service.InitClientService(hubRepo)
	clientHandler := handler.InitClientHandler(clientSvc)

	// routing
	r, err := handler.InitRouter(
		conf.HTTP,
		*userHandler,
		*authHandler,
		*hubHandler,
		*clientHandler,
	)
	if err != nil {
		panic(err)
	}

	// start server
	uri := fmt.Sprintf("%v:%v", conf.HTTP.Host, conf.HTTP.Port)
	r.Serve(uri)
}
