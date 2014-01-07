package user_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "user"
)

var _ = Describe("User", func() {
    var u *user.User

    BeforeEach(func() {
        u = user.New()
    })

    It("stores the Username and Password", func() {
        u.Login("mister_tee", "sekret")
        Expect(u.Username).To(Equal("mister_tee"))
        Expect(u.Password).To(Equal("sekret"))
    })
})
