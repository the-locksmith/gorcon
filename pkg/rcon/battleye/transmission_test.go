package battleye_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	be "github.com/playnet-public/gorcon/pkg/rcon/battleye"
)

var _ = Describe("Transmission", func() {
	var (
		t *be.Transmission
	)

	BeforeEach(func() {
		t = be.NewTransmission("test")
	})

	Describe("Key", func() {
		It("should return zero", func() {
			Expect(t.Key()).To(BeEquivalentTo(0))
		})
	})

	Describe("Request", func() {
		It("should return test", func() {
			Expect(t.Request()).To(BeEquivalentTo("test"))
		})
	})

	Describe("Done", func() {
		It("should return false", func() {
			Expect(t.Done()).To(BeEquivalentTo(false))
		})
	})

	Describe("Response", func() {
		It("should return empty string", func() {
			Expect(t.Response()).To(BeEquivalentTo(""))
		})
	})
})
