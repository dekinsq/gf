// Copyright 2018 gf Author(https://github.com/dekinsq/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/dekinsq/gf.

package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/net/ghttp"
	"github.com/dekinsq/gf/test/gtest"
)

func Test_GetUrl(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/url", func(r *ghttp.Request) {
		r.Response.Write(r.GetUrl())
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", p)
		client := ghttp.NewClient()
		client.SetBrowserMode(true)
		client.SetPrefix(prefix)

		t.Assert(client.GetContent("/url"), prefix+"/url")
	})
}
