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

package barEndpoint

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/thriftrw/ptr"
	"go.uber.org/zap"

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
)

// ArgWithQueryHeaderHandler is the handler for "/bar/argWithQueryHeader"
type ArgWithQueryHeaderHandler struct {
	Clients *module.ClientDependencies
}

// NewArgWithQueryHeaderHandler} creates a handler
func NewArgWithQueryHeaderHandler(
	gateway *zanzibar.Gateway,
	deps *module.Dependencies,
) *ArgWithQueryHeaderHandler {
	return &ArgWithQueryHeaderHandler{
		Clients: &deps.Client,
	}
}

// Register adds the http handler to the gateway's http router
func (handler *ArgWithQueryHeaderHandler) Register(g *zanzibar.Gateway) error {
	g.HTTPRouter.Register(
		"GET", "/bar/argWithQueryHeader",
		zanzibar.NewRouterEndpoint(
			g,
			"bar",
			"argWithQueryHeader",
			handler.HandleRequest,
		),
	)
	// TODO: register should return errors on route conflicts
	return nil
}

// HandleRequest handles "/bar/argWithQueryHeader".
func (handler *ArgWithQueryHeaderHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	var requestBody endpointsBarBar.Bar_ArgWithQueryHeader_Args

	xUUIDValue, _ := req.Header.Get("x-uuid")
	requestBody.UserUUID = ptr.String(xUUIDValue)

	workflow := ArgWithQueryHeaderEndpoint{
		Clients: handler.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, cliRespHeaders, err := workflow.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		switch errValue := err.(type) {

		default:
			req.Logger.Warn("Workflow for endpoint returned error",
				zap.String("error", errValue.Error()),
			)
			res.SendErrorString(500, "Unexpected server error")
			return
		}
	}
	// TODO(jakev): implement writing fields into response headers

	res.WriteJSON(200, cliRespHeaders, response)
}

// ArgWithQueryHeaderEndpoint calls thrift client Bar.ArgWithQueryHeader
type ArgWithQueryHeaderEndpoint struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w ArgWithQueryHeaderEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBarBar.Bar_ArgWithQueryHeader_Args,
) (*endpointsBarBar.BarResponse, zanzibar.Header, error) {
	clientRequest := convertToArgWithQueryHeaderClientRequest(r)

	clientHeaders := map[string]string{}

	clientRespBody, _, err := w.Clients.Bar.ArgWithQueryHeader(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		default:
			w.Logger.Warn("Could not make client request",
				zap.String("error", errValue.Error()),
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertArgWithQueryHeaderClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToArgWithQueryHeaderClientRequest(in *endpointsBarBar.Bar_ArgWithQueryHeader_Args) *clientsBarBar.Bar_ArgWithQueryHeader_Args {
	out := &clientsBarBar.Bar_ArgWithQueryHeader_Args{}

	out.UserUUID = (*string)(in.UserUUID)

	return out
}

func convertArgWithQueryHeaderClientResponse(in *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
	out := &endpointsBarBar.BarResponse{}

	out.StringField = string(in.StringField)
	out.IntWithRange = int32(in.IntWithRange)
	out.IntWithoutRange = int32(in.IntWithoutRange)
	out.MapIntWithRange = make(map[string]int32, len(in.MapIntWithRange))
	for key, value := range in.MapIntWithRange {
		out.MapIntWithRange[key] = int32(value)
	}
	out.MapIntWithoutRange = make(map[string]int32, len(in.MapIntWithoutRange))
	for key, value := range in.MapIntWithoutRange {
		out.MapIntWithoutRange[key] = int32(value)
	}

	return out
}
