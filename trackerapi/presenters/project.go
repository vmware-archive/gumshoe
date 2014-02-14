package presenters

import (
    "fmt"

    "github.com/pivotal/gumshoe/trackerapi/domain"
)

type Project struct {
    Project domain.Project
}

func (p Project) String() string {
    formatString := `
%s (%d)
  Current Iteration : %d
`
    output := fmt.Sprintf(
        formatString,
        p.Project.Name,
        p.Project.ID,
        p.Project.CurrentIteration)

    if p.Project.Description != "" {
        formatDesc := `  Description       : %s
`
        output += fmt.Sprintf(formatDesc, p.Project.Description)
    }
    return output
}
