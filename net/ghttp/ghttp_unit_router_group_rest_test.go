// Copyright 2018 gf Author(https://github.com/dekinsq/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/dekinsq/gf.

// 分组路由测试
package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/frame/gmvc"
	"github.com/dekinsq/gf/net/ghttp"
	"github.com/dekinsq/gf/test/gtest"
)

type GroupCtlRest struct {
	gmvc.Controller
}

func (c *GroupCtlRest) Init(r *ghttp.Request) {
	c.Controller.Init(r)
	c.Response.Write("1")
}

func (c *GroupCtlRest) Shut() {
	c.Response.Write("2")
}

func (c *GroupCtlRest) Get() {
	c.Response.Write("Controller Get")
}

func (c *GroupCtlRest) Put() {
	c.Response.Write("Controller Put")
}

func (c *GroupCtlRest) Post() {
	c.Response.Write("Controller Post")
}

func (c *GroupCtlRest) Delete() {
	c.Response.Write("Controller Delete")
}

func (c *GroupCtlRest) Patch() {
	c.Response.Write("Controller Patch")
}

func (c *GroupCtlRest) Options() {
	c.Response.Write("Controller Options")
}

func (c *GroupCtlRest) Head() {
	c.Response.Header().Set("head-ok", "1")
}

type GroupObjRest struct{}

func (o *GroupObjRest) Init(r *ghttp.Request) {
	r.Response.Write("1")
}

func (o *GroupObjRest) Shut(r *ghttp.Request) {
	r.Response.Write("2")
}

func (o *GroupObjRest) Get(r *ghttp.Request) {
	r.Response.Write("Object Get")
}

func (o *GroupObjRest) Put(r *ghttp.Request) {
	r.Response.Write("Object Put")
}

func (o *GroupObjRest) Post(r *ghttp.Request) {
	r.Response.Write("Object Post")
}

func (o *GroupObjRest) Delete(r *ghttp.Request) {
	r.Response.Write("Object Delete")
}

func (o *GroupObjRest) Patch(r *ghttp.Request) {
	r.Response.Write("Object Patch")
}

func (o *GroupObjRest) Options(r *ghttp.Request) {
	r.Response.Write("Object Options")
}

func (o *GroupObjRest) Head(r *ghttp.Request) {
	r.Response.Header().Set("head-ok", "1")
}

func Test_Router_GroupRest1(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	g := s.Group("/api")
	ctl := new(GroupCtlRest)
	obj := new(GroupObjRest)
	g.REST("/ctl", ctl)
	g.REST("/obj", obj)
	g.REST("/{.struct}/{.method}", ctl)
	g.REST("/{.struct}/{.method}", obj)
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/api/ctl"), "1Controller Get2")
		t.Assert(client.PutContent("/api/ctl"), "1Controller Put2")
		t.Assert(client.PostContent("/api/ctl"), "1Controller Post2")
		t.Assert(client.DeleteContent("/api/ctl"), "1Controller Delete2")
		t.Assert(client.PatchContent("/api/ctl"), "1Controller Patch2")
		t.Assert(client.OptionsContent("/api/ctl"), "1Controller Options2")
		resp1, err := client.Head("/api/ctl")
		if err == nil {
			defer resp1.Close()
		}
		t.Assert(err, nil)
		t.Assert(resp1.Header.Get("head-ok"), "1")

		t.Assert(client.GetContent("/api/obj"), "1Object Get2")
		t.Assert(client.PutContent("/api/obj"), "1Object Put2")
		t.Assert(client.PostContent("/api/obj"), "1Object Post2")
		t.Assert(client.DeleteContent("/api/obj"), "1Object Delete2")
		t.Assert(client.PatchContent("/api/obj"), "1Object Patch2")
		t.Assert(client.OptionsContent("/api/obj"), "1Object Options2")
		resp2, err := client.Head("/api/obj")
		if err == nil {
			defer resp2.Close()
		}
		t.Assert(err, nil)
		t.Assert(resp2.Header.Get("head-ok"), "1")

		t.Assert(client.GetContent("/api/group-ctl-rest"), "Not Found")
		t.Assert(client.GetContent("/api/group-ctl-rest/get"), "1Controller Get2")
		t.Assert(client.PutContent("/api/group-ctl-rest/put"), "1Controller Put2")
		t.Assert(client.PostContent("/api/group-ctl-rest/post"), "1Controller Post2")
		t.Assert(client.DeleteContent("/api/group-ctl-rest/delete"), "1Controller Delete2")
		t.Assert(client.PatchContent("/api/group-ctl-rest/patch"), "1Controller Patch2")
		t.Assert(client.OptionsContent("/api/group-ctl-rest/options"), "1Controller Options2")
		resp3, err := client.Head("/api/group-ctl-rest/head")
		if err == nil {
			defer resp3.Close()
		}
		t.Assert(err, nil)
		t.Assert(resp3.Header.Get("head-ok"), "1")

		t.Assert(client.GetContent("/api/group-obj-rest"), "Not Found")
		t.Assert(client.GetContent("/api/group-obj-rest/get"), "1Object Get2")
		t.Assert(client.PutContent("/api/group-obj-rest/put"), "1Object Put2")
		t.Assert(client.PostContent("/api/group-obj-rest/post"), "1Object Post2")
		t.Assert(client.DeleteContent("/api/group-obj-rest/delete"), "1Object Delete2")
		t.Assert(client.PatchContent("/api/group-obj-rest/patch"), "1Object Patch2")
		t.Assert(client.OptionsContent("/api/group-obj-rest/options"), "1Object Options2")
		resp4, err := client.Head("/api/group-obj-rest/head")
		if err == nil {
			defer resp4.Close()
		}
		t.Assert(err, nil)
		t.Assert(resp4.Header.Get("head-ok"), "1")
	})
}

