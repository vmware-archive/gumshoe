package trackerapi

import (
    "encoding/json"
    "fmt"
)

type Client struct {
    Resolver *Resolver
    user     *User
    store    *Store
}

func NewClient() (*Client, error) {
    store := NewStore()
    token, err := store.Get("APIToken")
    if err != nil {
        return nil, err
    }

    c := Client{
        Resolver: NewDefaultResolver(),
        store:    store,
        user: &User{
            APIToken:      token,
            authenticator: NewAPIAuthenticator(),
        },
    }

    return &c, nil
}

func (c *Client) SetResolver(resolver *Resolver) {
    c.Resolver = resolver
}

func (c *Client) Me() fmt.Stringer {
    structure := &MeResponseStructure{}
    c.executeRequest(structure, c.Resolver.MeRequestURL())
    return &OutputForMeCommand{
        user: structure,
    }
}

func (c *Client) Projects() fmt.Stringer {
    structure := &[]ProjectResponseStructure{}
    c.executeRequest(structure, c.Resolver.ProjectsRequestURL())
    return &OutputForProjectsCommand{
        projects: structure,
    }
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

    request := &Request{
        url:          url,
        authStrategy: strategy,
    }

    responseBody, err := request.Execute()
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
