package main

import (
	"fmt"
	"log"
	"net/http"
	"social-media-go/src/config"
	"social-media-go/src/router"
)

func main() {
	config.Carregar()
	r := router.Gerar()

	fmt.Println("Escultando na porta", config.Porta)

	log.Fatal(http.ListenAndServe(":3000", r))
}
