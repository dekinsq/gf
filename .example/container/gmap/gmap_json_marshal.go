package main

import (
	"encoding/json"
	"fmt"
	"github.com/dekinsq/gf/frame/g"

	"github.com/dekinsq/gf/container/gmap"
)

func main() {
	m := gmap.New()
	m.Sets(g.MapAnyAny{
		"name":  "john",
		"score": 100,
	})
	b, _ := json.Marshal(m)
	fmt.Println(string(b))
}
