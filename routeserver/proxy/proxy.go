package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func New() http.Handler {
	p := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			forwardedURL := req.Header.Get("X-CF-Forwarded-URL")
			url, err := url.Parse(forwardedURL)
			if err != nil {
				log.Fatalln(err.Error())
			}

			req.URL = url
			req.Host = url.Host
		},
	}

	return p
}
