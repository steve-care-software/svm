package lexers

import (
    "errors"
)

type instructionBuilder struct {
    module string
    kind Kind
    variable Variable
    assignment Assignment
    action Action
    execution Execution
}

func createInstructionBuilder() InstructionBuilder {
    out:= instructionBuilder{
        module: "",
        kind: nil,
        variable: nil,
        assignment: nil,
        action: nil,
        execution: nil,
    }

    return &out
}

// Create initializes the builder
func (app *instructionBuilder) Create() InstructionBuilder {
    return createInstructionBuilder()
}

// WithModule adds a module to the builder
func (app *instructionBuilder) WithModule(module string) InstructionBuilder {
    app.module = module
    return app
}

// WithKind adds a kind to the builder
func (app *instructionBuilder) WithKind(kind Kind) InstructionBuilder {
    app.kind = kind
    return app
}

// WithVariable adds a variable to the builder
func (app *instructionBuilder) WithVariable(variable Variable) InstructionBuilder {
    app.variable = variable
    return app
}

// WithAssignment adds an assignment to the builder
func (app *instructionBuilder) WithAssignment(assignment Assignment) InstructionBuilder {
    app.assignment = assignment
    return app
}

// WithAction adds an action to the builder
func (app *instructionBuilder) WithAction(action Action) InstructionBuilder {
    app.action = action
    return app
}

// WithExecution adds an execution to the builder
func (app *instructionBuilder) WithExecution(execution Execution) InstructionBuilder {
    app.execution = execution
    return app
}

// Now builds a new Instruction instance
func (app *instructionBuilder) Now() (Instruction, error) {
    if app.module != "" {
        return createInstructionWithModule(app.module), nil
    }

    if app.kind != nil {
        return createInstructionWithKind(app.kind), nil
    }

    if app.variable != nil {
        return createInstructionWithVariable(app.variable), nil
    }

    if app.assignment != nil {
        return createInstructionWithAssignment(app.assignment), nil
    }

    if app.action != nil {
        return createInstructionWithAction(app.action), nil
    }

    if app.execution != nil {
        return createInstructionWithExecution(app.execution), nil
    }

    return nil, errors.New("the Instruction is invalid")
}
