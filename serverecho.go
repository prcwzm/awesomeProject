package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int = 0

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous/", lissajousHandlerCyc)
	http.HandleFunc("/lissajous", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func lissajousHandlerCyc(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles, err := strconv.Atoi(request.Form.Get("cycles"))
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("%d cycles\n", cycles)
	lissajousGreenBlackCyc(writer, cycles)
}

func lissajousHandler(writer http.ResponseWriter, request *http.Request) {
	lissajousGreenBlack(writer)
}

func GetUrlArgs(request *http.Request, name string) (arg string) {
	value := request.URL.Query()
	arg = value.Get(name)
	return arg
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	count++

	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "From[%q] = %q \n", k, v)
	}
	fmt.Fprintf(w, "URL.Path: %v\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "counter: %d\n", count)
	defer mu.Unlock()

}
