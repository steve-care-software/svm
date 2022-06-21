package parsers

import (
	"errors"
	"fmt"
)

type variablesBuilder struct {
	list []Variable
}

func createVariablesBuilder() VariablesBuilder {
	out := variablesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *variablesBuilder) Create() VariablesBuilder {
	return createVariablesBuilder()
}

// WithList adds a list to the builder
func (app *variablesBuilder) WithList(list []Variable) VariablesBuilder {
	app.list = list
	return app
}

// Now builds a new Variables instance
func (app *variablesBuilder) Now() (Variables, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Variable in order to build a Variables instance")
	}

	mp := map[string]Variable{}
	for _, oneVariable := range app.list {
		keyname := fmt.Sprintf(moduleVariablePatern, oneVariable.Kind().Module(), moduleVariableDelimiter, oneVariable.Name())
		mp[keyname] = oneVariable
	}

	if len(mp) != len(app.list) {
		diff := len(app.list) - len(mp)
		str := fmt.Sprintf("%d Variable instances were duplicates", diff)
		return nil, errors.New(str)
	}

	return createVariables(app.list, mp), nil
}
