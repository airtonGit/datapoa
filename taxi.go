package main

import (
	"encoding/csv"
	"os"
)

func carregarPontosTaxiArquivo() ([][]string, error) {
	arquivo, err := os.OpenFile("dados/pontos.csv", os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()
	csvReader := csv.NewReader(arquivo)
	csvReader.Comma = rune('#')
	return csvReader.ReadAll()
}

func carregarPontosTaxi(registros [][]string) error {
	const (
		nome int = iota
		lat
		long
		datahora
	)
	registros = registros[1:]
	for _, item := range registros {
		ponto := PontoTaxi{item[nome], item[lat], item[long], item[datahora]}
		linhasCache.PontosTaxi = append(linhasCache.PontosTaxi, ponto)
	}
	return nil
}

func storePontosTaxi() error {
	arquivo, err := os.OpenFile("dados/pontos.csv", os.O_RDWR|os.O_SYNC, 0755)
	if err != nil {
		return err
	}
	defer arquivo.Close()
	csvWriter := csv.NewWriter(arquivo)
	for _, item := range linhasCache.PontosTaxi {
		campos := []string{item.nome, item.Latitude, item.Longitude, item.Datahora}
		csvWriter.Write(campos)
	}
	return nil
}
