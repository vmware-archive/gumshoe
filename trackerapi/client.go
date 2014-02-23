package trackerapi

import (
    "fmt"

    "github.com/pivotal/gumshoe/trackerapi/domain"
    "github.com/pivotal/gumshoe/trackerapi/presenters"
    "github.com/pivotal/gumshoe/trackerapi/request"
    "github.com/pivotal/gumshoe/trackerapi/responses"
    "github.com/pivotal/gumshoe/trackerapi/store"
)

type Configuration struct {
    Store    store.Store
    Resolver request.Resolver
}

func NewConfiguration() Configuration {
    return Configuration{
        Store:    store.NewFileStore(),
        Resolver: request.NewDefaultResolver(),
    }
}

func (c Configuration) User() domain.User {
    token, err := c.Store.Get("APIToken")
    handleError(err)

    user := domain.User{
        APIToken: token,
    }
    user.SetAuthenticator(NewAPIAuthenticator())
    return user
}

type Client struct {
    resolver request.Resolver
    user     domain.User
    store    store.Store
}

func NewClient(c Configuration) (*Client, error) {
    client := Client{
        resolver: c.Resolver,
        store:    c.Store,
        user:     c.User(),
    }

    return &client, nil
}

func (c *Client) SetResolver(resolver request.Resolver) {
    c.resolver = resolver
}

func (c *Client) Me() fmt.Stringer {
    response := responses.Me{}
    responseBody := c.executeRequest(c.resolver.MeRequestURL())
    response.Parse(responseBody)
    return presenters.User{response.User()}
}

func (c *Client) Projects() fmt.Stringer {
    response := responses.Projects{}
    responseBody := c.executeRequest(c.resolver.ProjectsRequestURL())
    response.Parse(responseBody)
    return presenters.Projects{response.Projects()}
}

func (c *Client) Activity(projectID int) fmt.Stringer {
    response := responses.Activities{}
    responseBody := c.executeRequest(c.resolver.ActivityRequestURL(projectID))
    response.Parse(responseBody)
    return presenters.Activities{response.Activities()}
}

func (c *Client) executeRequest(url string) []byte {
    strategy := &request.APITokenStrategy{
        APIToken: c.user.APIToken,
    }
    requester := request.New(url, strategy)
    responseBody, err := requester.Execute()
    handleError(err)
    return responseBody
}

func (c *Client) Cleanup() {
    c.store.Clear()
}

func (c *Client) IsAuthenticated() bool {
    return c.user.IsAuthenticated()
}

func (c *Client) Authenticate(username, password string) {
    c.user.Login(username, password)
    c.user.Authenticate()
    c.store.Set("APIToken", c.user.APIToken)
}
