package parsers

import (
    "errors"
    "github.com/steve-care-software/svm/domain/lexers"
)

type variableBuilder struct {
    kind lexers.Kind
    name string
    pContent *string
}

func createVariableBuilder() VariableBuilder {
    out := variableBuilder{
        kind: nil,
        name: "",
        pContent: nil,
    }

    return &out
}

// Create initializes the builder
func (app *variableBuilder) Create() VariableBuilder {
    return createVariableBuilder()
}

// WithKind adds a kind to the builder
func (app *variableBuilder) WithKind(kind lexers.Kind) VariableBuilder {
    app.kind = kind
    return app
}

// WithName adds a name to the builder
func (app *variableBuilder) WithName(name string) VariableBuilder {
    app.name = name
    return app
}

// WithContent adds a content to the builder
func (app *variableBuilder) WithContent(content string) VariableBuilder {
    app.pContent = &content
    return app
}

// Now builds a new Variable instance
func (app *variableBuilder) Now() (Variable, error) {
    if app.kind == nil {
        return nil, errors.New("the kind is mandatory in order to build a Variable instance")
    }

    if app.name == "" {
        return nil, errors.New("the name is mandatory in order to build a Variable instance")
    }

    if app.pContent != nil {
        return createVariableWithContent(app.kind, app.name, app.pContent), nil
    }

    return createVariable(app.kind, app.name), nil
}
