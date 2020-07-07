package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func carregarItinerarioTodos() {
	linhas, err := carregarLinhas()
	if err != nil {
		log.Fatalf("carregar itinerarios %s", err)
	}

	for _, item := range linhas {
		fmt.Println("Carregando itinerario linha", item.Nome)
		arquivoJSON, err := os.OpenFile(fmt.Sprintf("json/itinerarios-%s.json", item.ID), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			log.Println("Falha ao abrir arquivo json para escrita", err)
			continue
		}
		jsonPayload, err := chamarItinerarioAPI(item.ID)
		runes := []rune(string(jsonPayload))
		preview := string(runes[0:15])
		fmt.Println("Recebido", preview)
		if err != nil {
			log.Fatalf("request %s", err)
		}
		if n, err := arquivoJSON.Write(jsonPayload); err != nil {
			log.Println("Falha ao escrever arquivo", n, err)
		}
		arquivoJSON.Close()
		fmt.Println("Aguardando 2s...")
		time.Sleep(2 * time.Second)
	}
}
