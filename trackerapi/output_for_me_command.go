package trackerapi

import "fmt"

type OutputForMeCommand struct {
    user *MeResponseStructure
}

func (o *OutputForMeCommand) String() string {
    formatString := `
%s (%s)
  Email     : %s
  API Token : %s
  Timezone  : %s
  Initials  : %s
`
    return fmt.Sprintf(formatString, o.user.Name, o.user.Username, o.user.Email, o.user.APIToken, o.user.Timezone.OlsonName, o.user.Initials)
}
