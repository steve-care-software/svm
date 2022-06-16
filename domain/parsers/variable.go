package parsers

import (
    "github.com/steve-care-software/svm/domain/lexers"
)

type variable struct {
    kind lexers.Kind
    name string
    value interface{}
}

func createVariable(
    kind lexers.Kind,
    name string,
    value interface{},
) Variable {
    out := variable{
        kind: kind,
        name: name,
        value: value,
    }

    return &out
}

// Kind returns the kind
func (obj *variable) Kind() lexers.Kind {
    return obj.kind
}

// Name returns the name
func (obj *variable) Name() string {
    return obj.name
}

// Value returns the value
func (obj *variable) Value() interface{} {
    return obj.value
}
