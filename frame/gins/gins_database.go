// Copyright 2019 gf Author(https://github.com/dekinsq/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/dekinsq/gf.

package gins

import (
	"fmt"
	"github.com/dekinsq/gf/internal/intlog"
	"github.com/dekinsq/gf/text/gstr"
	"github.com/dekinsq/gf/util/gutil"

	"github.com/dekinsq/gf/database/gdb"
	"github.com/dekinsq/gf/text/gregex"
	"github.com/dekinsq/gf/util/gconv"
)

const (
	gFRAME_CORE_COMPONENT_NAME_DATABASE = "gf.core.component.database"
	gDATABASE_NODE_NAME                 = "database"
)

// Database returns an instance of database ORM object
// with specified configuration group name.
func Database(name ...string) gdb.DB {
	group := gdb.DEFAULT_GROUP_NAME
	if len(name) > 0 && name[0] != "" {
		group = name[0]
	}
	instanceKey := fmt.Sprintf("%s.%s", gFRAME_CORE_COMPONENT_NAME_DATABASE, group)
	db := instances.GetOrSetFuncLock(instanceKey, func() interface{} {
		var (
			configMap     map[string]interface{}
			configNodeKey string
		)
		// It firstly searches the configuration of the instance name.
		if Config().Available() {
			configNodeKey, _ = gutil.MapPossibleItemByKey(Config().GetMap("."), gDATABASE_NODE_NAME)
			if configNodeKey == "" {
				configNodeKey = gDATABASE_NODE_NAME
			}
			configMap = Config().GetMap(configNodeKey)
		}
		if len(configMap) == 0 && !gdb.IsConfigured() {
			panic(fmt.Sprintf(`database init failed: "%s" node not found, is config file or configuration missing?`, gDATABASE_NODE_NAME))
		}
		if len(configMap) == 0 {
			configMap = make(map[string]interface{})
		}
		// Parse <m> as map-slice and adds it to gdb's global configurations.
		for g, groupConfig := range configMap {
			cg := gdb.ConfigGroup{}
			switch value := groupConfig.(type) {
			case []interface{}:
				for _, v := range value {
					if node := parseDBConfigNode(v); node != nil {
						cg = append(cg, *node)
					}
				}
			case map[string]interface{}:
				if node := parseDBConfigNode(value); node != nil {
					cg = append(cg, *node)
				}
			}
			if len(cg) > 0 {
				if gdb.GetConfig(group) == nil {
					intlog.Printf("add configuration for group: %s, %#v", g, cg)
					gdb.SetConfigGroup(g, cg)
				} else {
					intlog.Printf("ignore configuration as it already exists for group: %s, %#v", g, cg)
					intlog.Printf("%s, %#v", g, cg)
				}
			}
		}
		// Parse <m> as a single node configuration,
		// which is the default group configuration.
		if node := parseDBConfigNode(configMap); node != nil {
			cg := gdb.ConfigGroup{}
			if node.LinkInfo != "" || node.Host != "" {
				cg = append(cg, *node)
			}

			if len(cg) > 0 {
				if gdb.GetConfig(group) == nil {
					intlog.Printf("add configuration for group: %s, %#v", gdb.DEFAULT_GROUP_NAME, cg)
					gdb.SetConfigGroup(gdb.DEFAULT_GROUP_NAME, cg)
				} else {
					intlog.Printf("ignore configuration as it already exists for group: %s, %#v", gdb.DEFAULT_GROUP_NAME, cg)
					intlog.Printf("%s, %#v", gdb.DEFAULT_GROUP_NAME, cg)
				}
			}
		}
		// Create a new ORM object with given configurations.
		if db, err := gdb.New(name...); err == nil {
			if Config().Available() {
				// Initialize logger for ORM.
				var loggerConfigMap map[string]interface{}
				loggerConfigMap = Config().GetMap(fmt.Sprintf("%s.%s", configNodeKey, gLOGGER_NODE_NAME))
				if len(loggerConfigMap) == 0 {
					loggerConfigMap = Config().GetMap(configNodeKey)
				}
				if len(loggerConfigMap) > 0 {
					if err := db.GetLogger().SetConfigWithMap(loggerConfigMap); err != nil {
						panic(err)
					}
				}
			}
			return db
		} else {
			// It panics often because it dose not find its configuration for given group.
			panic(err)
		}
		return nil
	})
	if db != nil {
		return db.(gdb.DB)
	}
	return nil
}

func parseDBConfigNode(value interface{}) *gdb.ConfigNode {
	nodeMap, ok := value.(map[string]interface{})
	if !ok {
		return nil
	}
	node := &gdb.ConfigNode{}
	err := gconv.Struct(nodeMap, node)
	if err != nil {
		panic(err)
	}
	if _, v := gutil.MapPossibleItemByKey(nodeMap, "link"); v != nil {
		node.LinkInfo = gconv.String(v)
	}
	// Parse link syntax.
	if node.LinkInfo != "" && node.Type == "" {
		match, _ := gregex.MatchString(`([a-z]+):(.+)`, node.LinkInfo)
		if len(match) == 3 {
			node.Type = gstr.Trim(match[1])
			node.LinkInfo = gstr.Trim(match[2])
		}
	}
	return node
}
