package main

import (
    "github.com/codegangsta/cli"
    "os"
    "trackerapi"
)

func main() {
    app := cli.NewApp()
    app.Name = "gumshoe"
    app.Usage = "talks to tracker"

    app.Commands = []cli.Command{
        {
            Name:  "me",
            Usage: "prints out Tracker's representation of your account",
            Action: func(c *cli.Context) {
                trackerapi.Me()
            },
        },
    }

    app.Run(os.Args)
}
