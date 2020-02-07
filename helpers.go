package goqradar

import (
	"context"
	"net/http"
)

// Do the request
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	// Do the request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return nil, nil
}
