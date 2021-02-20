package main

import (
	"fmt"
	"jupiter"
	"net/http"
)

func main()  {
	r := jupiter.New()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	r.POST("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header{
			fmt.Fprintf(w, "r.Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")

}
