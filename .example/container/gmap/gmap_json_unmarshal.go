package main

import (
	"encoding/json"
	"fmt"
	"github.com/dekinsq/gf/container/gmap"
)

func main() {
	m := gmap.Map{}
	s := []byte(`{"name":"john","score":100}`)
	json.Unmarshal(s, &m)
	fmt.Println(m.Map())
}
