package main

import (
    "os"

    "github.com/codegangsta/cli"
    "github.com/pivotal/gumshoe/trackerapi"
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
                client := trackerapi.NewClient()
                client.Me()
            },
        },
    }

    app.Run(os.Args)
}
