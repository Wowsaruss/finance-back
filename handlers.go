package main

import (
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/", homePage)

	http.HandleFunc("/insert", insertData)

	log.Fatal(http.ListenAndServe(":10000", nil))
}
