package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	infoURL = "http://www.poatransporte.com.br/php/facades/process.php?a=il&p="
)

func errorResponse(w http.ResponseWriter, info string) {
	w.WriteHeader(http.StatusFailedDependency)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, info)
}

func chamarItinerarioAPI(idlinha string) ([]byte, error) {
	datapoaClient := http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := datapoaClient.Get(infoURL + idlinha)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func jsonItinerarioDecode(jsonPayload []byte) (*Itinerario, error) {
	itinerario := &Itinerario{}
	if err := json.Unmarshal(jsonPayload, &itinerario); err != nil {
		return nil, fmt.Errorf("Slice itinerarios %s", err)
	}

	if err := json.Unmarshal(jsonPayload, &itinerario.Pontos); err != nil {
		return nil, fmt.Errorf("iterator pontos %s", err)
	}
	//Limpando unmarshaled campos da interface já carregados
	delete(itinerario.Pontos, "ID")
	delete(itinerario.Pontos, "Codigo")
	delete(itinerario.Pontos, "Nome")

	return itinerario, nil
}

func itinerarioResponse(w http.ResponseWriter, itinerario *Itinerario) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	jsonPayload, err := json.Marshal(itinerario)
	if err != nil {
		errorResponse(w, fmt.Sprintf("Falha montar response %s", err))
		return
	}
	w.Write(jsonPayload)
}

func carregaItinerario(idlinha string) (*Itinerario, error) {
	jsonPayload, err := chamarItinerarioAPI(idlinha)
	if err != nil {
		return nil, fmt.Errorf("request %s", err)
	}
	itinerario, err := jsonItinerarioDecode(jsonPayload)
	if err != nil {
		return nil, fmt.Errorf("Falha payload %s", err)
	}
	return itinerario, nil
}

func itinerariosHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	itinerario, err := carregaItinerario(vars["id"])
	if err != nil {
		errorResponse(w, fmt.Sprintf("Falha request %s", err))
		return
	}
	itinerarioResponse(w, itinerario)
}

func okHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "docs https://github.com/airtonGit/datapoa")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/linhas", linhasHandler)
	r.Path("/linhas/").Queries("nome", "{nome}").HandlerFunc(linhasPesquisaNomeHandler)
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
