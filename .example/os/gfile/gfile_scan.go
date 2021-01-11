package main

import (
	"github.com/dekinsq/gf/os/gfile"
	"github.com/dekinsq/gf/util/gutil"
)

func main() {
	gutil.Dump(gfile.ScanDir("/Users/john/Documents", "*.*"))
	gutil.Dump(gfile.ScanDir("/home/john/temp/newproject", "*", true))
}
