package jupiter

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct{
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: map[string]HandlerFunc{}}
}

func (e *Engine) addRoute(method string,pattern string, handler HandlerFunc)  {
	key := method + "-" + pattern
	e.router[key] = handler
}

func (e *Engine) Get(pattern string, handler HandlerFunc)  {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc)  {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string)  {
	http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.router[key]; ok{
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 Not Found: %s\n", r.URL.Path)
	}
}


