package main

import (
	"fmt"
	"log"
	"net/http"

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
	fmt.Printf("Server has started %s", cnf.Port)
	err := server.ListenAndServe() // server listening

	if err != nil {
		log.Fatalf("Error in starting the server %s", err) // if err
	}
}
