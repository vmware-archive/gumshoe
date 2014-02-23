package domain_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/trackerapi/domain"
)

type FakeAuthenticator struct {
    APIToken string
}

func (fa *FakeAuthenticator) Authenticate(user domain.User) (string, error) {
    return fa.APIToken, nil
}

var _ = Describe("User", func() {
    var u domain.User

    BeforeEach(func() {
        u = domain.User{}
        u.SetAuthenticator(&FakeAuthenticator{APIToken: "abcde12345"})
    })

    It("stores the Username and Password", func() {
        u.Login("mister_tee", "sekret")
        Expect(u.Username).To(Equal("mister_tee"))
        Expect(u.Password).To(Equal("sekret"))
    })

    Describe("IsAuthenticated", func() {
        It("returns the authenticated state of the user object", func() {
            u.APIToken = "abcde"
            Expect(u.IsAuthenticated()).To(BeTrue())

            u.APIToken = ""
            Expect(u.IsAuthenticated()).To(BeFalse())
        })
    })

    Describe("HasCredentials", func() {
        It("returns a boolean if username or password are empty", func() {
            Expect(u.HasCredentials()).To(BeFalse())
            u.Username = "RyanGosling"
            Expect(u.HasCredentials()).To(BeFalse())
            u.Password = "HeyGirl...."
            Expect(u.HasCredentials()).To(BeTrue())
        })
    })

    Describe("Authenticate", func() {
        It("retrieves an APIToken using the authenticator", func() {
            u.APIToken = ""
            u.Authenticate()
            Expect(u.APIToken).To(Equal("abcde12345"))
        })
    })
})
