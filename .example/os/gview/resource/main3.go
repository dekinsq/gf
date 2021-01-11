package main

import (
	"fmt"

	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/os/gres"
	_ "github.com/dekinsq/gf/os/gres/testdata"
)

func main() {
	gres.Dump()

	v := g.View()
	s, err := v.Parse("index.html")
	fmt.Println(err)
	fmt.Println(s)
}
