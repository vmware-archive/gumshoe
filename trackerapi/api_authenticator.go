package trackerapi

import (
    "encoding/json"
    "errors"

    "github.com/pivotal/gumshoe/trackerapi/domain"
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

    structure := &MeResponseStructure{}
    strategy := &BasicAuthStrategy{
        Username: a.user.Username,
        Password: a.user.Password,
    }
    requester := NewRequester(a.Resolver.AuthenticateRequestURL(), strategy)
    responseBody, err := requester.Execute()
    handleError(err)

    err = json.Unmarshal(responseBody, &structure)
    if err != nil {
        return "", err
    }
    return structure.APIToken, nil
}
