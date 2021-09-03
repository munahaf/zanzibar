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

package corgeclient

import (
	"context"
	"errors"
	"net/textproto"
	"strconv"
	"strings"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/uber/tchannel-go"
	"github.com/uber/zanzibar/config"
	zanzibar "github.com/uber/zanzibar/runtime"
	"github.com/uber/zanzibar/runtime/ruleengine"

	"go.uber.org/zap"

	module "github.com/uber/zanzibar/examples/example-gateway/build/clients/corge/module"
	clientsIDlClientsCorgeCorge "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients-idl/clients/corge/corge"
)

// CircuitBreakerConfigKey is key value for qps level to circuit breaker parameters mapping
const CircuitBreakerConfigKey = "circuitbreaking-configurations"

// Client defines corge client interface.
type Client interface {
	EchoString(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsIDlClientsCorgeCorge.Corge_EchoString_Args,
	) (context.Context, string, map[string]string, error)
}

// NewClient returns a new TChannel client for service corge.
func NewClient(deps *module.Dependencies) Client {
	serviceName := deps.Default.Config.MustGetString("clients.corge.serviceName")
	var routingKey string
	if deps.Default.Config.ContainsKey("clients.corge.routingKey") {
		routingKey = deps.Default.Config.MustGetString("clients.corge.routingKey")
	}
	var requestUUIDHeaderKey string
	if deps.Default.Config.ContainsKey("tchannel.clients.requestUUIDHeaderKey") {
		requestUUIDHeaderKey = deps.Default.Config.MustGetString("tchannel.clients.requestUUIDHeaderKey")
	}

	ip := deps.Default.Config.MustGetString("sidecarRouter.default.tchannel.ip")
	port := deps.Default.Config.MustGetInt("sidecarRouter.default.tchannel.port")
	gateway := deps.Default.Gateway
	var channel *tchannel.Channel

	// If dedicated.tchannel.client : true, each tchannel client will create a
	// dedicated connection with local sidecar, else it will use a shared connection
	if deps.Default.Config.ContainsKey("dedicated.tchannel.client") &&
		deps.Default.Config.MustGetBoolean("dedicated.tchannel.client") {
		channel = gateway.SetupClientTChannel(deps.Default.Config, serviceName)
		channel.Peers().Add(ip + ":" + strconv.Itoa(int(port)))
	} else {
		channel = gateway.ServerTChannel
		channel.GetSubChannel(serviceName, tchannel.Isolated).Peers().Add(ip + ":" + strconv.Itoa(int(port)))
	}

	/*Ex:
	{
	  "clients.rider-presentation.alternates": {
		"routingConfigs": [
		  {
			"headerName": "x-test-env",
			"headerValue": "*",
			"serviceName": "testservice"
		  },
		  {
			"headerName": "x-container",
			"headerValue": "container*",
			"serviceName": "relayer"
		  }
		],
		"servicesDetail": {
		  "testservice": {
			"ip": "127.0.0.1",
			"port": 5000
		  },
		  "relayer": {
			"ip": "127.0.0.1",
			"port": 12000
		  }
		}
	  }
	}*/
	var re ruleengine.RuleEngine
	var headerPatterns []string
	altChannelMap := make(map[string]*tchannel.SubChannel)
	headerPatterns, re = initializeDynamicChannel(channel, deps, headerPatterns, altChannelMap, re)

	timeoutVal := int(deps.Default.Config.MustGetInt("clients.corge.timeout"))
	timeout := time.Millisecond * time.Duration(
		timeoutVal,
	)
	timeoutPerAttempt := time.Millisecond * time.Duration(
		deps.Default.Config.MustGetInt("clients.corge.timeoutPerAttempt"),
	)

	methodNames := map[string]string{
		"Corge::echoString": "EchoString",
	}

	qpsLevels := map[string]string{
		"corge-EchoString": "default",
	}

	// circuitBreakerDisabled sets whether circuit-breaker should be disabled
	circuitBreakerDisabled := false
	if deps.Default.Config.ContainsKey("clients.corge.circuitBreakerDisabled") {
		circuitBreakerDisabled = deps.Default.Config.MustGetBoolean("clients.corge.circuitBreakerDisabled")
	}

	if !circuitBreakerDisabled {
		for _, methodName := range methodNames {
			circuitBreakerName := "corge" + "-" + methodName
			qpsLevel := "default"
			if level, ok := qpsLevels[circuitBreakerName]; ok {
				qpsLevel = level
			}
			configureCircuitBreaker(deps, timeoutVal, circuitBreakerName, qpsLevel)
		}
	}

	var client *zanzibar.TChannelClient

	var maxAttempts int
	if deps.Default.Config.ContainsKey("tchannelclients.retryCount.feature.enabled") && deps.Default.Config.MustGetBoolean("tchannelclients.retryCount.feature.enabled") && deps.Default.Config.ContainsKey("clients.corge.retryCount") {
		maxAttempts = int(deps.Default.Config.MustGetInt("clients.corge.retryCount"))
		client = zanzibar.NewTChannelClientContext(
			channel,
			deps.Default.ContextLogger,
			deps.Default.ContextMetrics,
			deps.Default.ContextExtractor,
			&zanzibar.TChannelClientOption{
				ServiceName:          serviceName,
				ClientID:             "corge",
				MethodNames:          methodNames,
				Timeout:              timeout,
				TimeoutPerAttempt:    timeoutPerAttempt,
				RoutingKey:           &routingKey,
				RuleEngine:           re,
				HeaderPatterns:       headerPatterns,
				RequestUUIDHeaderKey: requestUUIDHeaderKey,
				AltChannelMap:        altChannelMap,
				MaxAttempts:          maxAttempts,
			},
		)
	} else {
		client = zanzibar.NewTChannelClientContext(
			channel,
			deps.Default.ContextLogger,
			deps.Default.ContextMetrics,
			deps.Default.ContextExtractor,
			&zanzibar.TChannelClientOption{
				ServiceName:          serviceName,
				ClientID:             "corge",
				MethodNames:          methodNames,
				Timeout:              timeout,
				TimeoutPerAttempt:    timeoutPerAttempt,
				RoutingKey:           &routingKey,
				RuleEngine:           re,
				HeaderPatterns:       headerPatterns,
				RequestUUIDHeaderKey: requestUUIDHeaderKey,
				AltChannelMap:        altChannelMap,
			},
		)
	}

	return &corgeClient{
		client:                 client,
		circuitBreakerDisabled: circuitBreakerDisabled,
		defaultDeps:            deps.Default,
	}
}

