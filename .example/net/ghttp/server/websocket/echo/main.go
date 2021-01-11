package main

import (
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/net/ghttp"
	"github.com/dekinsq/gf/os/gfile"
	"github.com/dekinsq/gf/os/glog"
)

func main() {
	s := g.Server()
	s.BindHandler("/ws", func(r *ghttp.Request) {
		ws, err := r.WebSocket()
		if err != nil {
			glog.Error(err)
			return
		}
		for {
			msgType, msg, err := ws.ReadMessage()
			if err != nil {
				return
			}
			if err = ws.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})
	s.SetServerRoot(gfile.MainPkgPath())
	s.SetPort(8199)
	s.Run()
}
