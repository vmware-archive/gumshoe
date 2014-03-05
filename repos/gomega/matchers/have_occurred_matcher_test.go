package matchers_test

import (
    "errors"
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    . "github.com/pivotal/gumshoe/repos/gomega/matchers"
)

var _ = Describe("HaveOccurred", func() {
    It("should succeed if matching an error", func() {
        Ω(errors.New("Foo")).Should(HaveOccurred())
    })

    It("should not succeed with nil", func() {
        Ω(nil).ShouldNot(HaveOccurred())
    })

    It("should only support errors and nil", func() {
        success, _, err := (&HaveOccurredMatcher{}).Match("foo")
        Ω(success).Should(BeFalse())
        Ω(err).Should(HaveOccurred())

        success, _, err = (&HaveOccurredMatcher{}).Match("")
        Ω(success).Should(BeFalse())
        Ω(err).Should(HaveOccurred())
    })
})
