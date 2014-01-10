package trackerapi

import "fmt"

const (
    TrackerDomain           = "https://www.pivotaltracker.com/services/v5"
    MeRequestPath           = "/me"
    AuthenticateRequestPath = "/me"
    ProjectsRequestPath     = "/projects"
    ActivityRequestPath     = "/projects/%d/activity?limit=5"
)

type Resolver struct {
    TrackerDomain string
}

func NewDefaultResolver() *Resolver {
    return &Resolver{
        TrackerDomain: TrackerDomain,
    }
}

func (r *Resolver) ActivityRequestURL(projectID int) string {
    return fmt.Sprintf(r.TrackerDomain+ActivityRequestPath, projectID)
}

func (r *Resolver) MeRequestURL() string {
    return r.TrackerDomain + MeRequestPath
}

func (r *Resolver) ProjectsRequestURL() string {
    return r.TrackerDomain + ProjectsRequestPath
}

func (r *Resolver) AuthenticateRequestURL() string {
    return r.TrackerDomain + AuthenticateRequestPath
}
