package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/net/ghttp"
)

func main() {
	s := g.Server()
	group := s.Group("/api")
	group.ALL("/all", func(r *ghttp.Request) {
		r.Response.Write("all")
	})
	group.GET("/get", func(r *ghttp.Request) {
		r.Response.Write("get")
	})
	group.POST("/post", func(r *ghttp.Request) {
		r.Response.Write("post")
	})
	s.SetPort(8199)
	s.Run()
}
