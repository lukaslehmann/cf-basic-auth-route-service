package broker_test

import (
	"github.com/benlaplanche/cf-basic-auth-route-service/broker"
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

	Describe(".Services", func() {
		It("returns a single service", func() {
			services := basicAuthBroker.Services()
			Expect(len(services)).To(Equal(1))
		})

		It("returns the correct service id", func() {
			Expect(basicAuthService.ID).To(Equal("6a97b5b8-1d1f-44bc-98ae-01d8d1047555"))
		})

		It("returns the correct service name", func() {
			Expect(basicAuthService.Name).To(Equal("p-basic-auth"))
		})

		It("returns the correct description", func() {
			Expect(basicAuthService.Description).To(Equal("Protect applications with basic authentication in the routing path"))
		})

		It("returns the correct tags", func() {
			Expect(basicAuthService.Tags).To(Equal([]string{"route-service", "basic-auth"}))
		})

		It("returns the service as bindable", func() {
			Expect(basicAuthService.Bindable).To(BeTrue())
		})

		It("returns the service plan as not updateable", func() {
			Expect(basicAuthService.PlanUpdatable).To(BeFalse())
		})

		It("returns a single plan", func() {
			plans := basicAuthService.Plans
			Expect(len(plans)).To(Equal(1))
		})

		It("returns the correct plan ID", func() {
			Expect(basicAuthServicePlan.ID).To(Equal("7becb74f-ce9d-4f52-87a2-50cc1b2b4b8f"))
		})

		It("returns the correct plan name", func() {
			Expect(basicAuthServicePlan.Name).To(Equal("reverse-name"))
		})
	})

})
