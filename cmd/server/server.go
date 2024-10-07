package server

import (
	"github.com/LeMinh0706/music-go/internal/router"
	"github.com/LeMinh0706/music-go/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Config util.Config
	Router *gin.Engine
}

func NewServer(config util.Config) (*Server, error) {
	server := &Server{
		Config: config,
		Router: gin.Default(),
	}
	router.NewRouter(server.Router)
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
