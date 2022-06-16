package parsers

import (
    "errors"
    "github.com/steve-care-software/svm/domain/lexers"
)

type variableBuilder struct {
    kind lexers.Kind
    name string
    value interface{}
}

func createVariableBuilder() VariableBuilder {
    out := variableBuilder{
        kind: nil,
        name: "",
        value: nil,
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

// WithValue adds a value to the builder
func (app *variableBuilder) WithValue(value interface{}) VariableBuilder {
    app.value = value
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

    return createVariable(app.kind, app.name, app.value), nil
}
