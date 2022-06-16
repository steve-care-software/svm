package lexers

import (
    "errors"
)

type programBuilder struct {
    instructions  []Instruction
    parameters []Variable
}

func createProgramBuilder() ProgramBuilder {
    out:= programBuilder{
        instructions: nil,
        parameters: nil,
    }

    return &out
}

// Create initializes the builder
func (app *programBuilder) Create() ProgramBuilder {
    return createProgramBuilder()
}

// WithInstructions add instructions to the builder
func (app *programBuilder) WithInstructions(instructions []Instruction) ProgramBuilder {
    app.instructions = instructions
    return app
}

// WithParameters add parameters to the builder
func (app *programBuilder) WithParameters(parameters []Variable) ProgramBuilder {
    app.parameters = parameters
    return app
}

// Now builds a new Program instance
func (app *programBuilder) Now() (Program, error) {
    if app.instructions != nil && len(app.instructions) <= 0 {
        app.instructions = nil
    }

    if app.instructions == nil {
        return nil, errors.New("there must be at least 1 Instruction in order to build a Program instance")
    }

    if app.parameters != nil && len(app.parameters) <= 0 {
        app.parameters = nil
    }

    if app.parameters != nil {
        return createProgramWithParameters(app.instructions, app.parameters), nil
    }

    return createProgram(app.instructions), nil
}
