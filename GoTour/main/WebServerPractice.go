package main

import (
	"net/http"
	"fmt"
	"log"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Crowhyc WebS123ite")
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}

type Server struct {
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Crowhyc WebSite")
}
func main() {
	var s Server
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"hello", ":", "world"})
	http.Handle("/", s)
	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
