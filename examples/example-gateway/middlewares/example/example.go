// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package example

import (
	"context"

	"github.com/mcuadros/go-jsonschema-generator"
	zanzibar "github.com/uber/zanzibar/runtime"
)

// Options for middleware configuration
type Options struct {
	Foo string `json:"Foo"`
	Bar int    `json:",omitempty"`
}

// MiddlewareState accessible by other middlewares and endpoint handler
// though the context object.
type MiddlewareState struct {
	Baz string
}

type key string

// MiddlewareStateName
const (
	MiddlewareStateName key = "exampleState"
)

// NewMiddleWare creates a new middleware that executes the next middleware
// after performing it's operations.
func NewMiddleWare(
	gateway *zanzibar.Gateway,
	options Options,
	next zanzibar.HandlerFn) zanzibar.HandlerFn {
	return func(ctx context.Context,
		req *zanzibar.ServerHTTPRequest,
		res *zanzibar.ServerHTTPResponse) {
		scopedCtx := context.WithValue(
			ctx,
			MiddlewareStateName,
			MiddlewareState{
				Baz: options.Foo,
			})
		next(scopedCtx, req, res)
	}
}

// JSONSchema returns a schema definition of the configuration options for a middlware
func JSONSchema() *jsonschema.Document {
	s := &jsonschema.Document{}
	s.Read(&Options{})
	return s
}
