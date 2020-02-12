package goqradar

import (
	"net/http"
	"net/url"
)

type options struct {
	Method   string
	Endpoint string
	Headers  *http.Header
	Params   *url.Values
	Data     map[string]interface{}
	Result   interface{}
}

// Option adds a new option to options.
type Option func(*options) error

// WithHeader adds header.
func WithHeader(key, value string) Option {
	return func(opts *options) error {
		if opts.Headers == nil {
			opts.Headers = &http.Header{}
		}

		opts.Headers.Add(key, value)
		return nil
	}
}

// WithParam adds parameter.
func WithParam(key, value string) Option {
	return func(opts *options) error {
		if opts.Params == nil {
			opts.Params = &url.Values{}
		}

		opts.Params.Add(key, value)
		return nil
	}
}
