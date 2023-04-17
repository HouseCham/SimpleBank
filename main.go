package main

import (
	"database/sql"
	"log"

	"github.com/HouseCham/SimpleBank/api"
	db "github.com/HouseCham/SimpleBank/db/sqlc"
	"github.com/HouseCham/SimpleBank/util"
	_ "github.com/lib/pq"
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
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("unable to start the server:",err)
	}
}