package interpreters

import (
	"errors"
)

type moduleBuilder struct {
	name    string
	event   ExecuteFn
	watches Watches
}

func createModuleBuilder() ModuleBuilder {
	out := moduleBuilder{
		name:    "",
		event:   nil,
		watches: nil,
	}

	return &out
}

// Create initializes the builder
func (app *moduleBuilder) Create() ModuleBuilder {
	return createModuleBuilder()
}

// WithName adds a name to the builder
func (app *moduleBuilder) WithName(name string) ModuleBuilder {
	app.name = name
	return app
}

// WithEvent adds an event to the builder
func (app *moduleBuilder) WithEvent(event ExecuteFn) ModuleBuilder {
	app.event = event
	return app
}

// WithWatches add watches to the builder
func (app *moduleBuilder) WithWatches(watches Watches) ModuleBuilder {
	app.watches = watches
	return app
}

// Now builds a new Module instance
func (app *moduleBuilder) Now() (Module, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Module instance")
	}

	if app.event != nil && app.watches != nil {
		return createModuleWithEventAndWatches(app.name, app.event, app.watches), nil
	}

	if app.event == nil {
		return createModuleWithEvent(app.name, app.event), nil
	}

	if app.watches != nil {
		return createModuleWithWatches(app.name, app.watches), nil
	}

	return nil, errors.New("the Module is invalid")
}
