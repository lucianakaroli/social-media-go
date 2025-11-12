package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lucianakaroli/social-media-go/src/config"
	"github.com/lucianakaroli/social-media-go/src/router"
)


func main() {
	config.Carregar()

	fmt.Println(config.Porta)
	fmt.Println("Rodando API")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":3000", r))
}
