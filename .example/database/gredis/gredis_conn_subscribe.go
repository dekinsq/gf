package main

import (
	"fmt"

	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/util/gconv"
)

func main() {
	conn := g.Redis().Conn()
	defer conn.Close()
	_, err := conn.Do("SUBSCRIBE", "channel")
	if err != nil {
		panic(err)
	}
	for {
		reply, err := conn.Receive()
		if err != nil {
			panic(err)
		}
		fmt.Println(gconv.Strings(reply))
	}
}
