package main

import (
	"fmt"
	"github.com/dekinsq/gf/debug/gdebug"
)

func main() {
	gdebug.PrintStack()
	fmt.Println(gdebug.CallerPackage())
	fmt.Println(gdebug.CallerFunction())
}
