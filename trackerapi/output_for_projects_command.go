package trackerapi

import "fmt"

type OutputForProjectsCommand struct {
    projects *[]ProjectResponseStructure
}

type ProjectOutput struct {
    project *ProjectResponseStructure
}

func (o *OutputForProjectsCommand) String() string {
    projects := (*o.projects)
    outputString := ""
    for _, project := range projects {
        projectOutput := &ProjectOutput{
            project: &project,
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
