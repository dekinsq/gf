package main

import (
	"fmt"
	"time"

	"github.com/dekinsq/gf/util/gconv"
)

func main() {
	now := time.Now()
	t := gconv.Time(now.UnixNano() / 100)
	fmt.Println(now.UnixNano())
	fmt.Println(t.Nanosecond())
	fmt.Println(now.Nanosecond())
}
