package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/net/ghttp"
)

// https://github.com/dekinsq/gf/issues/437
func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.WriteTpl("client/layout.html")
	})
	s.SetPort(8199)
	s.Run()
}