func initializeDynamicChannel(channel *tchannel.Channel, deps *module.Dependencies, headerPatterns []string, altChannelMap map[string]*tchannel.SubChannel, re ruleengine.RuleEngine) ([]string, ruleengine.RuleEngine) {
	if deps.Default.Config.ContainsKey("clients.corge.alternates") {
		var alternateServiceDetail config.AlternateServiceDetail
		deps.Default.Config.MustGetStruct("clients.corge.alternates", &alternateServiceDetail)

		ruleWrapper := ruleengine.RuleWrapper{}
		for _, routingConfig := range alternateServiceDetail.RoutingConfigs {
			ruleValue := []string{routingConfig.ServiceName}
			rd := routingConfig.RoutingDelegate
			if rd != nil {
				ruleValue = append(ruleValue, *rd)
			}
			rawRule := ruleengine.RawRule{Patterns: []string{textproto.CanonicalMIMEHeaderKey(routingConfig.HeaderName),
				strings.ToLower(routingConfig.HeaderValue)}, Value: ruleValue}
			headerPatterns = append(headerPatterns, textproto.CanonicalMIMEHeaderKey(routingConfig.HeaderName))
			ruleWrapper.Rules = append(ruleWrapper.Rules, rawRule)

			scAlt := channel.GetSubChannel(routingConfig.ServiceName, tchannel.Isolated)
			serviceRouting, ok := alternateServiceDetail.ServicesDetailMap[routingConfig.ServiceName]
			if !ok {
				panic("service routing mapping incorrect for service: " + routingConfig.ServiceName)
			}
			scAlt.Peers().Add(serviceRouting.IP + ":" + strconv.Itoa(serviceRouting.Port))
			altChannelMap[routingConfig.ServiceName] = scAlt
		}

		re = ruleengine.NewRuleEngine(ruleWrapper)
	}
	return headerPatterns, re
}

// CircuitBreakerConfig is used for storing the circuit breaker parameters for each qps level
type CircuitBreakerConfig struct {
	Parameters map[string]map[string]int
}

