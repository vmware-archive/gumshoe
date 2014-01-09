package trackerapi

import "errors"

type APIAuthenticator struct {
    URL  string
    user *User
}

func NewAPIAuthenticator() *APIAuthenticator {
    return &APIAuthenticator{
        URL: AuthenticateURL,
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
        url:            a.URL,
        authStrategy:   strategy,
        responseStruct: structure,
    }

    err := request.Execute()
    handleError(err)
    return structure.APIToken, nil
}
