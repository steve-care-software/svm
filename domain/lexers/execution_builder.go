package lexers

import (
    "errors"
)

type executionBuilder struct {
    application string
    declaration Variable
}

func createExecutionBuilder() ExecutionBuilder {
    out := executionBuilder{
        application: "",
        declaration: nil,
    }

    return &out
}

// Create initializes the builder
func (app *executionBuilder) Create() ExecutionBuilder {
    return createExecutionBuilder()
}

// WithApplication adds an application to the builder
func (app *executionBuilder) WithApplication(application string) ExecutionBuilder {
    app.application = application
    return app
}

// WithDeclaration adds a declaration to the builder
func (app *executionBuilder) WithDeclaration(declaration Variable) ExecutionBuilder {
    app.declaration = declaration
    return app
}

// Now builds a new Execution instance
func (app *executionBuilder) Now() (Execution, error) {
    if app.application == "" {
        return nil, errors.New("the application is mandatory in order to build an Execution instance")
    }

    if app.declaration != nil {
        return createExecutionWithDeclaration(app.application, app.declaration), nil
    }

    return createExecution(app.application), nil
}
