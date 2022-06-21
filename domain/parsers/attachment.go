package parsers

type attachment struct {
	program string
	module  Variable
}

func createAttachment(
	program string,
	module Variable,
) Attachment {
	out := attachment{
		program: program,
		module:  module,
	}

	return &out
}

// Program returns the program
func (obj *attachment) Program() string {
	return obj.program
}

// Module returns the module
func (obj *attachment) Module() Variable {
	return obj.module
}
