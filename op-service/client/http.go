package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

const (
	DefaultTimeoutSeconds = 30
)

var _ HTTP = (*BasicHTTPClient)(nil)

type HTTP interface {
	Get(ctx context.Context, path string, query url.Values, headers http.Header) (*http.Response, error)
}

type BasicHTTPClient struct {
	endpoint string
	log      log.Logger
	client   *http.Client
}

func NewBasicHTTPClient(endpoint string, log log.Logger) *BasicHTTPClient {
	// Make sure the endpoint ends in trailing slash
	trimmedEndpoint := strings.TrimSuffix(endpoint, "/") + "/"
	return &BasicHTTPClient{
		endpoint: trimmedEndpoint,
		log:      log,
		client:   &http.Client{Timeout: DefaultTimeoutSeconds * time.Second},
	}
}

func (cl *BasicHTTPClient) Get(ctx context.Context, p string, query url.Values, headers http.Header) (*http.Response, error) {
	target, err := url.Parse(cl.endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse endpoint URL: %w", err)
	}
	// If we include the raw query in the path-join, it gets url-encoded,
	// and fails to prase as query, and ends up in the url.URL.Path part on the server side.
	// We want to avoid that, and insert the query manually. Real footgun in the url package.
	target = target.JoinPath(p)
	target.RawQuery = query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to construct request", err)
	}
	for k, values := range headers {
		for _, v := range values {
			req.Header.Add(k, v)
		}
	}
	return cl.client.Do(req)
}
