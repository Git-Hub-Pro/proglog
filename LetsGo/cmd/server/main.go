package main

import (
	"log"

	"github.com/Git-Hub-Pro/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
