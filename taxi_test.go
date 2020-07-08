package main

import (
	"testing"
)

func TestCarregarPontosTaxi(t *testing.T) {
	registros, err := carregarPontosTaxiArquivo()
	if err != nil {
		t.Fatalf("Falha carga arquivo %s", err)
	}
	err = carregarPontosTaxi(registros)
	if err != nil {
		t.Fatalf("Falha carga Pontos %s", err)
	}
}

func TestStorePontosTaxi(t *testing.T) {
	err := storePontosTaxi()
	if err != nil {
		t.Fatalf("Falha store arquivo %s", err)
	}
}
