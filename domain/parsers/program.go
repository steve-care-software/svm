package parsers

type program struct {
    executions Executions
    parameters Variables
}

func createProgram(
    executions Executions,
) Program {
    return createProgramInternally(executions, nil)
}

func createProgramWithParameters(
    executions Executions,
    parameters Variables,
) Program {
    return createProgramInternally(executions, parameters)
}

func createProgramInternally(
    executions Executions,
    parameters Variables,
) Program {
    out := program{
        executions: executions,
        parameters: parameters,
    }

    return &out
}

// Executions returns the executions
func (obj *program) Executions() Executions {
    return obj.executions
}

// HasParameters returns true if there is parameters, false otherwise
func (obj *program) HasParameters() bool {
    return obj.parameters != nil
}

// Parameters returns the parameters, if any
func (obj *program) Parameters() Variables {
    return obj.parameters
}
