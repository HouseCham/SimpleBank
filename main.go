package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/HouseCham/SimpleBank/api"
	db "github.com/HouseCham/SimpleBank/db/sqlc"
	"github.com/HouseCham/SimpleBank/gapi"
	"github.com/HouseCham/SimpleBank/pb"
	"github.com/HouseCham/SimpleBank/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config file:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("cannot connect to the database:", err)
	}
	store := db.NewStore(conn)

	//runGinServer(config, store)
	runGrpcServer(config, store)
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("unable to start the server:", err)
	}
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot listen to the port:", err)
	}

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	log.Printf("server is listening at %s", config.GRPCServerAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server:", err)
	}
}
