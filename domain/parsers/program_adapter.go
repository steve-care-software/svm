package parsers

import (
    "github.com/steve-care-software/svm/domain/lexers"
    "log"
)

type programAdapter struct {
    lexerAdapter lexers.ProgramAdapter
    computerFactory ComputerFactory
    commentLogger *log.Logger
}

func createProgramAdapter(
    lexerAdapter lexers.ProgramAdapter,
    computerFactory ComputerFactory,
    commentLogger *log.Logger,
    ) ProgramAdapter {
    out := programAdapter {
        lexerAdapter: lexerAdapter,
        computerFactory: computerFactory,
        commentLogger: commentLogger,
    }

    return &out
}

// ToProgram converts a lexed program to a parsed program
func (app *programAdapter) ToProgram(lexed lexers.Program) (Program, error) {
    computer := app.computerFactory.Create()
    instructions := lexed.Instructions()
    err := app.instructions(computer, instructions)
    if err != nil {
        return nil, err
    }

    return computer.Program()
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
    if instruction.IsComment() {
        if app.commentLogger == nil {
            return nil
        }
        
        comment := instruction.Comment()
        app.commentLogger.Println(comment)
        return nil
    }

    if instruction.IsParameter() {
        parameter := instruction.Parameter()
        return computer.Parameter(parameter)
    }

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
