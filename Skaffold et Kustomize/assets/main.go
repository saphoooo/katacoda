package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello skaffold with kustomize!")
}

func main() {
	http.HandleFunc("/", hello)
	log.Println("start listening on port :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