func Test_Router_GroupRest2(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.Group("/api", func(group *ghttp.RouterGroup) {
		ctl := new(GroupCtlRest)
		obj := new(GroupObjRest)
		group.REST("/ctl", ctl)
		group.REST("/obj", obj)
		group.REST("/{.struct}/{.method}", ctl)
		group.REST("/{.struct}/{.method}", obj)
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/api/ctl"), "1Controller Get2")
		t.Assert(client.PutContent("/api/ctl"), "1Controller Put2")
		t.Assert(client.PostContent("/api/ctl"), "1Controller Post2")
		t.Assert(client.DeleteContent("/api/ctl"), "1Controller Delete2")
		t.Assert(client.PatchContent("/api/ctl"), "1Controller Patch2")
		t.Assert(client.OptionsContent("/api/ctl"), "1Controller Options2")
		resp1, err := client.Head("/api/ctl")
		if err == nil {
			defer resp1.Close()
		}
		t.Assert(err, nil)
		t.Assert(resp1.Header.Get("head-ok"), "1")

		t.Assert(client.GetContent("/api/obj"), "1Object Get2")
		t.Assert(client.PutContent("/api/obj"), "1Object Put2")
		t.Assert(client.PostContent("/api/obj"), "1Object Post2")
		t.Assert(client.DeleteContent("/api/obj"), "1Object Delete2")
		t.Assert(client.PatchContent("/api/obj"), "1Object Patch2")
		t.Assert(client.OptionsContent("/api/obj"), "1Object Options2")
		resp2, err := client.Head("/api/obj")
		if err == nil {
			defer resp2.Close()
		}
		t.Assert(err, nil)
		t.Assert(resp2.Header.Get("head-ok"), "1")

		t.Assert(client.GetContent("/api/group-ctl-rest"), "Not Found")
		t.Assert(client.GetContent("/api/group-ctl-rest/get"), "1Controller Get2")
		t.Assert(client.PutContent("/api/group-ctl-rest/put"), "1Controller Put2")
		t.Assert(client.PostContent("/api/group-ctl-rest/post"), "1Controller Post2")
		t.Assert(client.DeleteContent("/api/group-ctl-rest/delete"), "1Controller Delete2")
		t.Assert(client.PatchContent("/api/group-ctl-rest/patch"), "1Controller Patch2")
		t.Assert(client.OptionsContent("/api/group-ctl-rest/options"), "1Controller Options2")
		resp3, err := client.Head("/api/group-ctl-rest/head")
		if err == nil {
			defer resp3.Close()
		}
		t.Assert(err, nil)
		t.Assert(resp3.Header.Get("head-ok"), "1")

		t.Assert(client.GetContent("/api/group-obj-rest"), "Not Found")
		t.Assert(client.GetContent("/api/group-obj-rest/get"), "1Object Get2")
		t.Assert(client.PutContent("/api/group-obj-rest/put"), "1Object Put2")
		t.Assert(client.PostContent("/api/group-obj-rest/post"), "1Object Post2")
		t.Assert(client.DeleteContent("/api/group-obj-rest/delete"), "1Object Delete2")
		t.Assert(client.PatchContent("/api/group-obj-rest/patch"), "1Object Patch2")
		t.Assert(client.OptionsContent("/api/group-obj-rest/options"), "1Object Options2")
		resp4, err := client.Head("/api/group-obj-rest/head")
		if err == nil {
			defer resp4.Close()
		}
		t.Assert(err, nil)
		t.Assert(resp4.Header.Get("head-ok"), "1")
	})
}
