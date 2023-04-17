package gapi

import (
	"fmt"

	db "github.com/HouseCham/SimpleBank/db/sqlc"
	"github.com/HouseCham/SimpleBank/pb"
	"github.com/HouseCham/SimpleBank/token"
	"github.com/HouseCham/SimpleBank/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer // This is the generated interface from the protobuf file.
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
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
	return server, nil
}