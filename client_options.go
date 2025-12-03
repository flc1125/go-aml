package aml

import "net/http"

type ClientOption func(*Client) error

// WithBaseURL sets the baseURL for the client
func WithBaseURL(urlStr string) ClientOption {
	return func(c *Client) error {
		return c.setBaseURL(urlStr)
	}
}

// WithAuth sets the account and password for the client
func WithAuth(account, password string) ClientOption {
	return func(c *Client) error {
		c.account = account
		c.password = password
		return nil
	}
}

// WithUserAgent sets the userAgent for the client
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.userAgent = userAgent
		return nil
	}
}

// WithHTTPClient sets the httpClient for the client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) error {
		c.httpClient = httpClient
		return nil
	}
}
