package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kwahome/go-haversine/pkg/haversine"
)

func carregarLinhasItinerarios() error {
	linhas, err := carregarLinhasCached()
	if err != nil {
		return fmt.Errorf("carregar itinerarios %s", err)
	}
	linhasCache.LinhaMap = make(map[string]LinhaItinerario, len(linhas))
	for _, item := range linhas {
		linhaItinerario, err := carregaItinerario(item.ID)
		if err != nil {
			log.Println(fmt.Sprintf("carregerLinhasItinerarios linha %s itnerarios %s", item.Nome, err))
			continue
		}
		linhasCache.LinhaMap[item.ID] = *linhaItinerario
	}
	return nil
}

func raio(ponto1, ponto2 haversine.Coordinate) {
	nairobi := haversine.Coordinate{
		Latitude:  1.2921,
		Longitude: 36.8219,
	}
	mombasa := haversine.Coordinate{
		Latitude:  4.0435,
		Longitude: 39.6682,
	}
	units := haversine.M
	distance := nairobi.DistanceTo(mombasa, units)
	fmt.Println("Distance from Nairobi =", nairobi, "to Mombasa =", mombasa, "in", units, "is", distance)
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
