package trackerapi

import "fmt"

type MeOutput struct {
    user *MeResponseStructure
}

func (o *MeOutput) String() string {
    formatString := `
Username  : %s
Name      : %s
Email     : %s
API Token : %s
Timezone  : %s
Initials  : %s
    `
    return fmt.Sprintf(formatString, o.user.Username, o.user.Name, o.user.Email, o.user.APIToken, o.user.Timezone.OlsonName, o.user.Initials)
}
