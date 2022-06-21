package applications

import (
	"io"

	"github.com/steve-care-software/svm/domain/interpreters"
	"github.com/steve-care-software/svm/domain/lexers"
	"github.com/steve-care-software/svm/domain/parsers"
)

// NewApplicationBuilder creates a new application builder
func NewApplicationBuilder(commentLogWriter io.Writer) Builder {
	lexerAdapter := lexers.NewProgramAdapter()
	parserAdapter := parsers.NewProgramAdapter(commentLogWriter)
	variableBuilder := parsers.NewVariableBuilder()
	variablesBuilder := parsers.NewVariablesBuilder()
	return createBuilder(
		lexerAdapter,
		parserAdapter,
		variableBuilder,
		variablesBuilder,
	)
}

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithModules(modules interpreters.Modules) Builder
	Now() (Application, error)
}

// Application represents the SVM application
type Application interface {
	Compile(script string) (parsers.Program, []byte, error)
	Execute(params map[string]string, program parsers.Program) (parsers.Variables, error)
}
