package main

import (
  "database/sql"
  "log"

  _ "github.com/lib/pq"

  "work-simplebank/api"
  db "work-simplebank/db/sqlc"
)

const (
  dbDriver      = "postgres"
  dbSource      = "postgresql://root:root@localhost:5432/work_simplebank?sslmode=disable"
  serverAddress = "0.0.0.0:8081"
)

func main() {

  conn, err := sql.Open(dbDriver, dbSource)

  if err != nil {
    log.Fatal("cannot connect to db:", err)
  }

  store := db.NewStore(conn)
  server := api.NewServer(store)

  err = server.Start(serverAddress)
  if err != nil {
    log.Fatal("cannot start server")
  }
}
