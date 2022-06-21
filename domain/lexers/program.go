package lexers

type program struct {
	instructions []Instruction
}

func createProgram(
	instructions []Instruction,
) Program {
	return createProgramInternally(instructions)
}

func createProgramInternally(
	instructions []Instruction,
) Program {
	out := program{
		instructions: instructions,
	}

	return &out
}

// Instructions returns the instructions
func (obj *program) Instructions() []Instruction {
	return obj.instructions
}
