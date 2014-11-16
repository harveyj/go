package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s)
}

type GreetingHandler struct {
	Greeting string
	Punct    string
	Who      string
}

func (gs GreetingHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, gs.Greeting, gs.Punct, gs.Who)
}

func main() {
	http.Handle("/greeting", &GreetingHandler{"Hello", ":", "Gophers!"})
	http.ListenAndServe("localhost:4000", nil)
}
