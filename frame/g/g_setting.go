// Copyright 2019 gf Author(https://github.com/dekinsq/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/dekinsq/gf.

package g

import (
	"github.com/dekinsq/gf/internal/intlog"
	"github.com/dekinsq/gf/net/ghttp"
)

// SetEnabled enables/disables the GoFrame internal logging manually.
// Note that this function is not concurrent safe, be aware of the DATA RACE,
// which means you should call this function in your boot but not the runtime.
func SetDebug(enabled bool) {
	intlog.SetEnabled(enabled)
}

// SetServerGraceful enables/disables graceful reload feature of http Web Server.
// This feature is disabled in default.
// Deprecated, use configuration of ghttp.Server for controlling this feature.
func SetServerGraceful(enabled bool) {
	ghttp.SetGraceful(enabled)
}
