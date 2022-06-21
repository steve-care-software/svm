package applications

import (
	"errors"

	"github.com/steve-care-software/svm/domain/interpreters"
	"github.com/steve-care-software/svm/domain/lexers"
	"github.com/steve-care-software/svm/domain/parsers"
)

type builder struct {
	lexerAdapter     lexers.ProgramAdapter
	parserAdapter    parsers.ProgramAdapter
	variableBuilder  parsers.VariableBuilder
	variablesBuilder parsers.VariablesBuilder
	modules          interpreters.Modules
}

func createBuilder(
	lexerAdapter lexers.ProgramAdapter,
	parserAdapter parsers.ProgramAdapter,
	variableBuilder parsers.VariableBuilder,
	variablesBuilder parsers.VariablesBuilder,
) Builder {
	out := builder{
		lexerAdapter:     lexerAdapter,
		parserAdapter:    parserAdapter,
		variableBuilder:  variableBuilder,
		variablesBuilder: variablesBuilder,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.lexerAdapter,
		app.parserAdapter,
		app.variableBuilder,
		app.variablesBuilder,
	)
}

// WithModules add modules to the builder
func (app *builder) WithModules(modules interpreters.Modules) Builder {
	return createBuilder(
		app.lexerAdapter,
		app.parserAdapter,
		app.variableBuilder,
		app.variablesBuilder,
	)
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.modules == nil {
		return nil, errors.New("the modules are mandatory in order to build an Application instance")
	}

	return createApplication(
		app.lexerAdapter,
		app.parserAdapter,
		app.variableBuilder,
		app.variablesBuilder,
		app.modules,
	), nil
}
