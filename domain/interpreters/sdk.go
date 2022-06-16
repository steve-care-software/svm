package interpreters

import (
    "github.com/steve-care-software/svm/domain/parsers"
)

// AssignFn represents the assign func
type AssignFn func(input parsers.Variables, value string) (parsers.Variables, error)

// ExecuteFn represents the execute func
type ExecuteFn func(input parsers.Variables, application string) (parsers.Variables, error)

// NewModulesBuilder creates a new modules builder
func NewModulesBuilder() ModulesBuilder {
    return createModulesBuilder()
}

// NewModuleBuilder creates a new module builder
func NewModuleBuilder() ModuleBuilder {
    return createModuleBuilder()
}

// NewWatchesBuilder creates a new watches builder
func NewWatchesBuilder() WatchesBuilder {
    return createWatchesBuilder()
}

// NewWatchBuilder creates a new watch builder
func NewWatchBuilder() WatchBuilder {
    return createWatchBuilder()
}

// NewEventBuilder creates a new event builder
func NewEventBuilder() EventBuilder {
    return createEventBuilder()
}

// NewEventDefinitionBuilder creates a new event definition builder
func NewEventDefinitionBuilder() EventDefinitionBuilder {
    return createEventDefinitionBuilder()
}

// ModulesBuilder represents the modules builder
type ModulesBuilder interface {
    Create() ModulesBuilder
    WithList(list []Module) ModulesBuilder
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
    HasEvent() bool
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
    Find(module string) (Watch, error)
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
    WithExecute(execute ExecuteFn) EventDefinitionBuilder
    WithAssign(assign AssignFn) EventDefinitionBuilder
    Now() (EventDefinition, error)
}

// EventDefinition represents an event definition
type EventDefinition interface {
    Execute() ExecuteFn
    HasAssign() bool
    Assign() AssignFn
}
