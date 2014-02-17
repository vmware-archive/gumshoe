package presenters

import (
    "fmt"

    "github.com/pivotal/gumshoe/trackerapi/domain"
)

type Activity struct {
    Activity domain.Activity
}

func (a Activity) String() string {
    return fmt.Sprintf("%s\n", a.Activity.Message)
}
