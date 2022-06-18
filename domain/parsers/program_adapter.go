package parsers

import (
    "github.com/steve-care-software/svm/domain/lexers"
)

type programAdapter struct {
    lexerAdapter lexers.ProgramAdapter
    computerFactory ComputerFactory
}

func createProgramAdapter(
    lexerAdapter lexers.ProgramAdapter,
    computerFactory ComputerFactory,
    ) ProgramAdapter {
    out := programAdapter {
        lexerAdapter: lexerAdapter,
        computerFactory: computerFactory,
    }

    return &out
}

// LexedToProgram converts a lexed program to a parsed program
func (app *programAdapter) LexedToProgram(lexed lexers.Program) (Program, error) {
    computer := app.computerFactory.Create()
    if lexed.HasParameters() {
        parameters := lexed.Parameters()
        err := app.parameters(computer, parameters)
        if err != nil {
            return nil, err
        }
    }

    instructions := lexed.Instructions()
    err := app.instructions(computer, instructions)
    if err != nil {
        return nil, err
    }

    return computer.Program()
}

func (app *programAdapter) parameters(computer Computer, parameters []lexers.Parameter) error {
    for _, oneParameter := range(parameters) {
        computer.Parameter(oneParameter)
    }

    return nil
}

func (app *programAdapter) instructions(computer Computer, instructions []lexers.Instruction) error {
    for _, oneInstruction := range(instructions) {
        err := app.instruction(computer, oneInstruction)
        if err != nil {
            return err
        }
    }

    return nil
}

func (app *programAdapter) instruction(computer Computer, instruction lexers.Instruction) error {
    if instruction.IsModule() {
        module := instruction.Module()
        return computer.Module(module)
    }

    if instruction.IsKind() {
        kind := instruction.Kind()
        return computer.Kind(kind)
    }

    if instruction.IsVariable() {
        variable := instruction.Variable()
        return computer.Variable(variable)
    }

    if instruction.IsAssignment() {
        assignment := instruction.Assignment()
        return computer.Assignment(assignment)
    }

    if instruction.IsAction() {
        action := instruction.Action()
        return computer.Action(action)
    }

    lexedExecution := instruction.Execution()
    return computer.Execute(lexedExecution)
}

// ProgramToByteCode converts a parsed program to bytecodes
func (app *programAdapter) ProgramToByteCode(program Program) ([]byte, error) {
    return nil, nil
}

// ByteCodeToProgram converts bytecodes to a parsed program
func (app *programAdapter) ByteCodeToProgram(bytecode []byte) (Program, error) {
    return nil, nil
}
