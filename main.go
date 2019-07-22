package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	router := newRouter()

	fmt.Println("Server is running on port: 10000!")
	log.Fatal(http.ListenAndServe(":10000", router))
}
