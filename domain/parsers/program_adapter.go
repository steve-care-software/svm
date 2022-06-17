package parsers

import (
    "github.com/steve-care-software/svm/domain/lexers"
)

type programAdapter struct {
    lexerAdapter lexers.ProgramAdapter
    computer Computer
}

func createProgramAdapter(
    lexerAdapter lexers.ProgramAdapter,
    computer Computer,
    ) ProgramAdapter {
    out := programAdapter {
        lexerAdapter: lexerAdapter,
        computer: computer,
    }

    return &out
}

// LexedToProgram converts a lexed program to a parsed program
func (app *programAdapter) LexedToProgram(lexed lexers.Program) (Program, error) {
    return nil, nil
}

func (app *programAdapter) instruction(instruction lexers.Instruction, variablesByModule map[string]map[string]Variable, variables map[string]Variable) error {
    if instruction.IsModule() {
        module := instruction.Module()
        return app.computer.Module(module)
    }

    if instruction.IsKind() {
        kind := instruction.Kind()
        return app.computer.Kind(kind)
    }

    if instruction.IsVariable() {
        variable := instruction.Variable()
        return app.computer.Variable(variable)
    }

    if instruction.IsAssignment() {
        assignment := instruction.Assignment()
        return app.computer.Assignment(assignment)
    }

    if instruction.IsAction() {

    }

    if instruction.IsExecution() {

    }

    return nil
}

// ProgramToByteCode converts a parsed program to bytecodes
func (app *programAdapter) ProgramToByteCode(program Program) ([]byte, error) {
    return nil, nil
}

// ByteCodeToProgram converts bytecodes to a parsed program
func (app *programAdapter) ByteCodeToProgram(bytecode []byte) (Program, error) {
    return nil, nil
}
