package term

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

type Terminal struct {
    InputFile   *os.File
    InputBuffer *bufio.Reader
}

func New() *Terminal {
    return &Terminal{
        InputFile: os.Stdin,
    }
}

func (t *Terminal) Prompt(prompt string, silent bool) string {
    if silent {
        t.DisableEcho()
    }
    fmt.Print(prompt)
    input := t.ReadLine()
    if silent {
        t.EnableEcho()
    }
    return input
}

func (t *Terminal) DisableEcho() {
    t.run(exec.Command("stty", "-echo"))
}

func (t *Terminal) EnableEcho() {
    t.run(exec.Command("stty", "echo"))
}

func (t *Terminal) ReadLine() string {
    buf := t.buffer()
    line, err := buf.ReadString('\n')
    if err != nil {
        fmt.Println(err)
    }
    return strings.TrimSpace(string(line))
}

func (t *Terminal) run(command *exec.Cmd) {
    command.Stdin = os.Stdin
    command.Stdout = os.Stdout
    command.Run()
}

func (t *Terminal) buffer() *bufio.Reader {
    if t.InputBuffer == nil {
        t.InputBuffer = bufio.NewReader(t.InputFile)
    }
    return t.InputBuffer
}
