package trackerapi

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "os"

    "github.com/pivotal/gumshoe/cmdutil"
)

type Client struct {
    URL          string
    FileLocation string
    Logger       *log.Logger
    user         *User
}

func NewClient() *Client {
    c := Client{
        URL:          DefaultAPIURL,
        FileLocation: homeDir() + "/.tracker",
        Logger:       NewLogger(os.Stdout),
    }
    return &c
}

func (c *Client) SetLogger(logger *log.Logger) {
    c.Logger = logger
}

func (c *Client) Me() {
    c.user = c.getUser()
    body, err := c.makeRequest()
    if err != nil {
        panic(err)
    }
    c.parse(body)
    if err != nil {
        panic(err)
    }
    c.Logger.Println(c.user)
}

func (c *Client) makeRequest() ([]byte, error) {
    httpClient := &http.Client{}
    req, err := http.NewRequest("GET", c.URL, nil)
    if err != nil {
        return nil, err
    }
    req.SetBasicAuth(c.user.Username, c.user.Password)
    resp, err := httpClient.Do(req)
    if err != nil {
        return nil, err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    return body, nil
}

func (c *Client) parse(body []byte) error {
    var resp = response{}
    err := json.Unmarshal(body, &resp)
    if err != nil {
        return err
    }

    c.user.APIToken = resp.APIToken
    c.user.Username = resp.Username
    c.user.Name = resp.Name
    c.user.Email = resp.Email
    c.user.Initials = resp.Initials
    c.user.Timezone = resp.Timezone.OlsonName
    return nil
}

func (c *Client) setCredentials(user *User) {
    c.Logger.Print("Username: ")
    var username = cmdutil.ReadLine()
    cmdutil.Silence()
    c.Logger.Print("Password: ")

    var password = cmdutil.ReadLine()
    user.Login(username, password)
    cmdutil.Unsilence()
}

func (c *Client) getUser() *User {
    user := &User{
        authenticator: &APIAuthenticator{},
    }
    if !user.IsAuthenticated() {
        c.setCredentials(user)
        user.Authenticate()
    }
    return user
}
