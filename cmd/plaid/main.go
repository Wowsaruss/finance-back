package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Wowsaruss/financial-back-go/pkg/config"
	"github.com/plaid/plaid-go/plaid"
)

func main() {
	cfg := config.NewConfig()

	clientOptions := plaid.ClientOptions{
		cfg.PlaidClientID,
		cfg.PlaidSecret,
		cfg.PlaidPublicKey,
		plaid.Development,
		&http.Client{},
	}

	client, err := plaid.NewClient(clientOptions)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v", *client)
}
