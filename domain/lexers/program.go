package lexers

type program struct {
    instructions  []Instruction
    parameters []Parameter
}

func createProgram(
    instructions  []Instruction,
) Program {
    return createProgramInternally(instructions, nil)
}

func createProgramWithParameters(
    instructions  []Instruction,
    parameters []Parameter,
) Program {
    return createProgramInternally(instructions, parameters)
}

func createProgramInternally(
    instructions  []Instruction,
    parameters []Parameter,
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
func (obj *program) Parameters() []Parameter {
    return obj.parameters
}
