package goqradar

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ListOffenses returns the offenses with given fields and filters.
func (service *SIEMService) ListOffenses(ctx context.Context, fields, filters string, min, max int) ([]*Offense, int, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(service.client.BaseURL)
	if err != nil {
		return nil, 0, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/api/siem/offenses"
	parameters := url.Values{}

	if fields != "" {
		parameters.Add("fields", fields)
	}
	if filters != "" {
		parameters.Add("filter", filters)
	}
	reqURL.RawQuery = parameters.Encode()

	// Create the request
	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return nil, 0, fmt.Errorf("error while creating the request : %s", err)
	}
	req = req.WithContext(ctx)

	// Set HTTP headers
	req.Header.Set("SEC", service.client.Token)
	req.Header.Set("Version", service.client.Version)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Range", fmt.Sprintf("items=%d-%d", min, max))

	// Do the request
	resp, err := service.client.client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("error while doing the request : %s", err)
	}
	defer resp.Body.Close()

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("error while reading the request : %s", err)
	}

	// Process the Content-Range
	contentRange := resp.Header.Get("Content-Range")
	contentRangeSplit := strings.Split(contentRange, " ")
	if len(contentRangeSplit) < 2 {
		return nil, 0, fmt.Errorf("Error with the Content-Range")
	}
	contentRangeSplit2 := strings.Split(contentRangeSplit[1], "/")
	total, err := strconv.Atoi(contentRangeSplit2[1])
	if err != nil {
		return nil, 0, fmt.Errorf("error while converting the total into int")
	}

	// Prepare the response
	var offenses []*Offense

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &offenses)
	if err != nil {
		return nil, 0, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return offenses, total, nil
}
