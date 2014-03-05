package matchers_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "testing"
)

type myStringer struct {
    a string
}

func (s *myStringer) String() string {
    return s.a
}

type StringAlias string

type myCustomType struct {
    s   string
    n   int
    f   float32
    arr []string
}

func Test(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Gomega")
}
