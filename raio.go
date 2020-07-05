package main

import (
	"fmt"

	"github.com/kwahome/go-haversine/pkg/haversine"
)

func carregarLinhasItinerarios() error {
	linhas, err := carregarLinhasCached()
	if err != nil {
		return fmt.Errorf("carregar itinerarios %s", err)
	}
	linhasCache.LinhaMap = make(map[string]LinhaItinerario, len(linhas))
	for _, item := range linhas {
		linhaItinerario := LinhaItinerario{item, []LatLong{}}

		linhasCache.LinhaMap[item.ID] = linhaItinerario
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
