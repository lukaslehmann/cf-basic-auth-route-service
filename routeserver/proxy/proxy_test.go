package proxy_test

import (
	"net/http"

	"github.com/benlaplanche/cf-basic-auth-route-service/routeserver/proxy"
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
		helloWorldServer.AppendHandlers(ghttp.RespondWith(200, []byte("Hello World!")))

		req, _ := http.NewRequest("GET", helloWorldServer.URL(), nil)
		transport = proxy.NewBasicAuthTransport()
		req.Header.Add("X-CF-Forwarded-Url", "https://hello-world.com")
		req.Header.Add("X-CF-Proxy-Metadata", "header-proxy-metadata-goes-here")
		req.Header.Add("X-CF-Proxy-Signature", "header-proxy-signature-goes-here")
	})

	Context("Missing the correct headers", func() {
		It("returns the correct error when there is no forwarded url", func() {
			req.Header.Del("X-CF-Forwarded-Url")

			req, res := transport.RoundTrip(req)
			Expect(req).To(BeNil())
			Expect(res).ToNot(BeNil())
		})

		It("return the correct error when there is no proxy metadata", func() {
			req.Header.Del("X-CF-Proxy-Metadata")

			req, res := transport.RoundTrip(req)
			Expect(req).To(BeNil())
			Expect(res).ToNot(BeNil())
		})

		It("returns the correct error when there is no proxy signature", func() {
			req.Header.Del("X-CF-Proxy-Signature")

			req, res := transport.RoundTrip(req)
			Expect(req).To(BeNil())
			Expect(res).ToNot(BeNil())
		})
	})

	PContext("Contains the correct headers", func() {
		Context("With no access details", func() {
			It("returns the correct HTTP Status code", func() {
			})

			It("returns the expected response body", func() {
			})

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
