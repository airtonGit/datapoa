package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

func carregaItinerario(idlinha string) (*LinhaItinerario, error) {
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

func pontoParser(pontos interface{}) (*LatLong, error) {
	pontoMap, ok := pontos.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("falha coordenadas map[string]interface{} assertion")
	}
	latlong := &LatLong{}
	lat, ok := pontoMap["lat"].(string)
	long, ok := pontoMap["lng"].(string)
	if !ok {
		return nil, fmt.Errorf("falha coordenas lat/lng assertion")
	}
	var err error
	latlong.Lat, err = strconv.ParseFloat(lat, 64)
	latlong.Long, err = strconv.ParseFloat(long, 64)
	if err != nil {
		return nil, fmt.Errorf("falha coordenas lat/lng string -> float")
	}
	return latlong, nil
}

func itinerarioParser(itinerario *Itinerario) ([]LatLong, error) {
	lista := []LatLong{}
	for _, item := range itinerario.Pontos {
		latlong, err := pontoParser(item)
		if err != nil {
			log.Println("itinerarioParser falha", item)
			continue
		}
		lista = append(lista, *latlong)
	}
	return lista, nil
}

func jsonItinerarioDecode(jsonPayload []byte) (*LinhaItinerario, error) {
	itinerario := &Itinerario{}
	if err := json.Unmarshal(jsonPayload, &itinerario); err != nil {
		return nil, fmt.Errorf("jsonItinerarioDecode: Slice itinerarios %s", err)
	}

	if err := json.Unmarshal(jsonPayload, &itinerario.Pontos); err != nil {
		return nil, fmt.Errorf("jsonItinerarioDecode: iterator pontos passo1 %s", err)
	}

	pontosItinerario, err := itinerarioParser(itinerario)
	if err != nil {
		return nil, fmt.Errorf("jsonItinerarioDecode: iterator pontos passo2 %s", err)
	}

	linhaItinerario := &LinhaItinerario{}
	linhaItinerario.ID = itinerario.ID
	linhaItinerario.Codigo = itinerario.Codigo
	linhaItinerario.Nome = itinerario.Nome
	linhaItinerario.Itinerario = pontosItinerario

	// //Limpando unmarshaled campos da interface já carregados
	// delete(itinerario.Pontos, "idlinha")
	// delete(itinerario.Pontos, "codigo")
	// delete(itinerario.Pontos, "nome")

	return linhaItinerario, nil
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

func itinerarioResponse(w http.ResponseWriter, itinerario *LinhaItinerario) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	jsonPayload, err := json.Marshal(itinerario)
	if err != nil {
		errorResponse(w, fmt.Sprintf("Falha montar response %s", err))
		return
	}
	w.Write(jsonPayload)
}
