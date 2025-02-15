package main

import (
	"fmt"

	"github.com/dekinsq/gf/frame/g"
)

func main() {
	g.Config().SetFileName("config3.toml")
	if r, err := g.DB().Table("user").Where("uid=?", 1).One(); err == nil {
		fmt.Println(r["uid"].Int())
		fmt.Println(r["name"].String())
	} else {
		fmt.Println(err)
	}

	if r, err := g.DB("user").Table("user").Where("uid=?", 1).One(); err == nil {
		fmt.Println(r["uid"].Int())
		fmt.Println(r["name"].String())
	} else {
		fmt.Println(err)
	}
}
