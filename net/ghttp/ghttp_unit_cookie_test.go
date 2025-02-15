// Copyright 2018 gf Author(https://github.com/dekinsq/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/dekinsq/gf.

package ghttp_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/net/ghttp"
	"github.com/dekinsq/gf/test/gtest"
)

func Test_Cookie(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/set", func(r *ghttp.Request) {
		r.Cookie.Set(r.GetString("k"), r.GetString("v"))
	})
	s.BindHandler("/get", func(r *ghttp.Request) {
		r.Response.Write(r.Cookie.Get(r.GetString("k")))
	})
	s.BindHandler("/remove", func(r *ghttp.Request) {
		r.Cookie.Remove(r.GetString("k"))
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := ghttp.NewClient()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		r1, e1 := client.Get("/set?k=key1&v=100")
		if r1 != nil {
			defer r1.Close()
		}

		t.Assert(e1, nil)
		t.Assert(r1.ReadAllString(), "")

		t.Assert(client.GetContent("/set?k=key2&v=200"), "")

		t.Assert(client.GetContent("/get?k=key1"), "100")
		t.Assert(client.GetContent("/get?k=key2"), "200")
		t.Assert(client.GetContent("/get?k=key3"), "")
		t.Assert(client.GetContent("/remove?k=key1"), "")
		t.Assert(client.GetContent("/remove?k=key3"), "")
		t.Assert(client.GetContent("/remove?k=key4"), "")
		t.Assert(client.GetContent("/get?k=key1"), "")
		t.Assert(client.GetContent("/get?k=key2"), "200")
	})
}

func Test_SetHttpCookie(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/set", func(r *ghttp.Request) {
		r.Cookie.SetHttpCookie(&http.Cookie{
			Name:  r.GetString("k"),
			Value: r.GetString("v"),
		})
	})
	s.BindHandler("/get", func(r *ghttp.Request) {
		r.Response.Write(r.Cookie.Get(r.GetString("k")))
	})
	s.BindHandler("/remove", func(r *ghttp.Request) {
		r.Cookie.Remove(r.GetString("k"))
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := ghttp.NewClient()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		r1, e1 := client.Get("/set?k=key1&v=100")
		if r1 != nil {
			defer r1.Close()
		}
		t.Assert(e1, nil)
		t.Assert(r1.ReadAllString(), "")

		t.Assert(client.GetContent("/set?k=key2&v=200"), "")

		t.Assert(client.GetContent("/get?k=key1"), "100")
		t.Assert(client.GetContent("/get?k=key2"), "200")
		t.Assert(client.GetContent("/get?k=key3"), "")
		t.Assert(client.GetContent("/remove?k=key1"), "")
		t.Assert(client.GetContent("/remove?k=key3"), "")
		t.Assert(client.GetContent("/remove?k=key4"), "")
		t.Assert(client.GetContent("/get?k=key1"), "")
		t.Assert(client.GetContent("/get?k=key2"), "200")
	})
}
