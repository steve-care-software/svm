package applications

import (
    "github.com/steve-care-software/svm/domain/parsers"
    "github.com/steve-care-software/svm/domain/interpreters"
)

// Application represents the SVM application
type Application interface {
    Compile(script string) (parsers.Program, error)
    Execute(params map[string]interface{}, modules interpreters.Modules, program parsers.Program) (map[string]interface{}, error)
}
