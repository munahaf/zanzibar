package googlenowClient

import (
	"bytes"
	"context"
	"net/http"
	"strconv"

	"github.com/uber/zanzibar/runtime"
)

// GoogleNowClient is the http client for service GoogleNow.
type GoogleNowClient zanzibar.HTTPClient

// NewClient returns a new http client for service GoogleNow.
func NewClient(opts *zanzibar.HTTPClientOptions) *GoogleNowClient {
	baseURL := "http://" + opts.IP + ":" + strconv.Itoa(int(opts.Port))
	return &GoogleNowClient{
		Client: &http.Client{
			Transport: &http.Transport{
				DisableKeepAlives:   false,
				MaxIdleConns:        500,
				MaxIdleConnsPerHost: 500,
			},
		},
		BaseURL: baseURL,
	}
}

// AddCredentials calls "/add-credentials" endpoint.
func (c *GoogleNowClient) AddCredentials(ctx context.Context, r *AddCredentialsHTTPRequest, h http.Header) (*http.Response, error) {
	// Generate full URL.
	fullURL := c.BaseURL + "/add-credentials"

	rawBody, err := r.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fullURL, bytes.NewReader(rawBody))
	if err != nil {
		return nil, err
	}
	if h != nil {
		req.Header = h
	}
	req.Header.Set("Content-Type", "application/json")
	return c.Client.Do(req.WithContext(ctx))
}

// CheckCredential calls "/check-credentials" endpoint.
func (c *GoogleNowClient) CheckCredential(ctx context.Context, h http.Header) (*http.Response, error) {
	// Generate full URL.
	fullURL := c.BaseURL + "/check-credentials"

	req, err := http.NewRequest("POST", fullURL, nil)
	if err != nil {
		return nil, err
	}
	if h != nil {
		req.Header = h
	}
	req.Header.Set("Content-Type", "application/json")
	return c.Client.Do(req.WithContext(ctx))
}
