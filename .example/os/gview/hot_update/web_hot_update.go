package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		g.View().SetPath(`D:\Workspace\Go\GOPATH\src\gitee.com\johng\gf\geg\os\gview`)
		b, _ := g.View().Parse("gview.tpl", nil)
		r.Response.Write(b)
	})
	s.Run()
}
