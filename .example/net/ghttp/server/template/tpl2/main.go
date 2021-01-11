package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/frame/gins"
	"github.com/dekinsq/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.SetServerRoot("public")
	s.SetNameToUriType(ghttp.URI_TYPE_ALLLOWER)
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)
	s.SetPort(2333)

	s.BindHandler("/", func(r *ghttp.Request) {
		content, _ := gins.View().Parse("test.html", nil)
		r.Response.Write(content)
	})

	s.Run()
}
