// Copyright 2019 gf Author(https://github.com/dekinsq/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/dekinsq/gf.

// go test *.go -bench=".*"

package gstr_test

import (
	"testing"

	"github.com/dekinsq/gf/test/gtest"
	"github.com/dekinsq/gf/text/gstr"
)

func Test_OctStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.OctStr(`\346\200\241`), "怡")
	})
}
