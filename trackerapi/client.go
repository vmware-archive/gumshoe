package trackerapi

import (
    "fmt"
    "log"
    "os"

    "github.com/pivotal/gumshoe/cmdutil"
)

type Client struct {
    Logger   *log.Logger
    Resolver *Resolver
    user     *User
    store    *Store
}

func NewClient() *Client {
    c := Client{
        Logger:   NewLogger(os.Stdout),
        Resolver: NewDefaultResolver(),
        store:    NewStore(),
    }
    return &c
}

func (c *Client) SetLogger(logger *log.Logger) {
    c.Logger = logger
}

func (c *Client) SetResolver(resolver *Resolver) {
    c.Resolver = resolver
}

func (c *Client) Me() fmt.Stringer {
    structure := &MeResponseStructure{}
    c.executeRequest(structure, c.Resolver.MeRequestURL())
    return &MeOutput{
        user: structure,
    }
}

func (c *Client) Projects() fmt.Stringer {
    structure := &[]ProjectResponseStructure{}
    c.executeRequest(structure, c.Resolver.ProjectsRequestURL())
    return &ProjectsOutput{
        projects: structure,
    }
}

func (c *Client) Activity(projectID int) fmt.Stringer {
    structure := &[]ActivityResponseStructure{}
    c.executeRequest(structure, c.Resolver.ActivityRequestURL(projectID))
    return &ActivitiesOutput{
        activities: structure,
    }
}

func (c *Client) executeRequest(structure interface{}, url string) {
    var err error
    c.user, err = c.setupUser()
    handleError(err)

    strategy := &APITokenStrategy{
        APIToken: c.user.APIToken,
    }

    request := &Request{
        url:            url,
        authStrategy:   strategy,
        responseStruct: structure,
    }

    err = request.Execute()
    handleError(err)
}

func (c *Client) setCredentials(user *User) {
    c.Logger.Print("Username: ")
    var username = cmdutil.ReadLine()
    cmdutil.Silence()
    c.Logger.Print("Password: ")

    var password = cmdutil.ReadLine()
    user.Login(username, password)
    cmdutil.Unsilence()
}

func (c *Client) setupUser() (*User, error) {
    token, err := c.store.Get("APIToken")
    if err != nil {
        return nil, err
    }
    user := &User{
        APIToken:      token,
        authenticator: NewAPIAuthenticator(),
    }
    if !user.IsAuthenticated() {
        c.setCredentials(user)
        user.Authenticate()
        c.store.Set("APIToken", user.APIToken)
    }
    return user, nil
}

func (c *Client) Cleanup() {
    c.store.Clear()
}
