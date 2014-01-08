package trackerapi

import "fmt"

func NewUser() *User {
    return &User{}
}

type User struct {
    Username string
    Password string
    Email    string
    Name     string
    Timezone string
    Initials string
    APIToken string
}

func (u *User) Login(name, pass string) {
    u.Username = name
    u.Password = pass
}

func (u *User) IsAuthenticated() bool {
    return u.APIToken != ""
}

func (u *User) String() string {
    return fmt.Sprintf("Username:  %s\nName:      %s\nEmail:     %s\nAPI Token: %s\nTimezone:  %s\nInitials:  %s", u.Username, u.Name, u.Email, u.APIToken, u.Timezone, u.Initials)
}
