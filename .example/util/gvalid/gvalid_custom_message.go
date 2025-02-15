package main

import (
	"fmt"
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/util/gvalid"
)

func main() {
	g.I18n().SetLanguage("cn")
	err := gvalid.Check("", "required", nil)
	fmt.Println(err.String())
}
