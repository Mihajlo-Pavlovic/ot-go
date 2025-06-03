package triplestore

import (
	"bytes"
	"io"
	"net/http"
)


type Client struct {
	endpoint string
	headers  map[string]string
	http     *http.Client
}

func New(endpoint string, headers map[string]string) *Client {
	return &Client{
		endpoint: endpoint,
		headers:  headers,
		http:     &http.Client{},
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
