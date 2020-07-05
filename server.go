package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func errorResponse(w http.ResponseWriter, info string) {
	w.WriteHeader(http.StatusFailedDependency)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, info)
}

func okHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "docs https://github.com/airtonGit/datapoa")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/linhas", linhasHandler)
	r.Path("/linhas/").Queries("nome", "{nome}").HandlerFunc(linhasPesquisaNomeHandler)
	r.Path("/linhas/").Queries("raio", "{raio}", "lat", "{lat}", "lng", "{lng}").HandlerFunc(linhasPesquisaRaioHandler)
	r.HandleFunc("/linha/{id}", itinerariosHandler)
	r.HandleFunc("/ws", serveWs)
	r.HandleFunc("/", okHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
