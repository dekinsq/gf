package main

import (
	"github.com/dekinsq/gf/os/glog"
)

func main() {
	glog.Line().Debug("this is the short file name with its line number")
	glog.Line(true).Debug("lone file name with line number")
}
