package main

import (
	"fmt"

	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/util/gconv"
)

func main() {
	conn := g.Redis().Conn()
	defer conn.Close()
	conn.Send("SET", "foo", "bar")
	conn.Send("GET", "foo")
	conn.Flush()
	// reply from SET
	conn.Receive()
	// reply from GET
	v, _ := conn.Receive()
	fmt.Println(gconv.String(v))
}
