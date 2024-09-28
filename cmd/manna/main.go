package main

import (
	"log"

	"github.com/8ideaz/manna/internal/server"
)

func main() {
	log.Fatal(server.Run())
}
