package parsers

import (
    "errors"
)

type executionBuilder struct {
    application Application
    output Variable
}

func createExecutionBuilder() ExecutionBuilder {
    out := executionBuilder{
        application: nil,
        output: nil,
    }

    return &out
}

// Create initializes the builder
func (app *executionBuilder) Create() ExecutionBuilder {
    return createExecutionBuilder()
}

// WithApplication adds an application to the builder
func (app *executionBuilder) WithApplication(application Application) ExecutionBuilder {
    app.application = application
    return app
}

// WithOutput adds an output to the builder
func (app *executionBuilder) WithOutput(output Variable) ExecutionBuilder {
    app.output = output
    return app
}

// Now builds a new Execution instance
func (app *executionBuilder) Now() (Execution, error) {
    if app.application == nil {
        return nil, errors.New("the application is mandatory in order to build an Execution instance")
    }

    if app.output != nil {
        return createExecutionWithOutput(app.application, app.output), nil
    }

    return createExecution(app.application), nil
}
