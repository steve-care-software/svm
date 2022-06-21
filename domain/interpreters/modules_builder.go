package interpreters

import (
	"errors"
	"fmt"
)

type modulesBuilder struct {
	list []Module
}

func createModulesBuilder() ModulesBuilder {
	out := modulesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *modulesBuilder) Create() ModulesBuilder {
	return createModulesBuilder()
}

// WithList adds a list to the builder
func (app *modulesBuilder) WithList(list []Module) ModulesBuilder {
	app.list = list
	return app
}

// Now builds a new Modules instance
func (app *modulesBuilder) Now() (Modules, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Module in order to build a Modules instance")
	}

	mp := map[string]Module{}
	for _, oneModule := range app.list {
		keyname := oneModule.Name()
		mp[keyname] = oneModule
	}

	if len(mp) != len(app.list) {
		diff := len(app.list) - len(mp)
		str := fmt.Sprintf("%d Module instances were duplicates", diff)
		return nil, errors.New(str)
	}

	return createModules(app.list, mp), nil
}
