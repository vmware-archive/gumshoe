package trackerapi

import "fmt"

type Authenticator interface {
    Authenticate(*User) (string, error)
}

type User struct {
    Username      string
    Password      string
    Email         string
    Name          string
    Timezone      string
    Initials      string
    APIToken      string
    authenticator Authenticator
}

func (u *User) Login(name, pass string) {
    u.Username = name
    u.Password = pass
}

func (u *User) IsAuthenticated() bool {
    return u.APIToken != ""
}

func (u *User) HasCredentials() bool {
    return u.Username != "" && u.Password != ""
}

func (u *User) SetAuthenticator(a Authenticator) {
    u.authenticator = a
}

func (u *User) Authenticate() error {
    token, err := u.authenticator.Authenticate(u)
    u.APIToken = token
    return err
}

func (u *User) String() string {
    return fmt.Sprintf("Username:  %s\nName:      %s\nEmail:     %s\nAPI Token: %s\nTimezone:  %s\nInitials:  %s", u.Username, u.Name, u.Email, u.APIToken, u.Timezone, u.Initials)
}
