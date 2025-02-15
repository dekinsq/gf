// Copyright 2018 gf Author(https://github.com/dekinsq/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/dekinsq/gf.

package ghttp

import (
	"github.com/dekinsq/gf/os/gfile"
	"strings"
	"time"

	"github.com/dekinsq/gf/os/gproc"
	"github.com/dekinsq/gf/os/gtimer"
	"github.com/dekinsq/gf/os/gview"
)

// utilAdmin is the controller for administration.
type utilAdmin struct{}

// Index shows the administration page.
func (p *utilAdmin) Index(r *Request) {
	data := map[string]interface{}{
		"pid":  gproc.Pid(),
		"path": gfile.SelfPath(),
		"uri":  strings.TrimRight(r.URL.Path, "/"),
	}
	buffer, _ := gview.ParseContent(`
            <html>
            <head>
                <title>GoFrame Web Server Admin</title>
            </head>
            <body>
                <p>Pid: {{.pid}}</p>
                <p>File Path: {{.path}}</p>
                <p><a href="{{$.uri}}/restart">Restart</a></p>
                <p><a href="{{$.uri}}/shutdown">Shutdown</a></p>
            </body>
            </html>
    `, data)
	r.Response.Write(buffer)
}

// Restart restarts all the servers in the process.
func (p *utilAdmin) Restart(r *Request) {
	var err error = nil
	// Custom start binary path when this process exits.
	path := r.GetQueryString("newExeFilePath")
	if path == "" {
		path = gfile.SelfPath()
	}
	if len(path) > 0 {
		err = RestartAllServer(path)
	} else {
		err = RestartAllServer()
	}
	if err == nil {
		r.Response.WriteExit("server restarted")
	} else {
		r.Response.WriteExit(err.Error())
	}
}

// Shutdown shuts down all the servers.
func (p *utilAdmin) Shutdown(r *Request) {
	if err := r.Server.Shutdown(); err != nil {
		r.Response.WriteExit(err.Error())
	}
	if err := ShutdownAllServer(); err == nil {
		r.Response.WriteExit("server shutdown")
	} else {
		r.Response.WriteExit(err.Error())
	}
}

// EnableAdmin enables the administration feature for the process.
// The optional parameter <pattern> specifies the URI for the administration page.
func (s *Server) EnableAdmin(pattern ...string) {
	p := "/debug/admin"
	if len(pattern) > 0 {
		p = pattern[0]
	}
	s.BindObject(p, &utilAdmin{})
}

// Shutdown shuts down current server.
func (s *Server) Shutdown() error {
	// It shuts down the server after 1 second, which is not triggered by system signal,
	// to ensure the response successfully to the client.
	gtimer.SetTimeout(time.Second, func() {
		// Only shut down current server.
		// It may have multiple underlying http servers.
		for _, v := range s.servers {
			v.close()
		}
	})
	return nil
}
