// Copyright 2018 gf Author(https://github.com/dekinsq/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/dekinsq/gf.

package ghttp

import (
	"fmt"
)

// getStatusHandler retrieves and returns the handler for given status code.
func (s *Server) getStatusHandler(status int, r *Request) HandlerFunc {
	domains := []string{r.GetHost(), gDEFAULT_DOMAIN}
	for _, domain := range domains {
		if f, ok := s.statusHandlerMap[s.statusHandlerKey(status, domain)]; ok {
			return f
		}
	}
	return nil
}

// setStatusHandler sets the handler for given status code.
// The parameter <pattern> is like: domain#status
func (s *Server) setStatusHandler(pattern string, handler HandlerFunc) {
	s.statusHandlerMap[pattern] = handler
}

// statusHandlerKey creates and returns key for given status and domain.
func (s *Server) statusHandlerKey(status int, domain string) string {
	return fmt.Sprintf("%s#%d", domain, status)
}

// BindStatusHandler registers handler for given status code.
func (s *Server) BindStatusHandler(status int, handler HandlerFunc) {
	s.setStatusHandler(s.statusHandlerKey(status, gDEFAULT_DOMAIN), handler)
}

// BindStatusHandlerByMap registers handler for given status code using map.
func (s *Server) BindStatusHandlerByMap(handlerMap map[int]HandlerFunc) {
	for k, v := range handlerMap {
		s.BindStatusHandler(k, v)
	}
}
