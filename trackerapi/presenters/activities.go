package presenters

import (
    "fmt"

    "github.com/pivotal/gumshoe/trackerapi/domain"
)

type Activities struct {
    Activities []domain.Activity
}

func (a Activities) String() string {
    outputString := "Activity:\n"
    for _, activity := range a.Activities {
        activityPresenter := Activity{
            Activity: activity,
        }
        outputString += fmt.Sprintf("  %s", activityPresenter.String())
    }
    return outputString
}
