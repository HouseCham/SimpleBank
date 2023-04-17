package api

import (
	"fmt"

	db "github.com/HouseCham/SimpleBank/db/sqlc"
	"github.com/HouseCham/SimpleBank/token"
	"github.com/HouseCham/SimpleBank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPaseToMaker([]byte(config.TokenSymmetricKey))
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", isValidCurrency)
	}

	// add routes to server
	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts", server.ListAccount)
	router.GET("/accounts/:id", server.GetAccount)
	router.POST("/transfers", server.CreateTransfer)
	router.POST("/users", server.CreateUser)

	server.router = router
	return server, nil
}

// Start runs the HTTP server on a  specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
