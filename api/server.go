package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	db "work-simplebank/db/sqlc"
	"work-simplebank/token"
	"work-simplebank/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey) // empty string as placeholder
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %v", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		router:     gin.Default(),
	}

	// con binding.Validator.Engine() ottengo il validator che gin sta usando correntemente
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	// v1alpha1 := server.router.Group("/v1alpha1")
	// {
	//  v1alpha1.POST("/accounts", server.createAccount)
	//  v1alpha1.GET("/accounts/:id", server.getAccount)
	//  v1alpha1.GET("/accounts", server.listAccount)
	// }

	server.router.POST("/accounts", server.createAccount)
	server.router.GET("/accounts/:id", server.getAccount)
	server.router.GET("/accounts", server.listAccount)
	server.router.POST("/transfers", server.createTransfer)
	server.router.POST("/users", server.createUser)
	server.router.POST("/users/login", server.loginUser)
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
