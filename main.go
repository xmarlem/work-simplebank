package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"work-simplebank/api"
	db "work-simplebank/db/sqlc"
	"work-simplebank/util"
)

func main() {

	conf, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	fmt.Println("Conf", conf)

	conn, err := sql.Open(conf.DBDriver, conf.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(*conf, store)
	if err != nil {
		log.Fatalf("cannot create the server (%v)", err)
	}

	err = server.Start(conf.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
