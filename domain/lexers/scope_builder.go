package lexers

import (
	"errors"
)

type scopeBuilder struct {
	program string
	module  string
}

func createScopeBuilder() ScopeBuilder {
	out := scopeBuilder{
		program: "",
		module:  "",
	}

	return &out
}

// Create initializes the builder
func (app *scopeBuilder) Create() ScopeBuilder {
	return createScopeBuilder()
}

// WithProgram adds a program to the builder
func (app *scopeBuilder) WithProgram(program string) ScopeBuilder {
	app.program = program
	return app
}

// WithModule adds a module to the builder
func (app *scopeBuilder) WithModule(module string) ScopeBuilder {
	app.module = module
	return app
}

// Now builds a new Scope instance
func (app *scopeBuilder) Now() (Scope, error) {
	if app.program == "" {
		return nil, errors.New("the program name is mandatory in order to build a Scope instance")
	}

	if app.module == "" {
		return nil, errors.New("the module name is mandatory in order to build a Scope instance")
	}

	return createScope(app.program, app.module), nil
}
