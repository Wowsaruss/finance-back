package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Wowsaruss/financial-back-go/pkg/config"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.NewConfig()
	router := newRouter()

	fmt.Println("Server is running!")
	log.Fatal(http.ListenAndServe(cfg.Port, router))
}
