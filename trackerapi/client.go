package trackerapi

import (
    "encoding/json"
    "fmt"

    "github.com/pivotal/gumshoe/trackerapi/domain"
    "github.com/pivotal/gumshoe/trackerapi/presenters"
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
    response := MeResponse{}
    c.executeRequest(&response.Structure, c.Resolver.MeRequestURL())
    return presenters.User{response.User()}
}

func (c *Client) Projects() fmt.Stringer {
    response := ProjectsResponse{}
    c.executeRequest(&response.Structure, c.Resolver.ProjectsRequestURL())
    return presenters.Projects{response.Projects()}
}

func (c *Client) Activity(projectID int) fmt.Stringer {
    structure := &[]ActivityResponseStructure{}
    c.executeRequest(structure, c.Resolver.ActivityRequestURL(projectID))
    return &OutputForActivitiesCommand{
        activities: structure,
    }
}

func (c *Client) executeRequest(structure interface{}, url string) {
    strategy := &APITokenStrategy{
        APIToken: c.user.APIToken,
    }
    requester := NewRequester(url, strategy)
    responseBody, err := requester.Execute()
    handleError(err)
    err = json.Unmarshal(responseBody, &structure)
    handleError(err)
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
