package parsers

import (
    "github.com/steve-care-software/svm/domain/lexers"
)

type variable struct {
    kind lexers.Kind
    name string
    pContent *string
}

func createVariable(
    kind lexers.Kind,
    name string,
) Variable {
    return createVariableInternally(kind, name, nil)
}

func createVariableWithContent(
    kind lexers.Kind,
    name string,
    pContent *string,
) Variable {
    return createVariableInternally(kind, name, pContent)
}

func createVariableInternally(
    kind lexers.Kind,
    name string,
    pContent *string,
) Variable {
    out := variable{
        kind: kind,
        name: name,
        pContent: pContent,
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

// HasContent returns true if there is content, false otherwise
func (obj *variable) HasContent() bool {
    return obj.pContent != nil
}

// Content returns the content, if any
func (obj *variable) Content() *string {
    return obj.pContent
}
