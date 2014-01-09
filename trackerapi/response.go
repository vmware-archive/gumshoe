package trackerapi

type MeResponseStructure struct {
    APIToken string `json:"api_token"`
    Username string `json:"username"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Initials string `json:"initials"`
    Timezone struct {
        Kind      string `json:"kind"`
        Offset    string `json:"offset"`
        OlsonName string `json:"olson_name"`
    }   `json:"time_zone"`
}
