package main

import (
	handler "Prisma/api"
	c "Prisma/utils/config"
	"log"
	"net/http"
)

func main() {
	c.Config()
	http.HandleFunc("/identify", handler.Identify)
	log.Fatal(http.ListenAndServe(c.Configuration.Server.Port, nil))
}
