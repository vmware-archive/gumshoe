package trackerapi

import (
    "encoding/json"
    "fmt"

    "github.com/pivotal/gumshoe/trackerapi/domain"
    "github.com/pivotal/gumshoe/trackerapi/presenters"
    "github.com/pivotal/gumshoe/trackerapi/responses"
)

type Client struct {
    Resolver *Resolver
    user     *domain.User
    store    *Store
}

func NewClient() (*Client, error) {
    store := NewStore()
    token, err := store.Get("APIToken")
    if err != nil {
        return nil, err
    }

    user := &domain.User{
        APIToken: token,
    }
    user.SetAuthenticator(NewAPIAuthenticator())
    c := Client{
        Resolver: NewDefaultResolver(),
        store:    store,
        user:     user,
    }

    return &c, nil
}

func (c *Client) SetResolver(resolver *Resolver) {
    c.Resolver = resolver
}

func (c *Client) Me() fmt.Stringer {
    response := responses.Me{}
    responseBody := c.executeRequest(c.Resolver.MeRequestURL())
    err := json.Unmarshal(responseBody, &response.Structure)
    handleError(err)
    return presenters.User{response.User()}
}

func (c *Client) Projects() fmt.Stringer {
    response := responses.Projects{}
    responseBody := c.executeRequest(c.Resolver.ProjectsRequestURL())
    err := json.Unmarshal(responseBody, &response.Structure)
    handleError(err)
    return presenters.Projects{response.Projects()}
}

func (c *Client) Activity(projectID int) fmt.Stringer {
    response := responses.Activities{}
    responseBody := c.executeRequest(c.Resolver.ActivityRequestURL(projectID))
    err := json.Unmarshal(responseBody, &response.Structure)
    handleError(err)
    return presenters.Activities{response.Activities()}
}

func (c *Client) executeRequest(url string) []byte {
    strategy := &APITokenStrategy{
        APIToken: c.user.APIToken,
    }
    requester := NewRequester(url, strategy)
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
