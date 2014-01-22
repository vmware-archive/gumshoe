package trackerapi

import (
    "os/user"
)

func homeDir() string {
    usr, _ := user.Current()
    return usr.HomeDir
}

func handleError(err error) {
    if err != nil {
        panic(err)
    }
}
