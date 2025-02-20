package gapi

import (
	"fmt"
	"simpleBank/authtoken"
	db "simpleBank/db/sqlc"
	"simpleBank/pb"
	"simpleBank/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker authtoken.Maker
}

// NewServer creates a new gRPC server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := authtoken.NewPasetoMaker(config.TokenSymmetricKey)
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
