package main

import (
	"github.com/caarlos0/env"
	"github.com/flipper-zero/sub-sync/ghost"
	"github.com/sendinblue/APIv3-go-library/client"
	"log"
	"net/http"
)

var cfg config
var gh *ghost.Ghost
var sib *client.SendinBlue

func main() {
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Config", err)
	}

	var err error

	gh, err = ghost.New(cfg.GhostBaseURL, cfg.GhostApiKey)
	if err != nil {
		log.Fatalln("Ghost", err)
	}

	sib = getSibClient(cfg.SendinblueApiKey)

	log.Println("Server started")
	http.HandleFunc("/sendinblue", handleSendinblueWebhook)
	http.HandleFunc("/ghost", handleGhostWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
