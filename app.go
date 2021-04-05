package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	err := http.ListenAndServe(":3000", loggedRouter)
	if err != nil {
		log.Panic(err.Error())
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Travix test sample web server")
}