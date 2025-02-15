// Copyright 2017 gf Author(https://github.com/dekinsq/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/dekinsq/gf.

// Package gview implements a template engine based on text/template.
//
// Reserved template variable names:
//     I18nLanguage: Assign this variable to define i18n language for each page.
package gview

import (
	"github.com/dekinsq/gf/container/gmap"
	"github.com/dekinsq/gf/internal/intlog"

	"github.com/dekinsq/gf"
	"github.com/dekinsq/gf/container/garray"
	"github.com/dekinsq/gf/internal/cmdenv"
	"github.com/dekinsq/gf/os/gfile"
	"github.com/dekinsq/gf/os/glog"
)

// View object for template engine.
type View struct {
	paths        *garray.StrArray       // Searching array for path, NOT concurrent-safe for performance purpose.
	data         map[string]interface{} // Global template variables.
	funcMap      map[string]interface{} // Global template function map.
	fileCacheMap *gmap.StrAnyMap        // File cache map.
	config       Config                 // Extra configuration for the view.
}

type (
	Params  = map[string]interface{} // Params is type for template params.
	FuncMap = map[string]interface{} // FuncMap is type for custom template functions.
)

var (
	// Default view object.
	defaultViewObj *View
)

// checkAndInitDefaultView checks and initializes the default view object.
// The default view object will be initialized just once.
func checkAndInitDefaultView() {
	if defaultViewObj == nil {
		defaultViewObj = New()
	}
}

// ParseContent parses the template content directly using the default view object
// and returns the parsed content.
func ParseContent(content string, params ...Params) (string, error) {
	checkAndInitDefaultView()
	return defaultViewObj.ParseContent(content, params...)
}

// New returns a new view object.
// The parameter <path> specifies the template directory path to load template files.
func New(path ...string) *View {
	view := &View{
		paths:        garray.NewStrArray(),
		data:         make(map[string]interface{}),
		funcMap:      make(map[string]interface{}),
		fileCacheMap: gmap.NewStrAnyMap(true),
		config:       DefaultConfig(),
	}
	if len(path) > 0 && len(path[0]) > 0 {
		if err := view.SetPath(path[0]); err != nil {
			intlog.Error(err)
		}
	} else {
		// Customized dir path from env/cmd.
		if envPath := cmdenv.Get("gf.gview.path").String(); envPath != "" {
			if gfile.Exists(envPath) {
				if err := view.SetPath(envPath); err != nil {
					intlog.Error(err)
				}
			} else {
				if errorPrint() {
					glog.Errorf("Template directory path does not exist: %s", envPath)
				}
			}
		} else {
			// Dir path of working dir.
			if err := view.SetPath(gfile.Pwd()); err != nil {
				intlog.Error(err)
			}
			// Dir path of binary.
			if selfPath := gfile.SelfDir(); selfPath != "" && gfile.Exists(selfPath) {
				if err := view.AddPath(selfPath); err != nil {
					intlog.Error(err)
				}
			}
			// Dir path of main package.
			if mainPath := gfile.MainPkgPath(); mainPath != "" && gfile.Exists(mainPath) {
				if err := view.AddPath(mainPath); err != nil {
					intlog.Error(err)
				}
			}
		}
	}
	view.SetDelimiters("{{", "}}")
	// default build-in variables.
	view.data["GF"] = map[string]interface{}{
		"version": gf.VERSION,
	}
	// default build-in functions.
	view.BindFuncMap(FuncMap{
		"eq":         view.funcEq,
		"ne":         view.funcNe,
		"lt":         view.funcLt,
		"le":         view.funcLe,
		"gt":         view.funcGt,
		"ge":         view.funcGe,
		"text":       view.funcText,
		"html":       view.funcHtmlEncode,
		"htmlencode": view.funcHtmlEncode,
		"htmldecode": view.funcHtmlDecode,
		"encode":     view.funcHtmlEncode,
		"decode":     view.funcHtmlDecode,
		"url":        view.funcUrlEncode,
		"urlencode":  view.funcUrlEncode,
		"urldecode":  view.funcUrlDecode,
		"date":       view.funcDate,
		"substr":     view.funcSubStr,
		"strlimit":   view.funcStrLimit,
		"concat":     view.funcConcat,
		"replace":    view.funcReplace,
		"compare":    view.funcCompare,
		"hidestr":    view.funcHideStr,
		"highlight":  view.funcHighlight,
		"toupper":    view.funcToUpper,
		"tolower":    view.funcToLower,
		"nl2br":      view.funcNl2Br,
		"include":    view.funcInclude,
		"dump":       view.funcDump,
	})

	return view
}
