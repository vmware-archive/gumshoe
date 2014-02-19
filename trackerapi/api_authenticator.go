package trackerapi

import (
    "errors"

    "github.com/pivotal/gumshoe/trackerapi/domain"
    "github.com/pivotal/gumshoe/trackerapi/responses"
)

type APIAuthenticator struct {
    Resolver *Resolver
    user     *domain.User
}

func NewAPIAuthenticator() *APIAuthenticator {
    return &APIAuthenticator{
        Resolver: NewDefaultResolver(),
    }
}

func (a *APIAuthenticator) Authenticate(u *domain.User) (string, error) {
    a.user = u
    if !a.user.HasCredentials() {
        return "", errors.New("Given domain.User does not have Username and Password")
    }

    response := responses.Me{}
    strategy := &BasicAuthStrategy{
        Username: a.user.Username,
        Password: a.user.Password,
    }
    requester := NewRequester(a.Resolver.AuthenticateRequestURL(), strategy)
    responseBody, err := requester.Execute()
    handleError(err)

    err = response.Parse(responseBody)
    if err != nil {
        return "", err
    }
    return response.User().APIToken, nil
}
