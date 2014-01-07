package user

func New() *User {
    return new(User)
}

type User struct {
    Username string
    Password string
    APIToken string
}

func (u *User) Login(name, pass string) {
    u.Username = name
    u.Password = pass
}
