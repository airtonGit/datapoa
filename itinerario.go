package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	infoURL = "http://www.poatransporte.com.br/php/facades/process.php?a=il&p="
)

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

func carregaItinerario(idlinha string) (*Itinerario, error) {
	jsonPayload, err := chamarItinerarioAPI(idlinha)
	if err != nil {
		return nil, fmt.Errorf("request %s", err)
	}
	itinerario, err := jsonItinerarioDecode(jsonPayload)
	if err != nil {
		return nil, fmt.Errorf("Falha decode payload %s", err)
	}
	return itinerario, nil
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
	delete(itinerario.Pontos, "idlinha")
	delete(itinerario.Pontos, "codigo")
	delete(itinerario.Pontos, "nome")

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
