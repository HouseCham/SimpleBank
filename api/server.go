package api

import (
	db "github.com/HouseCham/SimpleBank/db/sqlc"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", isValidCurrency)
	}

	// add routes to server
	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts", server.ListAccount)
	router.GET("/accounts/:id", server.GetAccount)
	router.POST("/transfers", server.CreateTransfer)

	server.router = router
	return server
}

// Start runs the HTTP server on a  specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
} 