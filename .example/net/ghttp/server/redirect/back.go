package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/page", func(r *ghttp.Request) {
		r.Response.Writeln(`<a href="/back">back</a>`)
	})
	s.BindHandler("/back", func(r *ghttp.Request) {
		r.Response.RedirectBack()
	})
	s.SetPort(8199)
	s.Run()
}
