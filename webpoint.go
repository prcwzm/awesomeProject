package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerWebPoint)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handlerWebPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
