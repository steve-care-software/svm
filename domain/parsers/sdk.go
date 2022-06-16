package parsers

import (
    "github.com/steve-care-software/svm/domain/lexers"
)

// ProgramAdapter represents a program adapter
type ProgramAdapter interface {
	LexedToProgram(lexed lexers.Program) (Program, error)
    ProgramToByteCode(program Program) ([]byte, error)
    ByteCodeToProgram(bytecode []byte) (Program, error)
}

// ProgramBuilder represents a program builder
type ProgramBuilder interface {
	Create() ProgramBuilder
	WithExecutions(executions Executions) ProgramBuilder
	WithParameters(parameters Variables) ProgramBuilder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Executions() Executions
	HasParameters() bool
	Parameters() Variables
}

// ExecutionsBuilder represents an executions builder
type ExecutionsBuilder interface {
	Create() ExecutionsBuilder
	WithList(list []Execution) ExecutionsBuilder
	Now() (Executions, error)
}

// Executions represents executions
type Executions interface {
	List() []Execution
}

// ExecutionBuilder represents an execution builder
type ExecutionBuilder interface {
	Create() ExecutionBuilder
	WithModule(module string) ExecutionBuilder
	WithApplication(application Variable) ExecutionBuilder
	WithAttachments(attachments Variables) ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents a program execution
type Execution interface {
	Module() string
	Application() Variable
	HasAttachments() bool
	Attachments() Variables
}

// VariablesBuilder represents variables builder
type VariablesBuilder interface {
	Create() VariablesBuilder
	WithList(list []Variable) VariablesBuilder
	Now() (Variables, error)
}

// Variables represents variables
type Variables interface {
	List() []Variable
}

// VariableBuilder represents a variable builder
type VariableBuilder interface {
	Create() VariableBuilder
	WithKind(kind lexers.Kind) VariableBuilder
	WithName(name string) VariableBuilder
	WithValue(value interface{}) VariableBuilder
	Now() (Variable, error)
}

// Variable represents a variable
type Variable interface {
	Kind() lexers.Kind
	Name() string
	HasValue() bool
	Value() interface{}
}
