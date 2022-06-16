package lexers

import (
    "errors"
)

type kindBuilder struct {
    pFlag *uint8
    module string
    name string
}

func createKindBuilder() KindBuilder {
    out := kindBuilder{
        pFlag: nil,
        module: "",
        name: "",
    }

    return &out
}

// Create initializes the builder
func (app *kindBuilder) Create() KindBuilder {
    return createKindBuilder()
}

// WithFlag adds a flag to the builder
func (app *kindBuilder) WithFlag(flag uint8) KindBuilder {
    app.pFlag = &flag
    return app
}

// WithModule adds a module to the builder
func (app *kindBuilder) WithModule(module string) KindBuilder {
    app.module = module
    return app
}

// WithName adds a name to the builder
func (app *kindBuilder) WithName(name string) KindBuilder {
    app.name = name
    return app
}

// Now builds a new Kind instance
func (app *kindBuilder) Now() (Kind, error) {
    if app.pFlag == nil {
        return nil, errors.New("the flag is mandatory in order to build a Kind instance")
    }

    if app.module == "" {
        return nil, errors.New("the module is mandatory in order to build a Kind instance")
    }

    if app.name == "" {
        return nil, errors.New("the name is mandatory in order to build a Kind instance")
    }

    return createKind(*app.pFlag, app.module, app.name), nil
}
