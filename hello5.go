package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"log"
)



func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Welcome Ania")
}

func main() {
	r :=mux.NewRouter()
	r.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
