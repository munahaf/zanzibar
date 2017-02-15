package bar

import (
	"github.com/uber/zanzibar/examples/example-gateway/gen-code/github.com/uber/zanzibar/endpoints/bar/bar"
	"github.com/uber/zanzibar/examples/example-gateway/gen-code/github.com/uber/zanzibar/endpoints/foo/foo"
)

// ArgNotStructHTTPRequest is the http body type for endpoint argNotStruct.
type ArgNotStructHTTPRequest struct {
	Request string
}

// NormalHTTPRequest is the http body type for endpoint normal.
type NormalHTTPRequest struct {
	Request bar.BarRequest
}

// TooManyArgsHTTPRequest is the http body type for endpoint tooManyArgs.
type TooManyArgsHTTPRequest struct {
	Request bar.BarRequest
	Foo     foo.FooStruct
}
