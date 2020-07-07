package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/kwahome/go-haversine/pkg/haversine"
)

func carregarItinerarioArquivo(id string) (*LinhaItinerario, error) {
	arquivoJSON, err := os.Open(fmt.Sprintf("json/itinerarios-%s.json", id))
	if err != nil {
		return nil, fmt.Errorf("falha abrir arquivo %s", err)
	}
	jsonPayload, err := ioutil.ReadAll(arquivoJSON)
	if err != nil {
		return nil, fmt.Errorf("falha ler arquivo %s", err)
	}
	itinerario, err := jsonItinerarioDecode(jsonPayload)
	if err != nil {
		return nil, fmt.Errorf("Falha decode payload %s", err)
	}
	return itinerario, nil
}

func carregarLinhasItinerarios() error {
	log.Println("Carregando itinerarios cached arquivos...")
	linhasCache.LinhaMap = make(map[string]LinhaItinerario, len(linhasCache.Linhas))
	for _, item := range linhasCache.Linhas {
		linhaItinerario, err := carregarItinerarioArquivo(item.ID)
		if err != nil {
			log.Println(fmt.Sprintf("carregerLinhasItinerarios linha %s itnerarios %s", item.Nome, err))
			continue
		}
		linhasCache.LinhaMap[item.ID] = *linhaItinerario
	}
	fmt.Println("Carregando itinerarios pronto.")
	return nil
}

func raio(lat, long float64, raioKm int) ([]Linha, error) {
	pontoInterese := haversine.Coordinate{
		Latitude:  lat,
		Longitude: long,
	}
	dentroRaioLista := []Linha{}
	for _, linha := range linhasCache.LinhaMap {
		for _, itinerario := range linha.Itinerario {
			pontoItinerario := haversine.Coordinate{
				Latitude:  itinerario.Lat,
				Longitude: itinerario.Long,
			}
			raioMetros := haversine.Distance(1000 * float64(raioKm))
			distancia := pontoInterese.DistanceTo(pontoItinerario, haversine.M)
			if distancia < raioMetros {
				linhaProxima := &Linha{
					ID:     linha.ID,
					Codigo: linha.Codigo,
					Nome:   linha.Nome}
				dentroRaioLista = append(dentroRaioLista, *linhaProxima)
				fmt.Println("Itinerario no perimetro ", linha.Nome, distancia)
				break
			}
		}
	}
	return dentroRaioLista, nil
}

func pesquisaRaioAdapter(raiostr, lng, lat string) ([]Linha, error) {
	raioKm, err := strconv.Atoi(raiostr)
	latf, err := strconv.ParseFloat(lat, 64)
	lngf, err := strconv.ParseFloat(lng, 64)
	if err != nil {
		return nil, fmt.Errorf("Falha converterção parametros %s", err)
	}
	linhasProximas, err := raio(latf, lngf, raioKm)
	if err != nil {
		return nil, fmt.Errorf("Falha pesquisa proximos %s", err)
	}

	return linhasProximas, nil
}

func pesquisaRaioResponse(w http.ResponseWriter, resultado []Linha) {
	jsonResponse, err := json.Marshal(resultado)
	if err != nil {
		log.Println("raio falha encode resposta", err)
		errorResponse(w, "Raio falha encode resposta")
		return
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println("raio falha envio resposta", err)
		errorResponse(w, "Raio falha envio resposta")
		return
	}
}

func linhasPesquisaRaioHandler(w http.ResponseWriter, req *http.Request) {
	raioKm := req.FormValue("raio")
	lng := req.FormValue("lng")
	lat := req.FormValue("lat")
	log.Println(fmt.Sprintf("raio: filtro raio:%s lat:%s lng:%s", raioKm, lat, lng))
	// err := carregarLinhasItinerarios()
	// if err != nil {
	// 	log.Println("raio falha carregar itinerarios", err)
	// 	errorResponse(w, "Raio falha carregar itinerarios")
	// 	return
	// }
	linhasProximas, err := pesquisaRaioAdapter(raioKm, lng, lat)
	if err != nil {
		log.Println("raio falha pesquisa itinerarios", err)
		errorResponse(w, "Raio falha pesquisa itinerarios")
		return
	}
	pesquisaRaioResponse(w, linhasProximas)
}
