// Copyright 2017 gf Author(https://github.com/dekinsq/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/dekinsq/gf.

package glog

import "github.com/dekinsq/gf/container/gmap"

const (
	// Default group name for instance usage.
	DEFAULT_NAME = "default"
)

var (
	// Instances map.
	instances = gmap.NewStrAnyMap(true)
)

// Instance returns an instance of Logger with default settings.
// The parameter <name> is the name for the instance.
func Instance(name ...string) *Logger {
	key := DEFAULT_NAME
	if len(name) > 0 && name[0] != "" {
		key = name[0]
	}
	return instances.GetOrSetFuncLock(key, func() interface{} {
		return New()
	}).(*Logger)
}
