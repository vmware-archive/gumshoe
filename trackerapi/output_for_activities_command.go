package trackerapi

import "fmt"

type OutputForActivitiesCommand struct {
    activities *[]ActivityResponseStructure
}

type ActivityOutput struct {
    activity *ActivityResponseStructure
}

func (o *OutputForActivitiesCommand) String() string {
    activities := (*o.activities)
    outputString := "Activity:\n"
    for _, activity := range activities {
        activityOutput := &ActivityOutput{
            activity: &activity,
        }
        outputString += activityOutput.String()
    }
    return outputString
}

func (o *ActivityOutput) String() string {
    return fmt.Sprintf("  %s\n", o.activity.Message)
}
