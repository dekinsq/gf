// 向指定进程发送进程消息。
package main

import (
	"fmt"

	"github.com/dekinsq/gf/os/gproc"
)

func main() {
	err := gproc.Send(22988, []byte{30})
	fmt.Println(err)
}
