package main

import (
	"strings"
	"testing"
)

func TestListRequest(t *testing.T) {
	got, err := carregaLinhas()
	if err != nil || false == strings.Contains(string(got), "1 DE MAIO") {
		t.Fatal("Falhei", err, string(got))
	}
}

func TestLinhaPayload(t *testing.T) {
	req, err := carregaLinhas()
	got, err := jsonLinhaDecode(req)
	if err != nil {
		t.Fatal("Falhei", err, got)
	}
}
