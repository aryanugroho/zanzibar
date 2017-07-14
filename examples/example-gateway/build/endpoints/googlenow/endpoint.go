// Code generated by zanzibar
// @generated

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

package googlenowEndpoint

import (
	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/googlenow/module"
	zanzibar "github.com/uber/zanzibar/runtime"
)

type Endpoint interface {
	Register(*zanzibar.Gateway) error
}

// NewEndpoint returns a collection of endpoints that can be registered on
// a gateway
func NewEndpoint(g *zanzibar.Gateway, deps *module.Dependencies) Endpoint {
	return &EndpointHandlers{
		AddCredentialsHandler:   NewAddCredentialsHandler(g, deps),
		CheckCredentialsHandler: NewCheckCredentialsHandler(g, deps),
	}
}

// EndpointHandlers is a collection of individual endpoint handlers
type EndpointHandlers struct {
	AddCredentialsHandler   *AddCredentialsHandler
	CheckCredentialsHandler *CheckCredentialsHandler
}

// Register registers the endpoint handlers with the gateway
func (handlers *EndpointHandlers) Register(gateway *zanzibar.Gateway) error {
	err0 := handlers.AddCredentialsHandler.Register(gateway)
	if err0 != nil {
		return err0
	}
	err1 := handlers.CheckCredentialsHandler.Register(gateway)
	if err1 != nil {
		return err1
	}
	return nil
}
