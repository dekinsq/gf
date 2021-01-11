package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln(r.Get("name"))
	})
	s.BindHookHandlerByMap("/", map[string]ghttp.HandlerFunc{
		ghttp.HookBeforeServe: func(r *ghttp.Request) {
			r.SetParam("name", "john")
		},
	})
	s.SetPort(8199)
	s.Run()
}
