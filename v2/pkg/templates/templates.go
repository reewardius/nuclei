package templates

import (
	"github.com/projectdiscovery/nuclei/v2/pkg/requests"
)

// Template is a request template parsed from a yaml file
type Template struct {
	// ID is the unique id for the template
	ID string `yaml:"id"`
	// Info contains information about the template
	Info map[string]string `yaml:"info"`
	// RequestsHTTP contains the http request to make in the template
	RequestsHTTP []*requests.BulkHTTPRequest `yaml:"requests,omitempty"`
	// RequestsDNS contains the dns request to make in the template
	RequestsDNS []*requests.DNSRequest `yaml:"dns,omitempty"`
	// RequestsNetwork contains the network request to make in the template
	RequestsNetwork []*requests.NetworkRequest `yaml:"network,omitempty"`

	path string
}

// GetPath of the workflow
func (t *Template) GetPath() string {
	return t.path
}

// GetHTTPRequestCount returns count of HTTP requests to make for template
func (t *Template) GetHTTPRequestCount() int64 {
	var count int64
	for _, request := range t.RequestsHTTP {
		count += request.GetRequestCount()
	}
	return count
}

// GetDNSRequestCount returns count of DNS requests to make for template
func (t *Template) GetDNSRequestCount() int64 {
	var count int64
	for _, request := range t.RequestsDNS {
		count += request.GetRequestCount()
	}
	return count
}

// GetNetworkRequestCount returns count of Network requests to make for template
func (t *Template) GetNetworkRequestCount() int64 {
	var count int64
	for _, request := range t.RequestsNetwork {
		count += request.GetRequestCount()
	}
	return count
}
