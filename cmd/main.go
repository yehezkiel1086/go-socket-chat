package main

import (
	"fmt"
	"go-socket/internal/adapter/config"
	"go-socket/internal/adapter/storage/postgres"
)

func main() {
	// init config
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Config imported successfully ✅")

	// init db
	_, err = postgres.InitDB(conf.DB)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connected successfully ✅")	
}
