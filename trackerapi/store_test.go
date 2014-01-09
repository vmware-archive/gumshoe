package trackerapi_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/trackerapi"
)

var _ = Describe("Store", func() {

    var store *trackerapi.Store

    BeforeEach(func() {
        store = trackerapi.NewStore()
    })

    It("stores arbitrary key-value pairs", func() {
        key := "Ryan Gosling"
        value := "Hey girl, I heard you like reading... Maybe you could read my lips."
        store.Set(key, value)
        retVal, _ := store.Get(key)
        Expect(retVal).To(Equal(value))
    })

    It("persists the key-value pairs across store instances", func() {
        key := "Ryan Gosling"
        value := "Hey girl, I heard you like reading... Maybe you could read my lips."
        store.Set(key, value)
        retVal, _ := store.Get(key)
        Expect(retVal).To(Equal(value))

        store = trackerapi.NewStore()
        retVal, _ = store.Get(key)
        Expect(retVal).To(Equal(value))
    })

    It("clears the cache and removes the .tracker file", func() {
        key := "Ryan Gosling"
        value := "Hey girl, I heard you like reading... Maybe you could read my lips."
        store.Set(key, value)
        retVal, _ := store.Get(key)
        Expect(retVal).To(Equal(value))
        store.Clear()

        retVal, _ = store.Get(key)
        Expect(retVal).To(Equal(""))
    })
})
