package domain_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"

    "testing"
)

func TestDomain(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Domain Suite")
}
