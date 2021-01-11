package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/os/gbuild"
)

func main() {
	g.Dump(gbuild.Info())
	g.Dump(gbuild.Map())
}
