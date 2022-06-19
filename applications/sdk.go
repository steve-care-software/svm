package applications

import (
    "github.com/steve-care-software/svm/domain/lexers"
    "github.com/steve-care-software/svm/domain/parsers"
    "github.com/steve-care-software/svm/domain/interpreters"
    "io"
)

// NewApplication creates a new application instance
func NewApplication(commentLogWriter io.Writer) Application {
    lexerAdapter := lexers.NewProgramAdapter()
    parserAdapter := parsers.NewProgramAdapter(commentLogWriter)
    variableBuilder := parsers.NewVariableBuilder()
    variablesBuilder := parsers.NewVariablesBuilder()
    return createApplication(lexerAdapter, parserAdapter, variableBuilder, variablesBuilder)
}

// Application represents the SVM application
type Application interface {
    Compile(script string) (parsers.Program, []byte, error)
    Execute(params map[string]string, modules interpreters.Modules, program parsers.Program) (parsers.Variables, error)
}
