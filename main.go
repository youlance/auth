package main

import (
	"log"

	server "github.com/youlance/auth/api/http"
	"github.com/youlance/auth/db"
	"github.com/youlance/auth/pkg/config"
)

func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := db.New(config)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	server, err := server.NewServer(config, db)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
