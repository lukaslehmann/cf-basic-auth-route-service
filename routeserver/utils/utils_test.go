package utils_test

import (
	"github.com/benlaplanche/cf-basic-auth-route-service/routeserver/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	var simpleURL string

	BeforeEach(func() {
		simpleURL = "https://myapp.pcf.io"
	})

	Describe(".Strip Special characters and reverse the string", func() {
		Context(".without special characters", func() {
			It("should correctly reverse the string", func() {
				result := utils.StripAndReverse(simpleURL)

				Expect(result).To(Equal("ppaym"))
			})
		})
	})
})
