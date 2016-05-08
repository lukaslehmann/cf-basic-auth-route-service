package broker_test

import (
	. "github.com/benlaplanche/cf-basic-auth-route-service/broker"
	"github.com/pivotal-cf/brokerapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic Auth Service Broker", func() {
	var basicAuthBroker *broker.BasicAuthBroker
	var basicAuthService *brokerapi.Service
	var basicAuthServicePlan *brokerapi.ServicePlan

	BeforeEach(func() {
		basicAuthBroker = &broker.BasicAuthBroker{}
		basicAuthService = &basicAuthBroker.Services()[0]
		basicAuthServicePlan = &basicAuthService.Plans[0]
	})

	Desccribe(".Services", func() {
		It("returns a single service", func() {
			services = basicAuthBroker.Services()
			Expect(len(services)).To(Equal(1))
		})

		It("returns the correct service id", func() {
			Expect(basicAuthService.ID).To(Equal("6a97b5b8-1d1f-44bc-98ae-01d8d1047555"))
		})

	})

})
