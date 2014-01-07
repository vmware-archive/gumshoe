package trackerapi

import (
    "cmdutil"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    u "os/user"
    "user"
)

var (
    URL          string     = "https://www.pivotaltracker.com/services/v5/me"
    FileLocation string     = homeDir() + "/.tracker"
    currentUser  *user.User = user.New()
    Stdout       *os.File   = os.Stdout
)

func Me() {
    setCredentials()
    parse(makeRequest())
    ioutil.WriteFile(FileLocation, []byte(currentUser.APIToken), 0644)
}

func makeRequest() []byte {
    client := &http.Client{}
    req, err := http.NewRequest("GET", URL, nil)
    req.SetBasicAuth(currentUser.Username, currentUser.Password)
    resp, err := client.Do(req)
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Print(err)
    }
    return body
}

func parse(body []byte) {
    var meResp = new(MeResponse)
    err := json.Unmarshal(body, &meResp)
    if err != nil {
        fmt.Println("error:", err)
    }

    currentUser.APIToken = meResp.APIToken
}

func setCredentials() {
    fmt.Fprint(Stdout, "Username: ")
    var username = cmdutil.ReadLine()
    cmdutil.Silence()
    fmt.Fprint(Stdout, "Password: ")

    var password = cmdutil.ReadLine()
    currentUser.Login(username, password)
    cmdutil.Unsilence()
}

func homeDir() string {
    usr, _ := u.Current()
    return usr.HomeDir
}

type MeResponse struct {
    APIToken string `json:"api_token"`
}
