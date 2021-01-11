package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/util/gvalid"
)

// string默认值校验
func main() {
	type User struct {
		Uid string `gvalid:"uid@integer"`
	}

	user := &User{}

	g.Dump(gvalid.CheckStruct(user, nil))
}
