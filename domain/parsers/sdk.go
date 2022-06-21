package parsers

import (
	"io"
	"log"

	"github.com/steve-care-software/svm/domain/lexers"
)

const moduleVariableDelimiter = "."

// NewProgramAdapter creates a new program adapter
func NewProgramAdapter(commentLogWriter io.Writer) ProgramAdapter {
	lexerAdapter := lexers.NewProgramAdapter()
	computerFactory := NewComputerFactory()
	var commentLogger *log.Logger
	if commentLogWriter != nil {
		commentLogger = log.New(commentLogWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	return createProgramAdapter(lexerAdapter, computerFactory, commentLogger)
}

// NewProgramBuilder creates a new program builder
func NewProgramBuilder() ProgramBuilder {
	return createProgramBuilder()
}

// NewComputerFactory creates a new computer factory
func NewComputerFactory() ComputerFactory {
	programBuilder := NewProgramBuilder()
	executionsBuilder := NewExecutionsBuilder()
	parametersBuilder := NewParametersBuilder()
	parameterBuilder := NewParameterBuilder()
	executionBuilder := NewExecutionBuilder()
	applicationBuilder := NewApplicationBuilder()
	variablesBuilder := NewVariablesBuilder()
	variableBuilder := NewVariableBuilder()
	return createComputerFactory(
		programBuilder,
		executionsBuilder,
		parametersBuilder,
		parameterBuilder,
		executionBuilder,
		applicationBuilder,
		variablesBuilder,
		variableBuilder,
	)
}

// NewParametersBuilder creates a new parameters builder
func NewParametersBuilder() ParametersBuilder {
	return createParametersBuilder()
}

// NewParameterBuilder creates a new parameter builder
func NewParameterBuilder() ParameterBuilder {
	return createParameterBuilder()
}

// NewExecutionsBuilder creates a new executions builder
func NewExecutionsBuilder() ExecutionsBuilder {
	return createExecutionsBuilder()
}

// NewExecutionBuilder creates a new execution builder
func NewExecutionBuilder() ExecutionBuilder {
	return createExecutionBuilder()
}

// NewApplicationBuilder creates a new application builder
func NewApplicationBuilder() ApplicationBuilder {
	return createApplicationBuilder()
}

// NewVariablesBuilder creates a new variables builder
func NewVariablesBuilder() VariablesBuilder {
	return createVariablesBuilder()
}

// NewVariableBuilder creates a new variable builder
func NewVariableBuilder() VariableBuilder {
	return createVariableBuilder()
}

// ProgramAdapter represents a program adapter
type ProgramAdapter interface {
	ToProgram(lexed lexers.Program) (Program, error)
}

// ProgramBuilder represents a program builder
type ProgramBuilder interface {
	Create() ProgramBuilder
	WithExecutions(executions Executions) ProgramBuilder
	WithParameters(parameters Parameters) ProgramBuilder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Executions() Executions
	HasParameters() bool
	Parameters() Parameters
}

// ParametersBuilder represents a parameters builder
type ParametersBuilder interface {
	Create() ParametersBuilder
	WithList(list []Parameter) ParametersBuilder
	Now() (Parameters, error)
}

// Parameters represents the parameters
type Parameters interface {
	List() []Parameter
}

// ParameterBuilder represents a parameter builder
type ParameterBuilder interface {
	Create() ParameterBuilder
	WithDeclaration(declaration Variable) ParameterBuilder
	IsInput() ParameterBuilder
	Now() (Parameter, error)
}

// Parameter represents a parameter
type Parameter interface {
	Declaration() Variable
	IsInput() bool
}

// ComputerFactory represents a computer factory
type ComputerFactory interface {
	Create() Computer
}

// Computer represents a parser computer
type Computer interface {
	Module(name string) error
	Kind(kind lexers.Kind) error
	Variable(variable lexers.Variable) error
	Assignment(assignment lexers.Assignment) error
	Action(action lexers.Action) error
	Execute(execution lexers.Execution) error
	Parameter(parameter lexers.Parameter) error
	Program() (Program, error)
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
	Find(name string) (Execution, error)
}

// ExecutionBuilder represents an execution builder
type ExecutionBuilder interface {
	Create() ExecutionBuilder
	WithApplication(application Application) ExecutionBuilder
	WithOutput(output Variable) ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	Application() Application
	HasOutput() bool
	Output() Variable
}

// ApplicationBuilder represents an application builder
type ApplicationBuilder interface {
	Create() ApplicationBuilder
	WithApplication(application Variable) ApplicationBuilder
	WithAttachments(attachments Variables) ApplicationBuilder
	Now() (Application, error)
}

// Application represents an application execution
type Application interface {
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
	Find(module string, variable string) (Variable, error)
}

// VariableBuilder represents a variable builder
type VariableBuilder interface {
	Create() VariableBuilder
	WithKind(kind lexers.Kind) VariableBuilder
	WithName(name string) VariableBuilder
	WithContent(content string) VariableBuilder
	Now() (Variable, error)
}

// Variable represents a variable
type Variable interface {
	Kind() lexers.Kind
	Name() string
	HasContent() bool
	Content() *string
}
