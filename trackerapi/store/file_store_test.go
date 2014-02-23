package store_test

import (
    "io/ioutil"
    "os"

    . "github.com/pivotal/gumshoe/repos/ginkgo"
    . "github.com/pivotal/gumshoe/repos/gomega"
    "github.com/pivotal/gumshoe/trackerapi/store"
)

var _ = Describe("FileStore", func() {

    var (
        fileStore *store.FileStore
        tempFile  *os.File
    )

    BeforeEach(func() {
        tempFile, _ = ioutil.TempFile("", ".tracker")
        fileStore = store.NewFileStore()
        fileStore.SetFilePath(tempFile.Name())
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
        fileStore.SetFilePath(tempFile.Name())
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
