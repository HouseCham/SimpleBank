package main

import (
	"database/sql"
	"log"

	"github.com/HouseCham/SimpleBank/api"
	db "github.com/HouseCham/SimpleBank/db/sqlc"
	_ "github.com/lib/pq"
)


const (
	serverAddress = "0.0.0.0:8080"
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("cannot connect to the database:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("unable to start the server:",err)
	}
}