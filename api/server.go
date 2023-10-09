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

	server.setupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.POST("/transfers", server.createTransfer)
	server.router = router
}
