package presenters

import (
    "fmt"
    "github.com/pivotal/gumshoe/trackerapi/domain"
)

type User struct {
    User domain.User
}

func (p User) String() string {
    formatString := `
%s (%s)
  Email     : %s
  API Token : %s
  Timezone  : %s
  Initials  : %s
`
    return fmt.Sprintf(
        formatString,
        p.User.Name,
        p.User.Username,
        p.User.Email,
        p.User.APIToken,
        p.User.Timezone,
        p.User.Initials)
}
