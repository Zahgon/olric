// The MIT License (MIT)
//
// Copyright (c) 2016 Josh Baker
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package server

import (
	"errors"

	"github.com/tidwall/redcon"
)

// errAuthRequired represents an error indicating that authentication is required to access the requested resource or operation.
var errAuthRequired = errors.New("Authentication required.")

// ServeMux is an RESP command multiplexer.
type ServeMux struct {
	config   *Config
	handlers map[string]redcon.Handler
}

// NewServeMux allocates and returns a new ServeMux.
func NewServeMux(c *Config) *ServeMux { _ = "STUB: not implemented"; return nil }

// HandleFunc registers the handler function for the given command.
func (m *ServeMux) HandleFunc(command string, handler redcon.Handler) {
	_ = "STUB: not implemented"
	return
}

// Handle registers the handler for the given command.
// If a handler already exists for command, Handle panics.
func (m *ServeMux) Handle(command string, handler redcon.Handler) {
	_ = "STUB: not implemented"
	return
}

// ServeRESP dispatches the command to the handler.
func (m *ServeMux) ServeRESP(conn redcon.Conn, cmd redcon.Command) {
	_ = "STUB: not implemented"
	return
}
