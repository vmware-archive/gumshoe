package trackerapi

import "fmt"

type MeOutput struct {
    user *MeResponseStructure
}

type ProjectsOutput struct {
    projects *[]ProjectResponseStructure
}

type ProjectOutput struct {
    project *ProjectResponseStructure
}

func (o *MeOutput) String() string {
    formatString := `
%s (%s)
  Email     : %s
  API Token : %s
  Timezone  : %s
  Initials  : %s
`
    return fmt.Sprintf(formatString, o.user.Name, o.user.Username, o.user.Email, o.user.APIToken, o.user.Timezone.OlsonName, o.user.Initials)
}

func (o *ProjectsOutput) String() string {
    projects := (*o.projects)
    outputString := ""
    for i := 0; i < len(projects); i++ {
        projectOutput := &ProjectOutput{
            project: &projects[i],
        }
        outputString += projectOutput.String()
    }
    return outputString
}

func (o *ProjectOutput) String() string {
    formatString := `
%s (%d)
  Current Iteration : %d
`
    output := fmt.Sprintf(formatString, o.project.Name, o.project.ID, o.project.CurrentIteration)
    if o.project.Description != "" {
        formatDesc := `  Description       : %s
`
        output += fmt.Sprintf(formatDesc, o.project.Description)
    }
    return output
}
