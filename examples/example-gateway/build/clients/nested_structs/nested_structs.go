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

package nestedStructsClient

import (
	"context"
	"strconv"

	"github.com/pkg/errors"
	clientsNestedStructsNestedStructs "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/nested_structs/nested_structs"
	"github.com/uber/zanzibar/runtime"
)

// NestedStructsClient is the http client for service NestedStructs.
type NestedStructsClient struct {
	ClientID   string
	HTTPClient *zanzibar.HTTPClient
}

// NewClient returns a new http client for service NestedStructs.
func NewClient(
	gateway *zanzibar.Gateway,
) *NestedStructsClient {
	ip := gateway.Config.MustGetString("clients.nestedStructs.ip")
	port := gateway.Config.MustGetInt("clients.nestedStructs.port")

	baseURL := "http://" + ip + ":" + strconv.Itoa(int(port))
	return &NestedStructsClient{
		ClientID:   "nestedStructs",
		HTTPClient: zanzibar.NewHTTPClient(gateway, baseURL),
	}
}

// First calls "/first" endpoint.
func (c *NestedStructsClient) First(
	ctx context.Context,
	headers map[string]string,
	r *clientsNestedStructsNestedStructs.NestedStructs_First_Args,
) (*clientsNestedStructsNestedStructs.RequiredStruct, map[string]string, error) {

	req := zanzibar.NewClientHTTPRequest(
		c.ClientID, "first", c.HTTPClient,
	)

	// Generate full URL.
	fullURL := c.HTTPClient.BaseURL + "/first"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return nil, nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return nil, nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200})

	switch res.StatusCode {
	case 200:
		var responseBody clientsNestedStructsNestedStructs.RequiredStruct
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return nil, respHeaders, err
		}

		return &responseBody, respHeaders, nil
	}

	return nil, respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}

// Second calls "/second" endpoint.
func (c *NestedStructsClient) Second(
	ctx context.Context,
	headers map[string]string,
	r *clientsNestedStructsNestedStructs.NestedStructs_Second_Args,
) (*clientsNestedStructsNestedStructs.OptionalStruct, map[string]string, error) {

	req := zanzibar.NewClientHTTPRequest(
		c.ClientID, "second", c.HTTPClient,
	)

	// Generate full URL.
	fullURL := c.HTTPClient.BaseURL + "/second"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return nil, nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return nil, nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200})

	switch res.StatusCode {
	case 200:
		var responseBody clientsNestedStructsNestedStructs.OptionalStruct
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return nil, respHeaders, err
		}

		return &responseBody, respHeaders, nil
	}

	return nil, respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}
