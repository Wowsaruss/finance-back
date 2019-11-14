package main

import (
	"fmt"
	"net/http"

	"github.com/Wowsaruss/financial-back-go/pkg/config"
	"github.com/gorilla/mux"
)

func getType(w http.ResponseWriter, r *http.Request) {
	cfg := config.NewConfig()
	eventID := mux.Vars(r)["type"]
	fmt.Println(eventID, cfg)
}
