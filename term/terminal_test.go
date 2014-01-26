package term_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/term"

    "io/ioutil"
    "os"
)

var _ = Describe("term", func() {
    Describe("New", func() {
        var terminal *term.Terminal
        BeforeEach(func() {
            terminal = term.New()
        })

        It("returns an instance of Terminal", func() {
            Expect(terminal).To(BeAssignableToTypeOf(&term.Terminal{}))
        })

        It("defaults the InputFile to os.Stdin", func() {
            Expect(terminal.InputFile).To(Equal(os.Stdin))
        })

        It("defaults the OutputFile to os.Stdout", func() {
            Expect(terminal.OutputFile).To(Equal(os.Stdout))
        })
    })

    Describe("Terminal#Prompt", func() {
        var (
            terminal      *term.Terminal
            stdin, stdout *os.File
        )
        BeforeEach(func() {
            terminal = term.New()
            stdin, _ = ioutil.TempFile(os.TempDir(), "stdin")
            stdout, _ = ioutil.TempFile(os.TempDir(), "stdout")
            terminal.SetInput(stdin)
            terminal.SetOutput(stdout)
        })

        It("prints the prompt, waits for and returns user input", func() {
            stdin.WriteString("Mister Tee")
            stdin.Seek(0, 0)
            output := terminal.Prompt("What is your name?: ", term.EnableEcho)
            Expect(output).To(ContainSubstring("Mister Tee"))
            stdoutContents, _ := ioutil.ReadFile(stdout.Name())
            Expect(string(stdoutContents)).To(ContainSubstring("What is your name?:"))
        })
    })
})
