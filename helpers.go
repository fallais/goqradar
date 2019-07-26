package qradar

import (
	"context"
	"net/http"
)

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	// Do the request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
}
