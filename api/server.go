package api

import (
	"github.com/gin-gonic/gin"

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

	// v1alpha1 := server.router.Group("/v1alpha1")
	// {
	//  v1alpha1.POST("/accounts", server.createAccount)
	//  v1alpha1.GET("/accounts/:id", server.getAccount)
	//  v1alpha1.GET("/accounts", server.listAccount)
	// }

	server.router.POST("/accounts", server.createAccount)
	server.router.GET("/accounts/:id", server.getAccount)
	server.router.GET("/accounts", server.listAccount)
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
