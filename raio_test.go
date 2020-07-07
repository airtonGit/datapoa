package main

import (
	"fmt"
	"testing"
)

func TestCarregarItinerarioArquivo(t *testing.T) {
	got, err := carregarItinerarioArquivo("5885")
	if err != nil {
		fmt.Println("got", got)
		t.Fatalf("Falhou %s", err)
	}
}

func TestCarregarLinhasItinerarios(t *testing.T) {
	err := carregarLinhasItinerarios()
	if err != nil {
		t.Fatalf("Falhou %s", err)
	}
	t.Fatal()
}

//"-30.01039857422600000","lng":"-51.14512013620700000"}

func TestRaio(t *testing.T) {
	_, err := carregarLinhasCached()
	if err != nil {
		t.Fatalf("Falhou %s", err)
	}
	got, err := raio(-30.010398574226, -51.145120136207, 1)
	if err != nil {
		fmt.Println("got", got)
		t.Fatalf("Falhou %s len %d", err, len(got))
	}
	t.Fatalf("Falhou %s len %d", err, len(got))
}
