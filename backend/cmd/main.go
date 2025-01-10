package main

import (
	"backend/api/handler"
	"backend/api/server"
	"backend/api/storer"
	"backend/db"
	"backend/service"
	"log"
)

const portNum = ":8090"

func main() {
	db, err := db.New()
	if err != nil {
		log.Fatalf("error opening db: %v", err)
	}
	defer db.Close()
	log.Print("connected to db ~")

	src := service.New()

	str := storer.NewPostgreSQLStorer(db.GetDB())

	h := handler.New(str, src)

	srv := server.New(portNum, h)

	if err := srv.Run(); err != nil {
		log.Fatalln(err)
	}

}
