package main

import (
	"../mandelbrot"
	"../surface"
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
	http.HandleFunc("/surface", surfaceHandler)
	http.HandleFunc("/mandelbrot", mandelbrotHandler)
	http.HandleFunc("/MandelbrotColor", mandelbrotColorHandler)
	http.HandleFunc("/SuperSample", SuperSampleHandler)
	http.HandleFunc("/SuperSampleMax", SuperSampleMaxHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func SuperSampleMaxHandler(writer http.ResponseWriter, request *http.Request) {
	mandelbrot.SuperSampleMax(writer)

}

func SuperSampleHandler(writer http.ResponseWriter, request *http.Request) {
	mandelbrot.SuperSample(writer)
}

func mandelbrotColorHandler(writer http.ResponseWriter, request *http.Request) {
	mandelbrot.MandelbrotColor(writer)

}

func mandelbrotHandler(w http.ResponseWriter, r *http.Request) {
	mandelbrot.Mandelbrot(w)
}

func surfaceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	surface.SurfaceWeb(w)
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
