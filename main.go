package main

import (
    "fmt"
    "os"
    "strconv"

    "github.com/pivotal/gumshoe/repos/cli"
    "github.com/pivotal/gumshoe/term"
    "github.com/pivotal/gumshoe/trackerapi"
)

func main() {
    app := cli.NewApp()
    app.Name = "gumshoe"
    app.Usage = "talks to tracker"

    client, err := trackerapi.NewClient()
    if err != nil {
        panic(err)
    }
    terminal := term.New()
    app.Commands = []cli.Command{
        {
            Name:  "me",
            Usage: "prints out Tracker's representation of your account",
            Action: func(c *cli.Context) {
                handleAuthentication(client, terminal)
                output := client.Me()
                fmt.Println(output)
            },
        },
        {
            Name:  "projects",
            Usage: "prints out a list of Tracker projects for your account",
            Action: func(c *cli.Context) {
                handleAuthentication(client, terminal)
                output := client.Projects()
                fmt.Println(output)
            },
        },
        {
            Name:  "activity",
            Usage: "lists last 5 activities for a given project",
            Action: func(c *cli.Context) {
                projectID, err := strconv.Atoi(c.Args()[0])
                if err != nil {
                    panic(err)
                }
                handleAuthentication(client, terminal)
                output := client.Activity(projectID)
                fmt.Println(output)
            },
        },
    }

    app.Run(os.Args)
}

func handleAuthentication(client *trackerapi.Client, terminal *term.Terminal) {
    if !client.IsAuthenticated() {
        username := terminal.Prompt("Username: ", false)
        password := terminal.Prompt("Password: ", true)
        client.Authenticate(username, password)
    }
}
