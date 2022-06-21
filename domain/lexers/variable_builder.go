package lexers

import (
	"errors"
)

type variableBuilder struct {
	module string
	kind   string
	name   string
}

func createVariableBuilder() VariableBuilder {
	out := variableBuilder{
		module: "",
		kind:   "",
		name:   "",
	}

	return &out
}

// Create initializes the builder
func (app *variableBuilder) Create() VariableBuilder {
	return createVariableBuilder()
}

// WithModule adds a module to the builder
func (app *variableBuilder) WithModule(module string) VariableBuilder {
	app.module = module
	return app
}

// WithKind adds a kind to the builder
func (app *variableBuilder) WithKind(kind string) VariableBuilder {
	app.kind = kind
	return app
}

// WithName adds a name to the builder
func (app *variableBuilder) WithName(name string) VariableBuilder {
	app.name = name
	return app
}

// Now builds a new Variable instance
func (app *variableBuilder) Now() (Variable, error) {
	if app.module == "" {
		return nil, errors.New("the module is mandatory in order to build a Variable instance")
	}

	if app.kind == "" {
		return nil, errors.New("the kind is mandatory in order to build a Variable instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Variable instance")
	}

	return createVariable(app.module, app.kind, app.name), nil
}
