// Code generated by zanzibar
// @generated

package googlenowClient

import (
	"bytes"
	"context"
	"net/http"
	"strconv"

	"github.com/uber/zanzibar/runtime"
)

// GoogleNowClient is the http client for service GoogleNow.
type GoogleNowClient struct {
	client *zanzibar.HTTPClient
}

// NewClient returns a new http client for service GoogleNow.
func NewClient(
	config *zanzibar.StaticConfig,
	gateway *zanzibar.Gateway,
) *GoogleNowClient {
	ip := config.MustGetString("clients.googleNow.ip")
	port := config.MustGetInt("clients.googleNow.port")

	baseURL := "http://" + ip + ":" + strconv.Itoa(int(port))
	return &GoogleNowClient{
		client: zanzibar.NewHTTPClient(gateway, baseURL),
	}
}

// AddCredentials calls "/add-credentials" endpoint.
func (c *GoogleNowClient) AddCredentials(ctx context.Context, r *AddCredentialsHTTPRequest) (*http.Response, error) {
	// Generate full URL.
	fullURL := c.client.BaseURL + "/add-credentials"

	rawBody, err := r.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fullURL, bytes.NewReader(rawBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.client.Client.Do(req.WithContext(ctx))
}

// CheckCredentials calls "/check-credentials" endpoint.
func (c *GoogleNowClient) CheckCredentials(ctx context.Context) (*http.Response, error) {
	// Generate full URL.
	fullURL := c.client.BaseURL + "/check-credentials"

	req, err := http.NewRequest("POST", fullURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.client.Client.Do(req.WithContext(ctx))
}
