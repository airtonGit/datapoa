package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	linhasURL = "http://www.poatransporte.com.br/php/facades/process.php?a=nc&p=%&t=o"
)

var (
	linhasCache = &linhasStore{}
)

type linhasStore struct {
	Validade time.Time //15 minutos
	LinhaMap map[string]LinhaItinerario
	Linhas   []Linha
}

func chamarLinhasAPI() ([]byte, error) {
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

func carregarLinhas() ([]Linha, error) {
	jsonPayload, err := chamarLinhasAPI()
	if err != nil {
		return nil, fmt.Errorf("Falha API request %s", err)
	}
	linhas, err := jsonLinhaDecode(jsonPayload)
	if err != nil {
		return nil, fmt.Errorf("Falha decode payload %s", err)
	}
	return linhas, nil
}

func atualizaCache(linhas []Linha) {
	log.Println("Linhas: atualizando cache")
	linhasCache.Validade = time.Now().Add(15 * time.Minute)
	linhasCache.Linhas = linhas
}

func carregarLinhasCached() ([]Linha, error) {
	var linhas []Linha
	var err error

	if time.Now().Sub(linhasCache.Validade) > 1*time.Second {
		linhas, err = carregarLinhas()
		if err != nil {
			return nil, err
		}
		atualizaCache(linhas)
	} else {
		log.Println("linhas: cache hit")
		linhas = linhasCache.Linhas
	}
	return linhas, nil
}

func linhasHandler(w http.ResponseWriter, req *http.Request) {
	linhas, err := carregarLinhasCached()
	if err != nil {
		errorResponse(w, fmt.Sprintf("Falha %s", err))
		return
	}
	linhasResponse(w, linhas)
}

func linhasPesquisaNomeHandler(w http.ResponseWriter, req *http.Request) {
	nome := req.FormValue("nome")
	log.Println("linhas: pesquisa nome por:" + nome)
	linhas, err := carregarLinhasCached()
	if err != nil {
		errorResponse(w, fmt.Sprintf("Falha %s", err))
		return
	}
	resultados := []Linha{}
	for _, item := range linhas {
		if strings.Contains(strings.ToLower(item.Nome), strings.ToLower(nome)) {
			resultados = append(resultados, item)
		}
	}
	jsonResultado, err := json.Marshal(resultados)
	if err != nil {
		errorResponse(w, fmt.Sprintf("Falha resultados encode %s", err))
		return
	}
	w.Write(jsonResultado)
}
