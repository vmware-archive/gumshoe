package trackerapi

import "errors"

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
        url:            a.Resolver.AuthenticateRequestURL(),
        authStrategy:   strategy,
        responseStruct: structure,
    }

    err := request.Execute()
    handleError(err)
    return structure.APIToken, nil
}
