package modules

import (
    "github.com/steve-care-software/svm/domain/interpreters/variables"
)

// AssignFn represents the assign func
type AssignFn func(input variables.Variables, value string) (variables.Variables, error)

// ExecuteFn represents the execute func
type ExecuteFn func(input variables.Variables, application string) (variables.Variables, error)

// Builder represents the modules builder
type Builder interface {
    Create() Builder
    WithList(list []Module) Builder
    Now() (Modules, error)
}

// Modules represents a modules list
type Modules interface {
    List() []Module
    Find(name string) (Module, error)
}

// ModuleBuilder represents a module builder
type ModuleBuilder interface {
    Create() ModuleBuilder
    WithName(name string) ModuleBuilder
    WithEvent(event Event) ModuleBuilder
    WithWatches(watches Watches) ModuleBuilder
    Now() (Module, error)
}

// Module represents a module
type Module interface {
    Name() string
    Event() Event
    HasWatches() bool
    Watches() Watches
}

// WatchesBuilder represents the watches builder
type WatchesBuilder interface {
    Create() WatchesBuilder
    WithList(list []Watch) WatchesBuilder
    Now() (Watches, error)
}

// Watches represents watches
type Watches interface {
    List() []Watch
    Find(name string) (Watch, error)
}

// WatchBuilder represents a watch builder
type WatchBuilder interface {
    Create() WatchBuilder
    WithModule(module string) WatchBuilder
    WithEvent(event Event) WatchBuilder
    Now() (Watch, error)
}

// Watch represents a watch
type Watch interface {
    Module() string
    Event() Event
}

// EventBuilder represents an event builder
type EventBuilder interface {
    Create() EventBuilder
    WithEnter(enter EventDefinition) EventBuilder
    WithExit(exit EventDefinition) EventBuilder
    Now() (Event, error)
}

// Event represents an event
type Event interface {
    HasEnter() bool
    Enter() EventDefinition
    HasExit() bool
    Exit() EventDefinition
}

// EventDefinitionBuilder represents an event definition builder
type EventDefinitionBuilder interface {
    Create() EventDefinitionBuilder
    WithAssign(assign AssignFn) EventDefinitionBuilder
    WithExecute(execute ExecuteFn) EventDefinitionBuilder
    Now() (EventDefinition, error)
}

// EventDefinition represents an event definition
type EventDefinition interface {
    Assign() AssignFn
    Execute() ExecuteFn
}
