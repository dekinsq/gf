package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/net/ghttp"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func main() {
	s := g.Server()
	s.Group("/api.v2", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.ALL("/user/list", func(r *ghttp.Request) {
			r.Response.Write("list")
		})
	})
	s.SetPort(8199)
	s.Run()
}
