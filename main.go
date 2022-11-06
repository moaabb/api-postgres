package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/moaabb/api-postgres/config"
	"github.com/moaabb/api-postgres/driver"
)

var app config.Application

func main() {
	cfg := config.ReadConfig()

	app.L = hclog.Default()
	conn, err := driver.ConnectDB(cfg.DB.DSN)
	if err != nil {
		app.L.Error("Could not connect to DB", err.Error())
		os.Exit(1)
	}
	app.L.Info("Connected to DB!")

	app.DBModel = conn

	s := http.Server{
		Addr:    cfg.Server.Address,
		Handler: routes(),
	}

	go func() {
		log.Println(fmt.Sprintf("Server Listening on pont %s", cfg.Server.Address))

		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err.Error())
		}

	}()

	// Create a channel  to receive the termination signals
	c := make(chan os.Signal, 1)

	// Using the Notify method to send the determined signals to the channel created
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Sig awaits to receive the signal from the channel when it's available
	sig := <-c

	log.Println(fmt.Sprintf("Got Signal: %s", sig))

	// Try to gracefully shutdown the server, or force the shutdown in 30 seconds
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
