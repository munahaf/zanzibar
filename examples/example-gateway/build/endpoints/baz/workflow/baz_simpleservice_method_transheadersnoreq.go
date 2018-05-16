// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
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

package workflow

import (
	"context"
	"strconv"

	zanzibar "github.com/uber/zanzibar/runtime"

	clientsBazBase "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
	endpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/baz/baz"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
	"go.uber.org/zap"
)

// SimpleServiceTransHeadersNoReqWorkflow defines the interface for SimpleServiceTransHeadersNoReq workflow
type SimpleServiceTransHeadersNoReqWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
	) (*endpointsBazBaz.TransHeader, zanzibar.Header, error)
}

// NewSimpleServiceTransHeadersNoReqWorkflow creates a workflow
func NewSimpleServiceTransHeadersNoReqWorkflow(deps *module.Dependencies) SimpleServiceTransHeadersNoReqWorkflow {
	return &simpleServiceTransHeadersNoReqWorkflow{
		Clients: deps.Client,
		Logger:  deps.Default.Logger,
	}
}

// simpleServiceTransHeadersNoReqWorkflow calls thrift client Baz.TransHeadersNoReq
type simpleServiceTransHeadersNoReqWorkflow struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
}

// Handle calls thrift client.
func (w simpleServiceTransHeadersNoReqWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
) (*endpointsBazBaz.TransHeader, zanzibar.Header, error) {
	clientRequest := propagateHeadersTransHeadersNoReqClientRequests(nil, reqHeaders)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	h, ok = reqHeaders.Get("B3")
	if ok {
		clientHeaders["B3"] = h
	}
	h, ok = reqHeaders.Get("I2")
	if ok {
		clientHeaders["I2"] = h
	}
	h, ok = reqHeaders.Get("S1")
	if ok {
		clientHeaders["S1"] = h
	}
	h, ok = reqHeaders.Get("X-Zanzibar-Use-Staging")
	if ok {
		clientHeaders["X-Zanzibar-Use-Staging"] = h
	}

	clientRespBody, _, err := w.Clients.Baz.TransHeadersNoReq(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBazBaz.AuthErr:
			serverErr := convertTransHeadersNoReqAuthErr(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, serverErr

		default:
			w.Logger.Warn("Could not make client request",
				zap.Error(errValue),
				zap.String("client", "Baz"),
			)

			// TODO(sindelar): Consider returning partial headers

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertSimpleServiceTransHeadersNoReqClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertTransHeadersNoReqAuthErr(
	clientError *clientsBazBaz.AuthErr,
) *endpointsBazBaz.AuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.AuthErr{}
	return serverError
}

func convertSimpleServiceTransHeadersNoReqClientResponse(in *clientsBazBase.TransHeaders) *endpointsBazBaz.TransHeader {
	out := &endpointsBazBaz.TransHeader{}

	return out
}

func propagateHeadersTransHeadersNoReqClientRequests(in *clientsBazBaz.SimpleService_TransHeadersNoReq_Args, headers zanzibar.Header) *clientsBazBaz.SimpleService_TransHeadersNoReq_Args {
	if in == nil {
		in = &clientsBazBaz.SimpleService_TransHeadersNoReq_Args{}
	}
	if key, ok := headers.Get("b3"); ok {
		if v, err := strconv.ParseBool(key); err == nil {
			in.B4 = &v
		}

	}
	if key, ok := headers.Get("i2"); ok {
		if v, err := strconv.ParseInt(key, 10, 32); err == nil {
			val := int32(v)
			in.I3 = val
		}

	}
	if key, ok := headers.Get("i2"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBase.NestedStruct{}
		}
		if v, err := strconv.ParseInt(key, 10, 32); err == nil {
			val := int32(v)
			in.Req.Check = &val
		}

	}
	if key, ok := headers.Get("s1"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBase.NestedStruct{}
		}
		in.Req.Msg = key

	}
	if key, ok := headers.Get("s1"); ok {
		in.S2 = &key

	}
	return in
}
