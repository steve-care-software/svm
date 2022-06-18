package parsers

type parameter struct {
    declaration Variable
    isInput bool
}

func createParameter(
    declaration Variable,
    isInput bool,
) Parameter {
    out := parameter{
        declaration: declaration,
        isInput: isInput,
    }

    return &out
}

// Declaration returns the variable declaration
func (obj *parameter) Declaration() Variable {
    return obj.declaration
}

// IsInput returns true if input, false otherwise
func (obj *parameter) IsInput() bool {
    return obj.isInput
}
