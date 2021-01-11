package main

import (
	"fmt"

	"github.com/dekinsq/gf/net/ghttp"
	"github.com/dekinsq/gf/os/glog"
	"github.com/dekinsq/gf/text/gregex"
)

type MyWriter struct {
	logger *glog.Logger
}

func (w *MyWriter) Write(p []byte) (n int, err error) {
	s := string(p)
	if gregex.IsMatchString(`\[(PANI|FATA)\]`, s) {
		fmt.Println("SERIOUS ISSUE OCCURRED!! I'd better tell monitor in first time!")
		ghttp.PostContent("http://monitor.mydomain.com", s)
	}
	return w.logger.Write(p)
}

func main() {
	glog.SetWriter(&MyWriter{
		logger: glog.New(),
	})
	glog.Debug("DEBUG")
	glog.Fatal("FATAL ERROR")

}
