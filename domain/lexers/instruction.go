package lexers

type instruction struct {
    module string
    kind Kind
    variable Variable
    assignment Assignment
    action Action
    execution Execution
}

func createInstructionWithModule(
    module string,
) Instruction {
    return createInstructionInternally(module, nil, nil, nil, nil, nil)
}

func createInstructionWithKind(
    kind Kind,
) Instruction {
    return createInstructionInternally("", kind, nil, nil, nil, nil)
}

func createInstructionWithVariable(
    variable Variable,
) Instruction {
    return createInstructionInternally("", nil, variable, nil, nil, nil)
}

func createInstructionWithAssignment(
    assignment Assignment,
) Instruction {
    return createInstructionInternally("", nil, nil, assignment, nil, nil)
}

func createInstructionWithAction(
    action Action,
) Instruction {
    return createInstructionInternally("", nil, nil, nil, action, nil)
}

func createInstructionWithExecution(
    execution Execution,
) Instruction {
    return createInstructionInternally("", nil, nil, nil, nil, execution)
}

func createInstructionInternally(
    module string,
    kind Kind,
    variable Variable,
    assignment Assignment,
    action Action,
    execution Execution,
) Instruction {
    out := instruction{
        module: module,
        kind: kind,
        variable: variable,
        assignment: assignment,
        action: action,
        execution: execution,
    }

    return &out
}

// IsModule returns true if there is a module, false otherwise
func (obj *instruction) IsModule() bool {
    return obj.module != ""
}

// Module returns the module, if any
func (obj *instruction) Module() string {
    return obj.module
}

// IsKind returns true if there is a kind, false otherwise
func (obj *instruction) IsKind() bool {
    return obj.kind != nil
}

// Kind returns the kind, if any
func (obj *instruction) Kind() Kind {
    return obj.kind
}

// IsVariable returns true if there is a variable, false otherwise
func (obj *instruction) IsVariable() bool {
    return obj.variable != nil
}

// Variable returns the variable, if any
func (obj *instruction) Variable() Variable {
    return obj.variable
}

// IsAssignment returns true if there is an assignment, false otherwise
func (obj *instruction) IsAssignment() bool {
    return obj.assignment != nil
}

// Assignment returns the assignment, if any
func (obj *instruction) Assignment() Assignment {
    return obj.assignment
}

// IsAction returns true if there is an action, false otherwise
func (obj *instruction) IsAction() bool {
    return obj.action != nil
}

// Action returns the action, if any
func (obj *instruction) Action() Action {
    return obj.action
}

// IsExecution returns true if there is an execution, false otherwise
func (obj *instruction) IsExecution() bool {
    return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *instruction) Execution() Execution {
    return obj.execution
}
