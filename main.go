package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sibipro/financial-back-go/pkg/config"
)

// Config ...
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func main() {
	cfg := config.NewConfig{}
	fmt.Println("Successfully connected!")

	handleRequests()
}
