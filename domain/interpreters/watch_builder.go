package interpreters

import (
    "errors"
)

type watchBuilder struct {
    module string
    event WatchEvent
}

func createWatchBuilder() WatchBuilder {
    out := watchBuilder{
        module: "",
        event: nil,
    }

    return &out
}

// Create initializes the builder
func (app *watchBuilder) Create() WatchBuilder {
    return createWatchBuilder()
}

// WithModule adds a module to the builder
func (app *watchBuilder) WithModule(module string) WatchBuilder {
    app.module = module
    return app
}

// WithEvent adds an event to the builder
func (app *watchBuilder) WithEvent(event WatchEvent) WatchBuilder {
    app.event = event
    return app
}

// Now builds a new Watch instance
func (app *watchBuilder) Now() (Watch, error) {
    if app.module == "" {
        return nil, errors.New("the module is mandatory in order to build a Watch instance")
    }

    if app.event == nil {
        return nil, errors.New("the event is mandatory in order to build a Watch instance")
    }

    return createWatch(app.module, app.event), nil
}
