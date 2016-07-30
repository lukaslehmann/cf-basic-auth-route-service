package proxy

import (
	"crypto/tls"
	"net/http"
)

type BasicAuthTransport struct {
	transport http.RoundTripper
}

func NewBasicAuthTransport() http.RoundTripper {
	return &BasicAuthTransport{
		transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

func (t *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, nil
}
