package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	linhas, err := carregarLinhasCached()
	if err != nil {
		return fmt.Errorf("carregar itinerarios %s", err)
	}
	linhasCache.LinhaMap = make(map[string]LinhaItinerario, len(linhas))
	for _, item := range linhas {
		fmt.Println("Carregando itinerario linha", item.Nome)
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
				fmt.Println("Linha com itinerario no perimetro ", linha.Nome, distancia)
				break
			}
		}
	}
	return dentroRaioLista, nil
}

func linhasPesquisaRaioHandler(w http.ResponseWriter, req *http.Request) {
	raio := req.FormValue("raio")
	lng := req.FormValue("lng")
	lat := req.FormValue("lat")
	log.Println(fmt.Sprintf("raio: filtro raio:%s lat:%s lng:%s", raio, lat, lng))
	err := carregarLinhasItinerarios()
	if err != nil {
		fmt.Fprintf(w, "Falha raio %s", err)
	}
}
