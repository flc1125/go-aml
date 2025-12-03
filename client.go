package aml

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL   = "https://aml.tdcc.com.tw/"
	defaultUserAgent = "go-aml"
)

var defaultHTTPClient = NewRetryableHTTPClient()

type Client struct {
	// baseURL for API requests.
	baseURL *url.URL

	// account, password for authentication.
	account, password string

	// accessToken for Personal Access Token (PAT) authentication.
	accessToken string

	// userAgent used for HTTP requests
	userAgent string

	// httpClient is the HTTP client used to communicate with the API.
	httpClient *http.Client

	// services used for talking to different parts of the Tapd API.
}

// NewClient returns a new AML API client with authentication.
func NewClient(account, password string, opts ...ClientOption) (*Client, error) {
	return newClient(append(opts, WithAuth(account, password))...)
}

// newClient returns a new AML API client.
func newClient(opts ...ClientOption) (*Client, error) {
	c := &Client{
		userAgent:  defaultUserAgent,
		httpClient: defaultHTTPClient,
	}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	// setup
	if err := c.setup(); err != nil {
		return nil, err
	}

	return c, nil
}

// setup sets up the client for API requests.
func (c *Client) setup() error {
	if c.baseURL == nil {
		if err := c.setBaseURL(defaultBaseURL); err != nil {
			return err
		}
	}

	return nil
}

// setBaseURL sets the base URL for API requests to a custom endpoint.
func (c *Client) setBaseURL(urlStr string) error {
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	c.baseURL = baseURL

	return nil
}

func (c *Client) NewRequest(ctx context.Context, method, path string, data any, opts []RequestOption) (*http.Request, error) { //nolint:lll
	u := *c.baseURL
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	// Set the encoded path data
	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	if c.userAgent != "" {
		reqHeaders.Set("User-Agent", c.userAgent)
	}

	var body io.Reader
	switch {
	case method == http.MethodPatch || method == http.MethodPost || method == http.MethodPut:
		reqHeaders.Set("Content-Type", "application/json")

		b, err := json.Marshal(reqRaw{
			UserID:   c.account,
			Password: c.password,
			Target:   data,
		})
		if err != nil {
			return nil, err
		}
		body = io.NopCloser(bytes.NewReader(b))

	default:
		q := make(url.Values)
		rawEncoder := reqRaw{
			UserID:   c.account,
			Password: c.password,
			Target:   data,
		}
		if err := rawEncoder.EncodeValues("", &q); err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}

	// Set the request specific headers.
	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	// Apply request options
	for _, opt := range opts {
		if opt != nil {
			if err := opt(req); err != nil {
				return nil, err
			}
		}
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v any) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()              // nolint:errcheck
	defer io.Copy(io.Discard, resp.Body) // nolint:errcheck

	// decode response body
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}

	return newResponse(resp), err
}

func par