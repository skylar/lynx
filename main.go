package main

import (
	"os"

	"lynx/internal/server"
)

func main() {
  config := server.DefaultConfiguration()
	port := os.Getenv("PORT")
	if len(port) > 0 {
	  config.Port = port
	}

  server := server.NewLynxServer(config)
  server.Start()
}
