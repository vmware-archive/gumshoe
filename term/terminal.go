package term

import (
    "bufio"
    "os"
    "os/exec"
    "strings"
)

const (
    EnableEcho  = iota
    DisableEcho = iota
)

type Terminal struct {
    InputFile   *os.File
    inputBuffer *bufio.Reader
    OutputFile  *os.File
}

func New() *Terminal {
    return &Terminal{
        InputFile:  os.Stdin,
        OutputFile: os.Stdout,
    }
}

func (t *Terminal) SetInput(input *os.File) {
    t.InputFile = input
}

func (t *Terminal) SetOutput(output *os.File) {
    t.OutputFile = output
}

func (t *Terminal) Prompt(prompt string, echo int) string {
    if echo == DisableEcho {
        t.DisableEcho()
    }
    t.OutputFile.WriteString(prompt)
    input := t.ReadLine()
    if echo == DisableEcho {
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
        t.OutputFile.WriteString(err.Error())
    }
    return strings.TrimSpace(string(line))
}

func (t *Terminal) run(command *exec.Cmd) {
    command.Stdin = t.InputFile
    command.Stdout = t.OutputFile
    command.Run()
}

func (t *Terminal) buffer() *bufio.Reader {
    if t.inputBuffer == nil {
        t.inputBuffer = bufio.NewReader(t.InputFile)
    }
    return t.inputBuffer
}
