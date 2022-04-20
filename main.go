package main

import (
	"log"
	"net/http"
	"os"

	"./handlers"
)

func main() {
	// create log
	log := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create handlers
	hh := handlers.NewHello(log)
	gh := handlers.NewGoodbye(log)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	sm.Handle("/goodbye", gh)

	// listen for http requests
	http.ListenAndServe(":9090", nil)
}
