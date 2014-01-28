package trackerapi

import (
    "encoding/json"
    "errors"
)

type APIAuthenticator struct {
    Resolver *Resolver
    user     *User
}

func NewAPIAuthenticator() *APIAuthenticator {
    return &APIAuthenticator{
        Resolver: NewDefaultResolver(),
    }
}

func (a *APIAuthenticator) Authenticate(u *User) (string, error) {
    a.user = u
    if !a.user.HasCredentials() {
        return "", errors.New("Given trackerapi.User does not have Username and Password")
    }

    structure := &MeResponseStructure{}
    strategy := &BasicAuthStrategy{
        Username: a.user.Username,
        Password: a.user.Password,
    }

    request := &Request{
        url:          a.Resolver.AuthenticateRequestURL(),
        authStrategy: strategy,
    }

    responseBody, err := request.Execute()
    handleError(err)

    err = json.Unmarshal(responseBody, &structure)
    if err != nil {
        return "", err
    }
    return structure.APIToken, nil
}
