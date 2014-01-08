package trackerapi_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "github.com/pivotal/gumshoe/trackerapi"
)

var _ = Describe("User", func() {
    var u *trackerapi.User

    BeforeEach(func() {
        u = trackerapi.NewUser()
    })

    It("stores the Username and Password", func() {
        u.Login("mister_tee", "sekret")
        Expect(u.Username).To(Equal("mister_tee"))
        Expect(u.Password).To(Equal("sekret"))
    })

    Describe("isAuthenticated", func() {
        It("returns the authenticated state of the user object", func() {
            u.APIToken = "abcde"
            Expect(u.IsAuthenticated()).To(BeTrue())
        })
    })
})
