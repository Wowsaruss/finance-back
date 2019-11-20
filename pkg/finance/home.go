package finance

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w, "Welcome to the HomePage!!")
	fmt.Println("Home!")
}
