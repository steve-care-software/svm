package interpreters

import (
    "fmt"
    "errors"
)

type watchesBuilder struct {
    list []Watch
}

func createWatchesBuilder() WatchesBuilder {
    out := watchesBuilder{
        list: nil,
    }

    return &out
}

// Create initializes the builder
func (app *watchesBuilder) Create() WatchesBuilder {
    return createWatchesBuilder()
}

// WithList adds a list to the builder
func (app *watchesBuilder) WithList(list []Watch) WatchesBuilder {
    app.list = list
    return app
}

// Now builds a new Watches instance
func (app *watchesBuilder) Now() (Watches, error) {
    if app.list != nil && len(app.list) <= 0 {
        app.list = nil
    }

    if app.list == nil {
        return nil, errors.New("there must be at least 1 Watch in order to build a Watches instance")
    }

    mp := map[string]Watch{}
    for _, oneWatch := range(app.list) {
        keyname := oneWatch.Module()
        mp[keyname] = oneWatch
    }

    if len(mp) != len(app.list) {
        diff := len(app.list) - len(mp)
        str := fmt.Sprintf("%d Watch instances were duplicates", diff)
        return nil, errors.New(str)
    }

    return createWatches(app.list, mp), nil
}
