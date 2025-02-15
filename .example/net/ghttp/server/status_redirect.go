package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/status/:status", func(r *ghttp.Request) {
		r.Response.Write("woops, status ", r.Get("status"), " found")
	})
	s.BindStatusHandler(404, func(r *ghttp.Request) {
		r.Response.RedirectTo("/status/404")
	})
	s.SetPort(8199)
	s.Run()
}
