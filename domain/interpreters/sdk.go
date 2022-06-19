package interpreters

import (
    "github.com/steve-care-software/svm/domain/parsers"
)

// ExecuteFn represents the execute func
type ExecuteFn func(input parsers.Variables, application string) (string, error)

// EnterWatchFn represents the enterWatchFn func
type EnterWatchFn func(input parsers.Variables, execOutput parsers.Variable, application string) error

// ExitWatchFn represents the exitWatchFn func
type ExitWatchFn func(input parsers.Variables, application string) error

// NewModulesBuilder creates a new modules builder
func NewModulesBuilder() ModulesBuilder {
    return createModulesBuilder()
}

// NewModuleBuilder creates a new module builder
func NewModuleBuilder() ModuleBuilder {
    return createModuleBuilder()
}

// NewEventBuilder creates a new event builder
func NewEventBuilder() EventBuilder {
    return createEventBuilder()
}

// NewWatchesBuilder creates a new watches builder
func NewWatchesBuilder() WatchesBuilder {
    return createWatchesBuilder()
}

// NewWatchBuilder creates a new watch builder
func NewWatchBuilder() WatchBuilder {
    return createWatchBuilder()
}

// NewWatchEventBuilder creates a new watch event builder
func NewWatchEventBuilder() WatchEventBuilder {
    return createWatchEventBuilder()
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
    WithEvent(event ExecuteFn) ModuleBuilder
    WithWatches(watches Watches) ModuleBuilder
    Now() (Module, error)
}

// Module represents a module
type Module interface {
    Name() string
    HasEvent() bool
    Event() ExecuteFn
    HasWatches() bool
    Watches() Watches
}

// EventBuilder represents an event builder
type EventBuilder interface {
    Create() EventBuilder
    WithEnter(enter ExecuteFn) EventBuilder
    WithExit(exit ExecuteFn) EventBuilder
    Now() (Event, error)
}

// Event represents an event
type Event interface {
    HasEnter() bool
    Enter() ExecuteFn
    HasExit() bool
    Exit() ExecuteFn
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
    Find(module string) ([]Watch, error)
}

// WatchBuilder represents a watch builder
type WatchBuilder interface {
    Create() WatchBuilder
    WithModule(module string) WatchBuilder
    WithEvent(event WatchEvent) WatchBuilder
    Now() (Watch, error)
}

// Watch represents a watch
type Watch interface {
    Module() string
    Event() WatchEvent
}

// WatchEventBuilder represents a watch event builder
type WatchEventBuilder interface {
    Create() WatchEventBuilder
    WithEnter(enter EnterWatchFn) WatchEventBuilder
    WithExit(exit ExitWatchFn) WatchEventBuilder
    Now() (WatchEvent, error)
}

// WatchEvent represents a watch event
type WatchEvent interface {
    HasEnter() bool
    Enter() EnterWatchFn
    HasExit() bool
    Exit() ExitWatchFn
}
