package main

//[{"id":"5529","codigo":"250-1","nome":"1 DE MAIO"},

// Linha campos de linha de onibus
type Linha struct {
	ID     string `json:"id"`
	Codigo string `json:"codigo"`
	Nome   string `json:"nome"`
}

//LinhaItinerario linha com seu itinerario
type LinhaItinerario struct {
	Linha
	Itinerario []LatLong
}

//LatLong par de coordenadas
type LatLong struct {
	Lat  float64
	Long float64
}

//{"idlinha":"5566","nome":"VILA
//NOVA","codigo":"266-1",
//"0":{"lat":"-30.12419057422600000",
//"lng":"-51.22378313620700000"},
//"1":{"lat":"-30.12410057422600000",
//"lng":"-51.22352313620700000"},"2":{"lat":"-30.12373357422600000","lng":"-51.22265713620700000"},"3":{"lat":"-30.12305757422600000","lng":"-51.22116713620700000"},"4":{"lat":"-30.12301857422600000","lng":"-51.22119413620700000"},"5":{"lat":"-30.12262857422600000","lng":"-51.22032813620700000"},"6":{"lat":"-30.12223457422600000","lng":"-51.21949313620700000"},"7":{"lat":"-30.12161657422600000","lng":"-51.21815113620700000"},"8":{"lat":"-30.12123957422600000","lng":"-51.21735113620700000"},"9":{"lat":"-30.12094357422600000","lng":"-51.21668113620700000"},"10":{"lat":"-30.12087657422600000","lng":"-51.21652913620700000"},"11":{"lat":"-30.12084357422600000","lng":"-51.21638313620700000"},"

//Itinerario item da lista
type Itinerario struct {
	ID     string                 `json:"idlinha"`
	Nome   string                 `json:"nome"`
	Codigo string                 `json:"codigo"`
	Pontos map[string]interface{} `json:"-"`
}

//PontoTaxi informações basicas
type PontoTaxi struct {
	nome      string
	Latitude  string
	Longitude string
	Datahora  string
}
