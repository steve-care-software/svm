package lexers

import (
    "errors"
)

type actionBuilder struct {
    scope Scope
    application string
    isAttach bool
}

func createActionBuilder() ActionBuilder {
    out := actionBuilder{
        scope: nil,
        application: "",
        isAttach: false,
    }

    return &out
}

// Create initializes the builder
func (app *actionBuilder) Create() ActionBuilder {
    return createActionBuilder()
}

// WithScope adds a scope to the builder
func (app *actionBuilder) WithScope(scope Scope) ActionBuilder {
    app.scope = scope
    return app
}

// WithApplication adds an application to the builder
func (app *actionBuilder) WithApplication(application string) ActionBuilder {
    app.application = application
    return app
}

// IsAttach flags the builder as attach
func (app *actionBuilder) IsAttach() ActionBuilder {
    app.isAttach = true
    return app
}

// Now builds a new Action instance
func (app *actionBuilder) Now() (Action, error) {
    if app.scope == nil {
        return nil, errors.New("the scope is mandatory in order to build an Action instance")
    }

    if app.application == "" {
        return nil, errors.New("the application is mandatory in order to build an Action instance")
    }

    return createAction(app.scope, app.application, app.isAttach), nil
}
