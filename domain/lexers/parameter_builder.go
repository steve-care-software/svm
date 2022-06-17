package lexers

import (
    "errors"
)

type parameterBuilder struct {
    declaration Variable
    isInput bool
}

func createParameterBuilder() ParameterBuilder {
    out := parameterBuilder{
        declaration: nil,
        isInput: false,
    }

    return &out
}

// Create initializes the builder
func (app *parameterBuilder) Create() ParameterBuilder {
    return createParameterBuilder()
}

// WithDeclaration adds a declaration to the builder
func (app *parameterBuilder) WithDeclaration(declaration Variable) ParameterBuilder {
    app.declaration = declaration
    return app
}

// IsInput flags the builder as an input
func (app *parameterBuilder) IsInput() ParameterBuilder {
    app.isInput = true
    return app
}

// Now builds a new Parameter instance
func (app *parameterBuilder) Now() (Parameter, error) {
    if app.declaration == nil {
        return nil, errors.New("the variable declaration is mandatory in order to build a Parameter instance")
    }

    return createParameter(app.declaration, app.isInput), nil
}
