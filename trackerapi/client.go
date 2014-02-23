package trackerapi

import (
    "fmt"

    "github.com/pivotal/gumshoe/trackerapi/domain"
    "github.com/pivotal/gumshoe/trackerapi/presenters"
    "github.com/pivotal/gumshoe/trackerapi/request"
    "github.com/pivotal/gumshoe/trackerapi/responses"
    "github.com/pivotal/gumshoe/trackerapi/store"
)

type Client struct {
    Resolver *request.Resolver
    user     *domain.User
    store    store.Store
}

func NewClient(s store.Store) (*Client, error) {
    if s == nil {
        s = store.NewFileStore()
    }
    token, err := s.Get("APIToken")
    handleError(err)

    user := &domain.User{
        APIToken: token,
    }
    user.SetAuthenticator(NewAPIAuthenticator())
    c := Client{
        Resolver: request.NewDefaultResolver(),
        store:    s,
        user:     user,
    }

    return &c, nil
}

func (c *Client) SetResolver(resolver *request.Resolver) {
    c.Resolver = resolver
}

func (c *Client) Me() fmt.Stringer {
    response := responses.Me{}
    responseBody := c.executeRequest(c.Resolver.MeRequestURL())
    response.Parse(responseBody)
    return presenters.User{response.User()}
}

func (c *Client) Projects() fmt.Stringer {
    response := responses.Projects{}
    responseBody := c.executeRequest(c.Resolver.ProjectsRequestURL())
    response.Parse(responseBody)
    return presenters.Projects{response.Projects()}
}

func (c *Client) Activity(projectID int) fmt.Stringer {
    response := responses.Activities{}
    responseBody := c.executeRequest(c.Resolver.ActivityRequestURL(projectID))
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
