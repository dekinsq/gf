// Copyright 2020 gf Author(https://github.com/dekinsq/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/dekinsq/gf.

package gvar_test

import (
	"github.com/dekinsq/gf/container/gvar"
	"github.com/dekinsq/gf/frame/g"
	"github.com/dekinsq/gf/test/gtest"
	"testing"
)

func Test_Map(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := g.Map{
			"k1": "v1",
			"k2": "v2",
		}
		objOne := gvar.New(m, true)
		t.Assert(objOne.Map()["k1"], m["k1"])
		t.Assert(objOne.Map()["k2"], m["k2"])
	})
}
