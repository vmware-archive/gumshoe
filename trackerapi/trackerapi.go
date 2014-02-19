package trackerapi

func handleError(err error) {
    if err != nil {
        panic(err)
    }
}
