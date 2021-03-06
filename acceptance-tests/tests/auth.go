package tests

import (
	"os"
	"github.com/fabric8-analytics/cli-tools/acceptance-tests/helper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TestCRDAauth implements Test Cases to test auth crda command
func TestCRDAauth() {
	When("I run crda auth without snyk token", func() {
		It("should throw error", func() {
			session := helper.CmdShouldFailWithExit1(getCRDAcmd(), "auth", "--snyk-token")
			Expect(string(session)).To(ContainSubstring("flag needs an argument: --snyk-token"))
		})
	})
	When("I run crda auth with invalid snyk token", func() {
		It("it should throw error", func() {
			session := helper.CmdShouldFailWithExit1(getCRDAcmd(), "auth", "--snyk-token", "invalid-token")
			Expect(string(session)).To(ContainSubstring("Snyk API Token is invalid"))
		})

	})
	When("I run crda auth with valid snyk token", func() {
		validToken := os.Getenv("snyk_token")
		It("it should not throw error", func() {
			if validToken== ""{
				Skip("Running in PR Check")
			}
			session := helper.CmdShouldPassWithoutError(getCRDAcmd(), "auth", "--snyk-token", string(validToken))
			helper.PrintWithGinkgo(session)
		})

	})
}
