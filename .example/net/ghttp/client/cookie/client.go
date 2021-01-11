package main

import (
	"fmt"

	"github.com/dekinsq/gf/net/ghttp"
	"github.com/dekinsq/gf/os/glog"
)

func main() {
	c := ghttp.NewClient()
	c.SetHeader("Cookie", "name=john; score=100")
	if r, e := c.Get("http://127.0.0.1:8199/"); e != nil {
		glog.Error(e)
	} else {
		fmt.Println(string(r.ReadAll()))
	}
}
