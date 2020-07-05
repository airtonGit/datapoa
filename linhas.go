package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	linhasURL = "http://www.poatransporte.com.br/php/facades/process.php?a=nc&p=%&t=o"
)

func carregaLinhas() ([]byte, error) {
	datapoaClient := http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := datapoaClient.Get(linhasURL)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func jsonLinhaDecode(jsonPayload []byte) ([]Linha, error) {
	var linhas []Linha
	if err := json.Unmarshal(jsonPayload, &linhas); err != nil {
		return nil, err
	}
	return linhas, nil
}

func linhasResponse(w http.ResponseWriter, linhas []Linha) {

	jsonPayload, err := json.Marshal(linhas)
	if err != nil {
		errorResponse(w, "Falha ao jsonfy linhas "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonPayload)
}

func linhasHandler(w http.ResponseWriter, req *http.Request) {
	jsonPayload, err := carregaLinhas()
	if err != nil {
		errorResponse(w, fmt.Sprintf("Falha request %s", err))
		return
	}
	linhas, err := jsonLinhaDecode(jsonPayload)
	if err != nil {
		errorResponse(w, fmt.Sprintf("Falha payload %s", err))
		return
	}
	linhasResponse(w, linhas)
}
