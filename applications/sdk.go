package applications

import (
    "github.com/steve-care-software/svm/domain/parsers"
    "github.com/steve-care-software/svm/domain/interpreters/modules"
)

// Application represents the SVM application
type Application interface {
    Compile(script string) (parsers.Program, error)
    Execute(params map[string]interface{}, modules modules.Modules, program parsers.Program) (map[string]interface{}, error)
}
