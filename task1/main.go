package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type HttpHandlerDecorator func(http.HandlerFunc) http.HandlerFunc

func Handler(h http.HandlerFunc, decors ...HttpHandlerDecorator) http.HandlerFunc {
	for index := range decors {
		decor := decors[len(decors)-1-index]
		h = decor(h)
	}
	return h
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "ok")
}

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func WithServerHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("--->WithServerHeader() request %#v", r.Header)
		for key, value := range r.Header {
			fmt.Printf("key is %#v value is %#v\n", key, value)
			w.Header().Set(key, value[0])
		}
		h(w, r)
	}
}

func WithEnvHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("$$$->WithEnvHeader() %#v", os.Getenv("GOVERSION"))
		key := "GOVERSION"
		goversion := os.Getenv(key)
		log.Printf("%s", goversion)
		w.Header().Set(key, goversion)

		h(w, r)
	}
}

func version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", os.Getenv("GOVERSION"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", WithServerHeader(healthcheck))
	r.HandleFunc("/headers", WithServerHeader(headers))
	r.HandleFunc("/version", Handler(version, WithServerHeader, WithEnvHeader))

	n := negroni.Classic()
	n.UseHandler(r)
	err := http.ListenAndServe(":8090", n)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// test with
// export GOVERSION="hello"

// $ curl -I http://localhost:8090/healthz
// HTTP/1.1 200 OK
//  Accept: */*
//  User-Agent: curl/7.68.0
//  Date: Sat, 02 Oct 2021 13:53:19 GMT
//  Content-Length: 2
//  Content-Type: text/plain; charset=utf-8

// $ curl -I http://localhost:8090/version
// HTTP/1.1 200 OK
// Accept: */*
// Goversion: hello
// User-Agent: curl/7.68.0
// Date: Sat, 02 Oct 2021 13:53:22 GMT
// Content-Length: 5
// Content-Type: text/plain; charset=utf-8
