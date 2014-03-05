package matchers_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    . "github.com/pivotal/gumshoe/repos/gomega/matchers"
)

var _ = Describe("BeFalse", func() {
    It("should handle true and false correctly", func() {
        Ω(true).ShouldNot(BeFalse())
        Ω(false).Should(BeFalse())
    })

    It("should only support booleans", func() {
        success, _, err := (&BeFalseMatcher{}).Match("foo")
        Ω(success).Should(BeFalse())
        Ω(err).Should(HaveOccured())
    })
})
