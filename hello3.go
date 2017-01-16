package main

import (
	"fmt"
	"log"
	"net/http"
)

type indexHandler struct {
}

func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Anna, good luck! ")
}

func main() {
	mux := http.NewServeMux()
	index := &indexHandler{}
	mux.Handle("/", index)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}