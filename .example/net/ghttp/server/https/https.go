package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("来自于HTTPS的：哈喽世界！")
	})
	s.EnableHTTPS("./server.crt", "./server.key")
	s.SetAccessLogEnabled(true)
	s.SetPort(8199)
	s.Run()
}
