package store_test

import (
    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/trackerapi/store"
)

var _ = Describe("MemoryStore", func() {

    var memoryStore *store.MemoryStore

    BeforeEach(func() {
        memoryStore = store.NewMemoryStore()
    })

    It("stores arbitrary key-value pairs", func() {
        key := "Ryan Gosling"
        value := "Hey girl, I heard you like reading... Maybe you could read my lips."
        memoryStore.Set(key, value)
        retVal, _ := memoryStore.Get(key)
        Expect(retVal).To(Equal(value))
    })

    It("does not persist the key-value pairs across store instances", func() {
        key := "Ryan Gosling"
        value := "Hey girl, I heard you like reading... Maybe you could read my lips."
        memoryStore.Set(key, value)
        retVal, _ := memoryStore.Get(key)
        Expect(retVal).To(Equal(value))

        memoryStore = store.NewMemoryStore()
        retVal, _ = memoryStore.Get(key)
        Expect(retVal).ToNot(Equal(value))
    })

    It("clears the cache and removes the .tracker file", func() {
        key := "Ryan Gosling"
        value := "Hey girl, I heard you like reading... Maybe you could read my lips."
        memoryStore.Set(key, value)
        retVal, _ := memoryStore.Get(key)
        Expect(retVal).To(Equal(value))
        memoryStore.Clear()

        retVal, _ = memoryStore.Get(key)
        Expect(retVal).To(Equal(""))
    })
})
