package utils_test

import (
	"github.com/benlaplanche/cf-basic-auth-route-service/routeserver/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	var example1 string = "https://myapp.pcf.io"
	var example2 string = "http://my-app.pcf.io"
	var example3 string = "https://my-app-1.pcf.io"

	Describe(".Strip Special characters and reverse the string", func() {
		Context(".without special characters", func() {
			It("should correctly reverse the string", func() {
				result := utils.StripAndReverse(example1)

				Expect(result).To(Equal("ppaym"))
			})

			It("should work with numbers", func() {
				result := utils.StripAndReverse(example3)
				Expect(result).To(Equal("1ppaym"))
			})

		})

		Context(".with special characters", func() {
			It("should correctly reverse the string", func() {
				result := utils.StripAndReverse(example2)

				Expect(result).To(Equal("ppaym"))
			})
		})
	})
})
