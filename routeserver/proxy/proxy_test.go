package proxy_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/benlaplanche/cf-basic-auth-route-service/routeserver/proxy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Proxy", func() {

	const (
		CF_PROXY_SIGNATURE = "5ASjPwv2H3IUO1LzEQYxfH6ceTt_wFGmjG1ESFS4rkAvO1fTBRsVf9QT8pXPg8cRGx4LK1LZWX5WkrT2DB5iKq4w2FM80OoRAcM_LcNz7tRPcniqwMO1adkrvulP2-LuTktyVKN8w2KaPImKkTD7vrnxFA=="
		CF_PROXY_METADATA  = "eyJub25jZSI6IjBxcGdYZmZNVVNQQnZwV3UifQ=="
		CF_FORWARDED_URL   = "https://my-app-1.pcf.io"
	)

	var proxyServer http.Handler

	fakeProtectedApp := func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`hello world`))
		})

		port := os.Getenv("PORT")
		if port == "" {
			port = "9999"
		}

		if err := http.ListenAndServe(CF_FORWARDED_URL+port, nil); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	}

	BeforeEach(func() {
		proxyServer = proxy.New()
		fakeProtectedApp()
	})

	makeRequest := func() *httptest.ResponseRecorder {
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/", nil)

		request.Header.Add("X-CF-Forwarded-Url", CF_FORWARDED_URL)
		request.Header.Add("X-Cf-Proxy-Signature", CF_PROXY_SIGNATURE)
		request.Header.Add("X-Cf-Proxy-Metadata", CF_PROXY_METADATA)

		proxyServer.ServeHTTP(recorder, request)
		return recorder
	}

	Describe("maintains the correct X-CF headers", func() {
		It("should contain the X-CF-Forwarded-Url header", func() {
			response := makeRequest()

			header := response.Header().Get("X-CF-Forwarded-Url")
			Expect(header).To(Equal(CF_FORWARDED_URL))
		})

		It("should contain the X-CF-Proxy-Signature header", func() {
			response := makeRequest()

			header := response.Header().Get("X-Cf-Proxy-Signature")
			Expect(header).To(Equal(CF_PROXY_SIGNATURE))
		})

		It("should contain the X-CF-Proxy-Metadata header", func() {
			response := makeRequest()

			header := response.Header().Get("X-Cf-Proxy-Metadata")
			Expect(header).To(Equal(CF_PROXY_METADATA))
		})
	})

})
