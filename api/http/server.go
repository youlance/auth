package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/youlance/auth/db"
	"github.com/youlance/auth/pkg/config"
	"github.com/youlance/auth/pkg/token"
)

type Server struct {
	config     config.Config
	db         *db.DB
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config config.Config, db *db.DB) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		db:         db,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/login", server.loginUser)
	router.POST("/auth", server.verify)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
