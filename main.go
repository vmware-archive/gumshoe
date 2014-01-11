package main

import (
    "fmt"
    "os"
    "strconv"

    "github.com/pivotal/gumshoe/cmdutil"
    "github.com/pivotal/gumshoe/repos/cli"
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
    term := &Terminal{}
    app.Commands = []cli.Command{
        {
            Name:  "me",
            Usage: "prints out Tracker's representation of your account",
            Action: func(c *cli.Context) {
                handleAuthentication(client, term)
                output := client.Me()
                fmt.Println(output)
            },
        },
        {
            Name:  "projects",
            Usage: "prints out a list of Tracker projects for your account",
            Action: func(c *cli.Context) {
                handleAuthentication(client, term)
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
                handleAuthentication(client, term)
                output := client.Activity(projectID)
                fmt.Println(output)
            },
        },
    }

    app.Run(os.Args)
}

func handleAuthentication(client *trackerapi.Client, term *Terminal) {
    if !client.IsAuthenticated() {
        username := term.Prompt("Username: ", false)
        password := term.Prompt("Password: ", true)
        client.Authenticate(username, password)
    }
}

type Terminal struct{}

func (t *Terminal) Prompt(prompt string, silent bool) string {
    if silent {
        cmdutil.Silence()
    }
    fmt.Print(prompt)
    input := cmdutil.ReadLine()
    if silent {
        cmdutil.Unsilence()
    }
    return input
}
