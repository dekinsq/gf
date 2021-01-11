package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/os/gfile"
	"github.com/dekinsq/gf/os/glog"
)

func main() {
	path := "/tmp/glog-cat"
	glog.SetPath(path)
	glog.Stdout(false).Cat("cat1").Cat("cat2").Println("test")
	list, err := gfile.ScanDir(path, "*", true)
	g.Dump(err)
	g.Dump(list)
}
