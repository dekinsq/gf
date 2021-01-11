package main

import (
	"time"

	"github.com/dekinsq/gf/os/glog"
	"github.com/dekinsq/gf/os/gtime"
)

// 测试删除日志文件是否会重建日志文件
func main() {
	path := "/Users/john/Temp/test"
	glog.SetPath(path)
	for {
		glog.Println(gtime.Now().String())
		time.Sleep(time.Second)
	}
}
