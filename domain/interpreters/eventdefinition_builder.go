package interpreters

import (
    "errors"
)

type eventDefinitionBuilder struct {
    assign AssignFn
    execute ExecuteFn
}

func createEventDefinitionBuilder() EventDefinitionBuilder {
    out := eventDefinitionBuilder{
        assign: nil,
        execute: nil,
    }

    return &out
}

// Create initializes the builder
func (app *eventDefinitionBuilder) Create() EventDefinitionBuilder {
    return createEventDefinitionBuilder()
}

// WithAssign adds an assign func to the builder
func (app *eventDefinitionBuilder) WithExecute(execute ExecuteFn) EventDefinitionBuilder {
    app.execute = execute
    return app
}

// WithAssign adds an assign func to the builder
func (app *eventDefinitionBuilder) WithAssign(assign AssignFn) EventDefinitionBuilder {
    app.assign = assign
    return app
}

// Now builds a new EventDefinition instance
func (app *eventDefinitionBuilder) Now() (EventDefinition, error) {
    if app.execute == nil {
        return nil, errors.New("the execute func is mandatory in order to build an EventDefinition instance")
    }

    if app.assign != nil {
        return createEventDefinitionWithAssign(app.execute, app.assign), nil
    }

    return createEventDefinition(app.execute), nil
}
