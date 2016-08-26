package proxy

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type BasicAuthTransport struct {
	transport http.RoundTripper
}

func NewBasicAuthTransport(skipSSLValidation bool) http.RoundTripper {
	return &BasicAuthTransport{
		transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: skipSSLValidation},
		},
	}
}

func (b *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	err := checkHeaders(req)
	if err != nil {
		log.Printf("Invalid headers. %+v\n", req.Header)
		return nil, err
	}

	return nil, nil
}

func checkHeaders(r *http.Request) error {
	if r.Header.Get("X-CF-Forwarded-Url") == "" {
		return missingHeaderError("X-CF-Forwarded-Url")
	}

	if r.Header.Get("X-CF-Proxy-Metadata") == "" {
		return missingHeaderError("X-CF-Proxy-Metadata")
	}

	if r.Header.Get("X-CF-Proxy-Signature") == "" {
		return missingHeaderError("X-CF-Proxy-Signature")
	}

	return nil
}

func missingHeaderError(header string) error {
	return errors.New(fmt.Sprintf("Missing expected header: %s", header))
}
