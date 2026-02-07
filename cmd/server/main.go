package main

import (
	"log"

	"github.com/Yogi-1996/notes-backend/internal/config"
	"github.com/Yogi-1996/notes-backend/internal/servers"
)

func main() {
	cfg := config.Load()

	if err := servers.RunApp(cfg); err != nil {
		log.Fatal(err)
	}
}
