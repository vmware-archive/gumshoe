package term_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/term"

    "os"
)

var _ = Describe("Terminal", func() {
    Describe("New", func() {
        BeforeEach(func() {
        })

        It("returns an instance of Terminal with a defined InputFile of os.Stdin", func() {
            terminal := term.New()
            Expect(terminal).To(BeAssignableToTypeOf(&term.Terminal{}))
            Expect(terminal.InputFile).To(Equal(os.Stdin))
        })
    })
})
