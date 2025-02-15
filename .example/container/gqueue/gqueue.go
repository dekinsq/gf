package main

import (
	"fmt"
	"time"

	"github.com/dekinsq/gf/container/gqueue"
	"github.com/dekinsq/gf/os/gtime"
	"github.com/dekinsq/gf/os/gtimer"
)

func main() {
	q := gqueue.New()
	// 数据生产者，每隔1秒往队列写数据
	gtimer.SetInterval(time.Second, func() {
		v := gtime.Now().String()
		q.Push(v)
		fmt.Println("Push:", v)
	})

	// 3秒后关闭队列
	gtimer.SetTimeout(3*time.Second, func() {
		q.Close()
	})

	// 消费者，不停读取队列数据并输出到终端
	for {
		if v := q.Pop(); v != nil {
			fmt.Println(" Pop:", v)
		} else {
			break
		}
	}
}
