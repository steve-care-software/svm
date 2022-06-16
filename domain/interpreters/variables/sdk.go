package variables

// Builder represents the variables builder
type Builder interface {
    Create() Builder
    WithList(list []Variable) Builder
    Now() (Variables, error)
}

// Variables represents variables
type Variables interface {
    List() []Variable
    Find(name string) (Variable, error)
}

// VariableBuilder represents a variable builder
type VariableBuilder interface {
    Create() VariableBuilder
    WithModule(module string) VariableBuilder
    WithName(name string) VariableBuilder
    WithKind(kind uint8) VariableBuilder
    WithValue(value interface{}) VariableBuilder
    Now() (Variable, error)
}

// Variable represents a variable
type Variable interface {
    Module() string
    Name() string
    Kind() uint8
    Value() interface{}
}
