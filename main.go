package main

import (
    "fmt"
    "os"

    "github.com/codegangsta/cli"
    "github.com/pivotal/gumshoe/trackerapi"
)

func main() {
    app := cli.NewApp()
    app.Name = "gumshoe"
    app.Usage = "talks to tracker"

    client := trackerapi.NewClient()
    app.Commands = []cli.Command{
        {
            Name:  "me",
            Usage: "prints out Tracker's representation of your account",
            Action: func(c *cli.Context) {
                output := client.Me()
                fmt.Println(output)
            },
        },
        {
            Name:  "projects",
            Usage: "prints out a list of Tracker projects for your account",
            Action: func(c *cli.Context) {
                output := client.Projects()
                fmt.Println(output)
            },
        },
    }

    app.Run(os.Args)
}
