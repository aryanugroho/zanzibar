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

package bazEndpoint

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
	endpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/baz/baz"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
)

// CallHandler is the handler for "/baz/call"
type CallHandler struct {
	Clients *module.ClientDependencies
}

// NewCallHandler} creates a handler
func NewCallHandler(
	gateway *zanzibar.Gateway,
	deps *module.Dependencies,
) *CallHandler {
	return &CallHandler{
		Clients: &deps.Client,
	}
}

// Register adds the http handler to the gateway's http router
func (handler *CallHandler) Register(g *zanzibar.Gateway) error {
	g.HTTPRouter.Register(
		"POST", "/baz/call",
		zanzibar.NewRouterEndpoint(
			g,
			"baz",
			"call",
			handler.HandleRequest,
		),
	)
	// TODO: register should return errors on route conflicts
	return nil
}

// HandleRequest handles "/baz/call".
func (handler *CallHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	if !req.CheckHeaders([]string{"x-uuid", "x-token"}) {
		return
	}
	var requestBody endpointsBazBaz.SimpleService_Call_Args
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	workflow := CallEndpoint{
		Clients: handler.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	cliRespHeaders, err := workflow.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		switch errValue := err.(type) {

		case *endpointsBazBaz.AuthErr:
			res.WriteJSON(
				403, cliRespHeaders, errValue,
			)
			return

		default:
			req.Logger.Warn("Workflow for endpoint returned error",
				zap.String("error", errValue.Error()),
			)
			res.SendErrorString(500, "Unexpected server error")
			return
		}
	}
	// TODO(sindelar): implement check headers on response

	res.WriteJSONBytes(204, cliRespHeaders, nil)
}

// CallEndpoint calls thrift client Baz.Call
type CallEndpoint struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w CallEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBazBaz.SimpleService_Call_Args,
) (zanzibar.Header, error) {
	clientRequest := convertToCallClientRequest(r)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	h, ok = reqHeaders.Get("X-Token")
	if ok {
		clientHeaders["X-Token"] = h
	}
	h, ok = reqHeaders.Get("X-Uuid")
	if ok {
		clientHeaders["X-Uuid"] = h
	}

	cliRespHeaders, err := w.Clients.Baz.Call(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBazBaz.AuthErr:
			serverErr := convertCallAuthErr(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, serverErr

		default:
			w.Logger.Warn("Could not make client request",
				zap.String("error", errValue.Error()),
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	resHeaders.Set("Some-Res-Header", cliRespHeaders["Some-Res-Header"])

	return resHeaders, nil
}

func convertToCallClientRequest(in *endpointsBazBaz.SimpleService_Call_Args) *clientsBazBaz.SimpleService_Call_Args {
	out := &clientsBazBaz.SimpleService_Call_Args{}

	if in.Arg != nil {
		out.Arg = &clientsBazBaz.BazRequest{}
		out.Arg.B1 = bool(in.Arg.B1)
		out.Arg.S2 = string(in.Arg.S2)
		out.Arg.I3 = int32(in.Arg.I3)
	} else {
		out.Arg = nil
	}

	return out
}

func convertCallAuthErr(
	clientError *clientsBazBaz.AuthErr,
) *endpointsBazBaz.AuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.AuthErr{}
	return serverError
}
