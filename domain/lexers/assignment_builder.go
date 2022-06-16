package lexers

import (
    "errors"
)

type assignmentBuilder struct {
    name string
    declaration Variable
}

func createAssignmentBuilder() AssignmentBuilder {
    out := assignmentBuilder{
        name: "",
        declaration: nil,
    }

    return &out
}

// Create initializes the builder
func (app *assignmentBuilder) Create() AssignmentBuilder {
    return createAssignmentBuilder()
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
    if app.name != "" {
        return createAssignmentWithName(app.name), nil
    }

    if app.declaration != nil {
        return createAssignmentWithDeclaration(app.declaration), nil
    }

    return nil, errors.New("the Assignment is invalid")
}
