package lexers

const (
	// KindData represents the data type
	KindData uint8 = 1 << iota

	// KindApplication represents the application type
	KindApplication
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// NewProgramAdapter creates a new program adapter
func NewProgramAdapter() ProgramAdapter {
	builder := NewProgramBuilder()
	parameterBuilder := NewParameterBuilder()
	instructionBuilder := NewInstructionBuilder()
	executionBuilder := NewExecutionBuilder()
	actionBuilder := NewActionBuilder()
	scopeBuilder := NewScopeBuilder()
	assignmentBuilder := NewAssignmentBuilder()
	assigneeBuilder := NewAssigneeBuilder()
	variableBuilder := NewVariableBuilder()
	kindBuilder := NewKindBuilder()
	commentPrefix := "//"
	moduleKeyname := "module"
    typeKeyname := "type"
    dataKeyname := "data"
    inputKeyname := "->"
    outputKeyname := "<-"
    applicationKeyname := "application"
    attachKeyname := "attach"
    detachKeyname := "detach"
    toKeyname := "@"
    fromKeyname := "@"
    executeKeyname := "execute"
    moduleNameCharacters := []byte(letters)
    typeCharacters := []byte(letters)
    variableCharacters := []byte(letters)
	channelCharacters := []byte{
		[]byte("\t")[0],
		[]byte("\n")[0],
		[]byte("\r")[0],
		[]byte(" ")[0],
	}

    scopeDelimiter := []byte(":")[0]
    lineDelimiter := []byte(";")[0]
    escapeDelimiter := []byte("\\")[0]
    assignmentDelimiter :=  []byte("=")[0]
	moduleTypeDelimiter := []byte(".")[0]
	variableNameUsage := []byte("$")[0]
	return createProgramAdapter(
		builder,
		parameterBuilder,
		instructionBuilder,
		executionBuilder,
		actionBuilder,
		scopeBuilder,
		assignmentBuilder,
		assigneeBuilder,
		variableBuilder,
		kindBuilder,
		commentPrefix,
		moduleKeyname,
	    typeKeyname,
	    dataKeyname,
	    inputKeyname,
	    outputKeyname,
	    applicationKeyname,
	    attachKeyname,
	    detachKeyname,
	    toKeyname,
	    fromKeyname,
	    executeKeyname,
	    moduleNameCharacters,
	    typeCharacters,
	    variableCharacters,
	    channelCharacters,
	    scopeDelimiter,
	    lineDelimiter,
	    escapeDelimiter,
	    assignmentDelimiter,
		moduleTypeDelimiter,
		variableNameUsage,
	)
}

// NewProgramBuilder creates a new program builder
func NewProgramBuilder() ProgramBuilder {
    return createProgramBuilder()
}

// NewParameterBuilder creates a new parameter builder
func NewParameterBuilder() ParameterBuilder {
	return createParameterBuilder()
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
    return createInstructionBuilder()
}

// NewExecutionBuilder creates a new execution builder
func NewExecutionBuilder() ExecutionBuilder {
    return createExecutionBuilder()
}

// NewActionBuilder creates a new action builder
func NewActionBuilder() ActionBuilder {
    return createActionBuilder()
}

// NewScopeBuilder creates a new scope builder
func NewScopeBuilder() ScopeBuilder {
    return createScopeBuilder()
}

// NewAssignmentBuilder creates a new assignment builder
func NewAssignmentBuilder() AssignmentBuilder {
    return createAssignmentBuilder()
}

// NewAssigneeBuilder creates a new assignee builder
func NewAssigneeBuilder() AssigneeBuilder {
	return createAssigneeBuilder()
}

// NewVariableBuilder creates a new variable builder
func NewVariableBuilder() VariableBuilder {
    return createVariableBuilder()
}

// NewKindBuilder creates a new kind builder
func NewKindBuilder() KindBuilder {
    return createKindBuilder()
}


// ProgramAdapter represents a program adapter
type ProgramAdapter interface {
	ToProgram(script string) (Program, []byte, error)
}

// ProgramBuilder represents a program builder
type ProgramBuilder interface {
	Create() ProgramBuilder
	WithInstructions(instructions []Instruction) ProgramBuilder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Instructions() []Instruction
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithParameter(parameter Parameter) InstructionBuilder
	WithModule(module string) InstructionBuilder
	WithKind(kind Kind) InstructionBuilder
	WithVariable(variable Variable) InstructionBuilder
	WithAssignment(assignment Assignment) InstructionBuilder
	WithAction(action Action) InstructionBuilder
	WithExecution(execution Execution) InstructionBuilder
	WithComment(comment string) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	IsParameter() bool
	Parameter() Parameter
	IsModule() bool
	Module() string
	IsKind() bool
	Kind() Kind
	IsVariable() bool
	Variable() Variable
	IsAssignment() bool
	Assignment() Assignment
	IsAction() bool
	Action() Action
	IsExecution() bool
	Execution() Execution
	IsComment() bool
	Comment() string
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

// ExecutionBuilder represents an execution builder
type ExecutionBuilder interface {
	Create() ExecutionBuilder
	WithApplication(application string) ExecutionBuilder
	WithAssignee(assignee Assignee) ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	Application() string
	HasAssignee() bool
	Assignee() Assignee
}

// ActionBuilder represents an action builder
type ActionBuilder interface {
	Create() ActionBuilder
	WithScope(scope Scope) ActionBuilder
	WithApplication(application string) ActionBuilder
	IsAttach() ActionBuilder
	Now() (Action, error)
}

// Action represents an action
type Action interface{
	Scope() Scope
	Application() string
	IsAttach() bool
}

// ScopeBuilder represents a scope builder
type ScopeBuilder interface {
	Create() ScopeBuilder
	WithProgram(program string) ScopeBuilder
	WithModule(module string) ScopeBuilder
	Now() (Scope, error)
}

// Scope represents a scope
type Scope interface {
	Program() string
	Module() string
}

// AssignmentBuilder represents an assignment builder
type AssignmentBuilder interface {
	Create() AssignmentBuilder
	WithContent(content string) AssignmentBuilder
	WithAssignee(assignee Assignee) AssignmentBuilder
	Now() (Assignment, error)
}

// Assignment represents the assignment
type Assignment interface {
	Content() string
	Assignee() Assignee
}

// AssigneeBuilder represents an assignee builder
type AssigneeBuilder interface {
	Create() AssigneeBuilder
	WithName(name string) AssigneeBuilder
	WithDeclaration(declaration Variable) AssigneeBuilder
	Now() (Assignee, error)
}

// Assignee represents an assignee
type Assignee interface {
	IsName() bool
	Name() string
	IsDeclaration() bool
	Declaration() Variable
}

// VariableBuilder represents a variable builder
type VariableBuilder interface {
	Create() VariableBuilder
	WithModule(module string) VariableBuilder
	WithKind(kind string) VariableBuilder
	WithName(name string) VariableBuilder
	Now() (Variable, error)
}

// Variable represents a variable
type Variable interface {
	Module() string
	Kind() string
	Name() string
}

// KindBuilder represents a kind builder
type KindBuilder interface {
	Create() KindBuilder
	WithFlag(flag uint8) KindBuilder
	WithModule(module string) KindBuilder
	WithName(name string) KindBuilder
	Now() (Kind, error)
}

// Kind represents a kind
type Kind interface {
    Flag() uint8
	Module() string
    Name() string
}
