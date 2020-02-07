package goqradar

// Options are the configurable options for the QRadar client.
type Options struct {
	version  string
	timeout  int
	insecure bool
}

// DefaultOptions are the default options.
var DefaultOptions = Options{
	version:  "9.0",
	timeout:  10,
	insecure: false,
}
