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
	links := service.Handle()
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	store := db.NewStore(conn)
	service.CheckDataAndInsert(store, links)
	service.Routes(store)
}

func ginserver() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ???", err)
	}

	store := db.NewStore(conn)
	datas := service.Handle()
	service.CheckDataAndInsert(store, datas)

	if err != nil {
		log.Fatal("cannot start server :", err)
	}

}
