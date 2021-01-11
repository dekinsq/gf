package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/os/glog"
)

func main() {
	err := glog.SetConfigWithMap(g.Map{
		"prefix": "[TEST]",
	})
	if err != nil {
		panic(err)
	}
	glog.Info(1)
}
