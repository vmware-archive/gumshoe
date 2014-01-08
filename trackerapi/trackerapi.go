package trackerapi

import (
    "io"
    "log"
    "os"
    u "os/user"
)

const (
    DefaultAPIURL = "https://www.pivotaltracker.com/services/v5/me"
)

func NewLogger(wr io.Writer) *log.Logger {
    return log.New(wr, "", 0)
}

func NewFileLogger(path string) *log.Logger {
    file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }
    return NewLogger(file)
}

func NewStdoutLogger() *log.Logger {
    return NewLogger(os.Stdout)
}

func homeDir() string {
    usr, _ := u.Current()
    return usr.HomeDir
}
