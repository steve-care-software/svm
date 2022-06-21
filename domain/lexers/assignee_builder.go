package lexers

import (
	"errors"
)

type assigneeBuilder struct {
	name        string
	declaration Variable
}

func createAssigneeBuilder() AssigneeBuilder {
	out := assigneeBuilder{
		name:        "",
		declaration: nil,
	}

	return &out
}

// Create initializes the builder
func (app *assigneeBuilder) Create() AssigneeBuilder {
	return createAssigneeBuilder()
}

// WithName adds a name to the builder
func (app *assigneeBuilder) WithName(name string) AssigneeBuilder {
	app.name = name
	return app
}

// WithDeclaration adds a declaration to the builder
func (app *assigneeBuilder) WithDeclaration(declaration Variable) AssigneeBuilder {
	app.declaration = declaration
	return app
}

// Now builds a new Assignee instance
func (app *assigneeBuilder) Now() (Assignee, error) {
	if app.name != "" {
		return createAssigneeWithName(app.name), nil
	}

	if app.declaration != nil {
		return createAssigneeWithDeclaration(app.declaration), nil
	}

	return nil, errors.New("the Assignee is invalid")
}
