package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rahulvarma07/goo_backend/internal/config"
)

func main() {
	// TODO:
	// load the config file here
	// setup the router using http packagte
	// setup the server
	// close the server once task is done gracefully

	cnf := config.MustLoadConfig() // COMPLETED: (loading the config file)

	router := http.NewServeMux() // COMPLETED: (Setup router)

	// Testing by writing an handler
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Just testing port"))
	})

	// setting up server with Addr in config file
	server := http.Server{
		Addr:    cnf.Port,
		Handler: router,
	}
	// COMPLETED: (Seting up server)

	// TODO:
	// shut the server gracefully
	// by getting the commands from the os

	channel := make(chan os.Signal, 1) // this is to check for signal such as interupt

	signal.Notify(channel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		slog.Info("Server has started %s", cnf.Port)
		err := server.ListenAndServe() // server listening
		if err != nil {
			log.Fatalf("Error in starting the server %s", err) // if err
		}
	}()
	<-channel

	slog.Info("Shutting down the server")

	// now try to shutdown the server once interupted

	con, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(con)
	if err != nil {
		slog.Error("There is an error in shutting down", slog.String("error", err.Error()))
	}
	slog.Info("Server shutdown successfully")
}
