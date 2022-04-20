package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/geoff-maddock/golang-microservice/handlers"
)

func main() {
	// create log
	log := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create handlers
	// hh := handlers.NewHello(log)
	// gh := handlers.NewGoodbye(log)
	ph := handlers.NewProducts(log)

	sm := http.NewServeMux()

	// handle routes
	// sm.Handle("/hello", hh)
	// sm.Handle("/goodbye", gh)
	sm.Handle("/", ph)

	// configure our server instance
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// listen for http requests
	//s.ListenAndServe()

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	//signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	log.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
