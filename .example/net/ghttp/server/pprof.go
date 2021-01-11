package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.Domain("localhost").EnablePProf()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("哈喽世界！")
	})
	s.SetPort(8199)
	s.Run()
}
