package main

import (
	"fmt"

	"github.com/dekinsq/gf/frame/g"
)

// 使用默认的config.toml配置文件读取配置
func main() {
	c := g.Config()
	fmt.Println(c.GetArray("memcache"))
}
