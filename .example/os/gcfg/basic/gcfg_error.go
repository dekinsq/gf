package main

import (
	"fmt"

	"github.com/dekinsq/gf/frame/g"
)

func main() {
	fmt.Println(g.Config().Get("none"))
}
