package main

import (
	"net/http"
	"fmt"
)



func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Welcome Ania")
}

func main() {
	http.HandleFunc("/", handler3)
	http.ListenAndServe(":8080", nil)
}