package applications

import (
    "github.com/steve-care-software/svm/domain/lexers"
    "github.com/steve-care-software/svm/domain/parsers"
)

// Application represents the SVM application
type Application interface {
    Lex(script string) (lexers.Program, error)
    Parse(lexed lexers.Program) (parsers.Program, error)
    Interpret(params map[string]interface{}, program parsers.Program) (map[string]interface{}, error)
}
