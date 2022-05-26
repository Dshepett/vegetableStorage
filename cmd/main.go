package main

import (
	"log"
	"vegetableShop/internal/app"
	"vegetableShop/internal/config"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}
	a := app.Init(conf)
	a.Run()
}
