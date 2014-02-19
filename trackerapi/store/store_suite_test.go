package store_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"

    "testing"
)

func TestStore(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Store Suite")
}
