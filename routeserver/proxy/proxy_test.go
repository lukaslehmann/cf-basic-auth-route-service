package proxy_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Proxy", func() {
	var (
		transport        http.RoundTripper
		req              *http.Request
		helloWorldServer *ghttp.Server
	)

	BeforeEach(func() {
		helloWorldServer = ghttp.NewServer()
		helloWorldServer.AppendHandlers(ghttp.RespondWith(200, []byte{"Hello World!"}))

		req, _ := http.NewRequest("GET", helloWorldServer.URL(), nil)
		transport = NewBasicAuthTransport()
		req.Header.Add("X-CF-Forwarded-Url", "https://hello-world.com")
		req.Header.Add("X-CF-Proxy-Metadata", "header-proxy-metadata-goes-here")
		req.Header.Add("X-CF-Proxy-Signature", "header-proxy-signature-goes-here")
	})

	Context("Missing the correct headers", func() {
		It("returns the correct error when there is no forwarded url", func() {
		})

		It("return the correct error when there is no proxy metadata", func() {
		})

		It("returns the correct error when there is no proxy signature", func() {
		})
	})

	Context("Contains the correct headers", func() {
		Context("With no access details", func() {
			It("returns the correct HTTP Status code", func() {
			})

			It("returns the expected response body", func() {
			}

		})

		Context("With invalid login details", func() {
			It("returns the correct HTTP Status code", func() {
			})

			It("returns the expected response body", func() {
			})
		})

		Context("With valid login details", func() {
			It("returns the correct HTTP Status code", func() {
			})

			It("returns the expected response body", func() {
			})

		})
	})

})
