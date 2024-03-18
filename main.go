package main

import (
	"log"
)

func main() {
	database, err := NewPostgresDatabase()

	if err != nil {
		log.Fatal(err)
	}

	if err:= database.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":8080", database)
	server.Start()
}