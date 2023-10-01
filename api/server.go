package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	db "work-simplebank/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store:  store,
		router: gin.Default(),
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
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
