package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi...")
	})

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
