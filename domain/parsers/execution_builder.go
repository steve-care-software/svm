package parsers

import (
    "errors"
)

type executionBuilder struct {
    module string
    application Variable
    attachments Variables
}

func createExecutionBuilder() ExecutionBuilder {
    out := executionBuilder{
        module: "",
        application: nil,
        attachments: nil,
    }

    return &out
}

// Create initializes the builder
func (app *executionBuilder) Create() ExecutionBuilder {
    return createExecutionBuilder()
}

// WithModule adds a module to the builder
func (app *executionBuilder) WithModule(module string) ExecutionBuilder {
    app.module = module
    return app
}

// WithApplication adds an application to the builder
func (app *executionBuilder) WithApplication(application Variable) ExecutionBuilder {
    app.application = application
    return app
}

// WithAttachments add attachments to the builder
func (app *executionBuilder) WithAttachments(attachments Variables) ExecutionBuilder {
    app.attachments = attachments
    return app
}

// Now builds a new Execution instance
func (app *executionBuilder) Now() (Execution, error) {
    if app.module == "" {
        return nil, errors.New("the module is mandatory in order to build an Execution instance")
    }

    if app.application == nil {
        return nil, errors.New("the application is mandatory in order to build an Execution instance")
    }

    if app.attachments != nil {
        return createExecutionWithAttachments(app.module, app.application, app.attachments), nil
    }

    return createExecution(app.module, app.application), nil
}
