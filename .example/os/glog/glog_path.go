package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/os/gfile"
	"github.com/dekinsq/gf/os/glog"
)

// 设置日志输出路径
func main() {
	path := "/tmp/glog"
	glog.SetPath(path)
	glog.Println("日志内容")
	list, err := gfile.ScanDir(path, "*")
	g.Dump(err)
	g.Dump(list)
}
