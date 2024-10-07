package main

import (
	"log"

	"github.com/LeMinh0706/music-go/cmd/server"
	"github.com/LeMinh0706/music-go/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	server, err := server.NewServer(config)
	if err != nil {
		log.Fatal("Cannot load server: ", err)
	}
	server.Start(config.ServerAddress)
}
