package triplestore

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

type Client struct {
	endpoint string
	headers  map[string]string
	http     *http.Client
}

func New(endpoint string, headers map[string]string) *Client {
	transport := &http.Transport{
		MaxIdleConns:        100,              // Total idle connections
		MaxIdleConnsPerHost: 20,               // Per host
		IdleConnTimeout:     90 * time.Second, // Keep alive time
		DisableKeepAlives:   false,            // Enable keep-alive
	}
	return &Client{
		endpoint: endpoint,
		headers:  headers,
		http:     &http.Client{Transport: transport, Timeout: 30 * time.Second},
	}
}

func (c *Client) Query(query string) ([]byte, error) {
	req, err := http.NewRequest("POST", c.endpoint, bytes.NewBufferString(query))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/sparql-query")
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
