package parsers

import (
    "errors"
)

type programBuilder struct {
    executions Executions
    parameters Variables
}

func createProgramBuilder() ProgramBuilder {
    out := programBuilder{
        executions: nil,
        parameters: nil,
    }

    return &out
}

// Create initializes the builder
func (app *programBuilder) Create() ProgramBuilder {
    return createProgramBuilder()
}

// WithExecutions add executions to the builder
func (app *programBuilder) WithExecutions(executions Executions) ProgramBuilder {
    app.executions = executions
    return app
}

// WithParameters add parameters to the builder
func (app *programBuilder) WithParameters(parameters Variables) ProgramBuilder {
    app.parameters = parameters
    return app
}

// Now builds a new Program instance
func (app *programBuilder) Now() (Program, error) {
    if app.executions == nil {
        return nil, errors.New("the executions is mandatory in order to build a Program instance")
    }

    if app.parameters != nil {
        return createProgramWithParameters(app.executions, app.parameters), nil
    }

    return createProgram(app.executions), nil
}
