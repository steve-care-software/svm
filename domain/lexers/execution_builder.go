package lexers

import (
	"errors"
)

type executionBuilder struct {
	application string
	assignee    Assignee
}

func createExecutionBuilder() ExecutionBuilder {
	out := executionBuilder{
		application: "",
		assignee:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *executionBuilder) Create() ExecutionBuilder {
	return createExecutionBuilder()
}

// WithApplication adds an application to the builder
func (app *executionBuilder) WithApplication(application string) ExecutionBuilder {
	app.application = application
	return app
}

// WithAssignee adds a assignee to the builder
func (app *executionBuilder) WithAssignee(assignee Assignee) ExecutionBuilder {
	app.assignee = assignee
	return app
}

// Now builds a new Execution instance
func (app *executionBuilder) Now() (Execution, error) {
	if app.application == "" {
		return nil, errors.New("the application is mandatory in order to build an Execution instance")
	}

	if app.assignee != nil {
		return createExecutionWithAssignee(app.application, app.assignee), nil
	}

	return createExecution(app.application), nil
}
