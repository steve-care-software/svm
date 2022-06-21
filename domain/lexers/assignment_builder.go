package lexers

import (
	"errors"
)

type assignmentBuilder struct {
	content  string
	assignee Assignee
}

func createAssignmentBuilder() AssignmentBuilder {
	out := assignmentBuilder{
		content:  "",
		assignee: nil,
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

// WithAssignee adds an assignee to the builder
func (app *assignmentBuilder) WithAssignee(assignee Assignee) AssignmentBuilder {
	app.assignee = assignee
	return app
}

// Now builds a new Assignment instance
func (app *assignmentBuilder) Now() (Assignment, error) {
	if app.content == "" {
		return nil, errors.New("the content is mandatory in order to build an Assignment instance")
	}

	if app.assignee == nil {
		return nil, errors.New("the assignee is mandatory in order to build an Assignment instance")
	}

	return createAssignment(app.content, app.assignee), nil
}
