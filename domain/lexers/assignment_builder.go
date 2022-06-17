package lexers

import (
    "errors"
)

type assignmentBuilder struct {
    content string
    name string
    declaration Variable
}

func createAssignmentBuilder() AssignmentBuilder {
    out := assignmentBuilder{
        content: "",
        name: "",
        declaration: nil,
    }

    return &out
}

// Create initializes the builder
func (app *assignmentBuilder) Create() AssignmentBuilder {
    return createAssignmentBuilder()
}

// WithContent adds a content to the builder
func (app *assignmentBuilder) WithContent(content string) AssignmentBuilder {
    app.content = content
    return app
}

// WithName adds a name to the builder
func (app *assignmentBuilder) WithName(name string) AssignmentBuilder {
    app.name = name
    return app
}

// WithDeclaration adds a declaration to the builder
func (app *assignmentBuilder) WithDeclaration(declaration Variable) AssignmentBuilder {
    app.declaration = declaration
    return app
}

// Now builds a new Assignment instance
func (app *assignmentBuilder) Now() (Assignment, error) {
    if app.content == "" {
        return nil, errors.New("the content is mandatory in order to build an Assignment instance")
    }

    if app.name != "" {
        return createAssignmentWithName(app.content, app.name), nil
    }

    if app.declaration != nil {
        return createAssignmentWithDeclaration(app.content, app.declaration), nil
    }

    return nil, errors.New("the Assignment is invalid")
}
