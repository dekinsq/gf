// Copyright 2020 gf Author(https://github.com/dekinsq/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/dekinsq/gf.

package utils_test

import (
	"github.com/dekinsq/gf/internal/utils"
	"github.com/dekinsq/gf/test/gtest"
	"io/ioutil"
	"testing"
)

func Test_ReadCloser(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			n    int
			b    = make([]byte, 3)
			body = utils.NewReadCloser([]byte{1, 2, 3, 4}, false)
		)
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{1, 2, 3})
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{4})

		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{})
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{})
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			r    []byte
			body = utils.NewReadCloser([]byte{1, 2, 3, 4}, false)
		)
		r, _ = ioutil.ReadAll(body)
		t.Assert(r, []byte{1, 2, 3, 4})
		r, _ = ioutil.ReadAll(body)
		t.Assert(r, []byte{})
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			n    int
			r    []byte
			b    = make([]byte, 3)
			body = utils.NewReadCloser([]byte{1, 2, 3, 4}, true)
		)
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{1, 2, 3})
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{4})

		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{1, 2, 3})
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{4})

		r, _ = ioutil.ReadAll(body)
		t.Assert(r, []byte{1, 2, 3, 4})
		r, _ = ioutil.ReadAll(body)
		t.Assert(r, []byte{1, 2, 3, 4})
	})
}
