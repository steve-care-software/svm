package lexers

type program struct {
    instructions  []Instruction
    parameters []Variable
}

func createProgram(
    instructions  []Instruction,
) Program {
    return createProgramInternally(instructions, nil)
}

func createProgramWithParameters(
    instructions  []Instruction,
    parameters []Variable,
) Program {
    return createProgramInternally(instructions, parameters)
}

func createProgramInternally(
    instructions  []Instruction,
    parameters []Variable,
) Program {
    out := program{
        instructions: instructions,
        parameters: parameters,
    }

    return &out
}

// Instructions returns the instructions
func (obj *program) Instructions() []Instruction {
    return obj.instructions
}

// HasParameters returns true if there is parameters, false otherwise
func (obj *program) HasParameters() bool {
    return obj.parameters != nil
}

// Parameters returns the parameters, if any
func (obj *program) Parameters() []Variable {
    return obj.parameters
}