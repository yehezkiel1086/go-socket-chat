package main

import (
	"fmt"
	"go-socket/internal/adapter/config"
)

func main() {
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Config imported successfully ✅")

	fmt.Println(conf.App)
}
