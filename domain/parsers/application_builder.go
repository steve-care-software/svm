package parsers

import (
    "errors"
)

type applicationBuilder struct {
    application Variable
    attachments Variables
}

func createApplicationBuilder() ApplicationBuilder {
    out := applicationBuilder{
        application: nil,
        attachments: nil,
    }

    return &out
}

// Create initializes the builder
func (app *applicationBuilder) Create() ApplicationBuilder {
    return createApplicationBuilder()
}

// WithApplication adds an application to the builder
func (app *applicationBuilder) WithApplication(application Variable) ApplicationBuilder {
    app.application = application
    return app
}

// WithAttachments add attachments to the builder
func (app *applicationBuilder) WithAttachments(attachments Variables) ApplicationBuilder {
    app.attachments = attachments
    return app
}

// Now builds a new Application instance
func (app *applicationBuilder) Now() (Application, error) {
    if app.application == nil {
        return nil, errors.New("the application is mandatory in order to build an Application instance")
    }

    if app.attachments != nil {
        return createApplicationWithAttachments(app.application, app.attachments), nil
    }

    return createApplication(app.application), nil
}
