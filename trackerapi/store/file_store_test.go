package store_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/trackerapi/store"
)

var _ = Describe("Store", func() {

    var fileStore *store.FileStore

    BeforeEach(func() {
        fileStore = store.NewFileStore()
    })

    It("stores arbitrary key-value pairs", func() {
        key := "Ryan Gosling"
        value := "Hey girl, I heard you like reading... Maybe you could read my lips."
        fileStore.Set(key, value)
        retVal, _ := fileStore.Get(key)
        Expect(retVal).To(Equal(value))
    })

    It("persists the key-value pairs across store instances", func() {
        key := "Ryan Gosling"
        value := "Hey girl, I heard you like reading... Maybe you could read my lips."
        fileStore.Set(key, value)
        retVal, _ := fileStore.Get(key)
        Expect(retVal).To(Equal(value))

        fileStore = store.NewFileStore()
        retVal, _ = fileStore.Get(key)
        Expect(retVal).To(Equal(value))
    })

    It("clears the cache and removes the .tracker file", func() {
        key := "Ryan Gosling"
        value := "Hey girl, I heard you like reading... Maybe you could read my lips."
        fileStore.Set(key, value)
        retVal, _ := fileStore.Get(key)
        Expect(retVal).To(Equal(value))
        fileStore.Clear()

        retVal, _ = fileStore.Get(key)
        Expect(retVal).To(Equal(""))
    })
})
