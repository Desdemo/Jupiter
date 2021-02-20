package jupiter

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct{
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method string,pattern string, handler HandlerFunc)  {
	e.router.addRouter(method, pattern, handler)
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
	c := newContext(w, r)
	e.router.handle(c)
}


