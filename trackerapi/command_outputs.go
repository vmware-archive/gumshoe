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

type ActivitiesOutput struct {
    activities *[]ActivityResponseStructure
}

type ActivityOutput struct {
    activity *ActivityResponseStructure
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

func (o *ActivitiesOutput) String() string {
    activities := (*o.activities)
    outputString := ""
    for _, activity := range activities {
        activityOutput := &ActivityOutput{
            activity: &activity,
        }
        outputString += activityOutput.String()
    }
    return outputString
}

func (o *ActivityOutput) String() string {
    return o.activity.Message + "\n"
}
