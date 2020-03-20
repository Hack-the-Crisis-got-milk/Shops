package main

import (
	"fmt"
	"log"
)

func main() {
	server, err := InitializeServer()
	if err != nil {
		log.Fatal(fmt.Sprintf("could not create server: %s", err))
	}

	err = server.Run(fmt.Sprintf(":%s", server.Port))
	if err != nil {
		log.Fatal(fmt.Sprintf("could not start server: %s", err))
	}

	log.Printf("server started at: localhost:%s", server.Port)
}
