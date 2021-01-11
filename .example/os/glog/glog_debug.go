package main

import (
	"time"

	"github.com/dekinsq/gf/os/glog"
	"github.com/dekinsq/gf/os/gtime"
	"github.com/dekinsq/gf/os/gtimer"
)

func main() {
	gtimer.SetTimeout(3*time.Second, func() {
		glog.SetDebug(false)
	})
	for {
		glog.Debug(gtime.Datetime())
		time.Sleep(time.Second)
	}
}
