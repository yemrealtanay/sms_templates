package main

import (
	"database/sql"
	"github.com/yemrealtanay/sms_templates/api"
	db "github.com/yemrealtanay/sms_templates/db/sqlc"
	"github.com/yemrealtanay/sms_templates/util"
	"log"
)

func main() {
	config, err := util.LoadConfig("../config/")
	if err != nil {
		log.Fatal("cannot find config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
