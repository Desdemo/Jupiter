package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "r.URL.Path = %q\n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header{
			fmt.Fprintf(w, "r.Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 Not Found: %s\n", r.URL.Path)
	}
}

func main()  {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe("localhost:8080", engine))
}
