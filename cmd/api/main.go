package main

import (
	"log"
	"net/http"

	"github.com/Wowsaruss/financial-back-go/pkg/config"
	"github.com/Wowsaruss/financial-back-go/pkg/finance"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.NewConfig()
	router := finance.NewRouter()

	log.Println("Server is running on", cfg.HostPort)
	log.Fatal(http.ListenAndServe(cfg.HostPort, router))
}
