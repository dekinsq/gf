package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/net/ghttp"
	"github.com/dekinsq/gf/os/gsession"
	"github.com/dekinsq/gf/os/gtime"
	"time"
)

func main() {
	s := g.Server()
	s.SetConfigWithMap(g.Map{
		"SessionMaxAge":  time.Minute,
		"SessionStorage": gsession.NewStorageMemory(),
	})
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/set", func(r *ghttp.Request) {
			r.Session.Set("time", gtime.Timestamp())
			r.Response.Write("ok")
		})
		group.ALL("/get", func(r *ghttp.Request) {
			r.Response.Write(r.Session.Map())
		})
		group.ALL("/del", func(r *ghttp.Request) {
			r.Session.Clear()
			r.Response.Write("ok")
		})
	})
	s.SetPort(8199)
	s.Run()
}
