package goqradar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func parseContentRange(cr string) (int, int, int, error) {
	// Trim
	trimed := strings.Trim(cr, "items ")

	// Split min-max and total
	split := strings.Split(trimed, "/")
	if len(split) != 2 {
		return 0, 0, 0, fmt.Errorf("error when splitting the content-range with slash")
	}

	if split[0] == "*" {
		// Convert total
		total, err := strconv.Atoi(split[1])
		if err != nil {
			return 0, 0, 0, fmt.Errorf("error while converting the total into int")
		}

		return 0, 0, total, nil
	}

	// Split min and max
	minAndMax := strings.Split(split[0], "-")
	if len(minAndMax) != 2 {
		return 0, 0, 0, fmt.Errorf("error when splitting the content-range with dash")
	}

	// Convert min
	min, err := strconv.Atoi(minAndMax[0])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error while converting the min into int")
	}

	// Convert max
	max, err := strconv.Atoi(minAndMax[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error while converting the max into int")
	}

	// Convert total
	total, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error while converting the total into int")
	}

	return min, max, total, nil
}

func (c *Client) do(ctx context.Context, method, endpoint string, opts ...Option) (*http.Response, error) {
	// Options
	var apiOptions options

	// Add options
	for _, op := range opts {
		err := op(&apiOptions)
		if err != nil {
			return nil, err
		}
	}

	// Raw URL
	rawURL := fmt.Sprintf("%s/api/%s", c.BaseURL, endpoint)

	// Build query
	queryURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	// Assign query parameters
	if apiOptions.Params != nil {
		queryURL.RawQuery = apiOptions.Params.Encode()
	}

	// Initialize request
	req, err := http.NewRequestWithContext(ctx, method, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	// Default headers
	headers := http.Header{}
	headers.Add("Accept", "application/json")
	headers.Add("Version", c.Version)
	headers.Add("SEC", c.Token)

	// Optional headers
	if apiOptions.Headers != nil {
		for k := range *apiOptions.Headers {
			headers.Add(k, apiOptions.Headers.Get(k))
		}
	}

	// Assign new headers
	req.Header = headers

	// Do the query
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while doing the request: %s", err)
	}

	return resp, err
}
