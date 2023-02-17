package main

import (
	db "cpic/db/sqlc"
	service "cpic/servcie"
	"cpic/util"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ???", err)
	}

	store := db.NewStore(conn)
	datas := service.NewDBServer(store)
	service.CheckDataAndInsert(store, datas)
}

func ginserver() {
	//service.Routes()
	//service.WebSeseav()
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ???", err)
	}

	store := db.NewStore(conn)
	server := service.NewServer(store)
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server :", err)
	}

}
