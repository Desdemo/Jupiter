package main

import (
	"jupiter"
	"net/http"
)

func main()  {
	r := jupiter.New()
	
	r.Get("/", func(c *jupiter.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Jupiter</h1>")
	})
	
	r.Get("/hello", func(c *jupiter.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *jupiter.Context) {
		c.JSON(http.StatusOK, jupiter.H{
			"username" : c.PostForm("username"),
			"password" : c.PostForm("password"),
		})
	})

	r.Run(":9999")
}