func configureCircuitBreaker(deps *module.Dependencies, timeoutVal int, circuitBreakerName string, qpsLevel string) {
	// sleepWindowInMilliseconds sets the amount of time, after tripping the circuit,
	// to reject requests before allowing attempts again to determine if the circuit should again be closed
	sleepWindowInMilliseconds := 5000
	// maxConcurrentRequests sets how many requests can be run at the same time, beyond which requests are rejected
	maxConcurrentRequests := 20
	// errorPercentThreshold sets the error percentage at or above which the circuit should trip open
	errorPercentThreshold := 20
	// requestVolumeThreshold sets a minimum number of requests that will trip the circuit in a rolling window of 10s
	// For example, if the value is 20, then if only 19 requests are received in the rolling window of 10 seconds
	// the circuit will not trip open even if all 19 failed.
	requestVolumeThreshold := 20
	// parses circuit breaker configurations
	if deps.Default.Config.ContainsKey(CircuitBreakerConfigKey) {
		var config CircuitBreakerConfig
		deps.Default.Config.MustGetStruct(CircuitBreakerConfigKey, &config)
		parameters := config.Parameters
		// first checks if level exists in configurations then assigns parameters
		// if "default" qps level assigns default parameters from circuit breaker configurations
		if settings, ok := parameters[qpsLevel]; ok {
			if sleep, ok := settings["sleepWindowInMilliseconds"]; ok {
				sleepWindowInMilliseconds = sleep
			}
			if max, ok := settings["maxConcurrentRequests"]; ok {
				maxConcurrentRequests = max
			}
			if errorPercent, ok := settings["errorPercentThreshold"]; ok {
				errorPercentThreshold = errorPercent
			}
			if reqVolThreshold, ok := settings["requestVolumeThreshold"]; ok {
				requestVolumeThreshold = reqVolThreshold
			}
		}
	}
	// client settings override parameters
	if deps.Default.Config.ContainsKey("clients.corge.sleepWindowInMilliseconds") {
		sleepWindowInMilliseconds = int(deps.Default.Config.MustGetInt("clients.corge.sleepWindowInMilliseconds"))
	}
	if deps.Default.Config.ContainsKey("clients.corge.maxConcurrentRequests") {
		maxConcurrentRequests = int(deps.Default.Config.MustGetInt("clients.corge.maxConcurrentRequests"))
	}
	if deps.Default.Config.ContainsKey("clients.corge.errorPercentThreshold") {
		errorPercentThreshold = int(deps.Default.Config.MustGetInt("clients.corge.errorPercentThreshold"))
	}
	if deps.Default.Config.ContainsKey("clients.corge.requestVolumeThreshold") {
		requestVolumeThreshold = int(deps.Default.Config.MustGetInt("clients.corge.requestVolumeThreshold"))
	}
	hystrix.ConfigureCommand(circuitBreakerName, hystrix.CommandConfig{
		MaxConcurrentRequests:  maxConcurrentRequests,
		ErrorPercentThreshold:  errorPercentThreshold,
		SleepWindow:            sleepWindowInMilliseconds,
		RequestVolumeThreshold: requestVolumeThreshold,
		Timeout:                timeoutVal,
	})
}

// corgeClient is the TChannel client for downstream service.
type corgeClient struct {
	client                 *zanzibar.TChannelClient
	circuitBreakerDisabled bool
	defaultDeps            *zanzibar.DefaultDependencies
}

// EchoString is a client RPC call for method "Corge::echoString"
func (c *corgeClient) EchoString(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsIDlClientsCorgeCorge.Corge_EchoString_Args,
) (context.Context, string, map[string]string, error) {
	var result clientsIDlClientsCorgeCorge.Corge_EchoString_Result
	var resp string

	logger := c.client.ContextLogger

	var success bool
	respHeaders := make(map[string]string)
	var err error
	if c.circuitBreakerDisabled {
		success, respHeaders, err = c.client.Call(
			ctx, "Corge", "echoString", reqHeaders, args, &result,
		)
	} else {
		// We want hystrix ckt-breaker to count errors only for system issues
		var clientErr error
		scope := c.defaultDeps.Scope.Tagged(map[string]string{
			"client":     "corge",
			"methodName": "EchoString",
		})
		start := time.Now()
		circuitBreakerName := "corge" + "-" + "EchoString"
		err = hystrix.DoC(ctx, circuitBreakerName, func(ctx context.Context) error {
			elapsed := time.Now().Sub(start)
			scope.Timer("hystrix-timer").Record(elapsed)
			success, respHeaders, clientErr = c.client.Call(
				ctx, "Corge", "echoString", reqHeaders, args, &result,
			)
			if _, isSysErr := clientErr.(tchannel.SystemError); !isSysErr {
				// Declare ok if it is not a system-error
				return nil
			}
			return clientErr
		}, nil)
		if err == nil {
			// ckt-breaker was ok, bubble up client error if set
			err = clientErr
		}
	}

	if err == nil && !success {
		switch {
		case result.Success != nil:
			ctx = logger.ErrorZ(ctx, "Internal error. Success flag is not set for EchoString. Overriding", zap.Error(err))
			success = true
		default:
			err = errors.New("corgeClient received no result or unknown exception for EchoString")
		}
	}
	if err != nil {
		ctx = logger.WarnZ(ctx, "Client failure: TChannel client call returned error", zap.Error(err))
		return ctx, resp, respHeaders, err
	}

	resp, err = clientsIDlClientsCorgeCorge.Corge_EchoString_Helper.UnwrapResponse(&result)
	if err != nil {
		ctx = logger.WarnZ(ctx, "Client failure: unable to unwrap client response", zap.Error(err))
	}
	return ctx, resp, respHeaders, err
}
