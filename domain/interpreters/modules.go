package interpreters

import (
	"errors"
	"fmt"
)

type modules struct {
	list []Module
	mp   map[string]Module
}

func createModules(
	list []Module,
	mp map[string]Module,
) Modules {
	out := modules{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list of modules
func (obj *modules) List() []Module {
	return obj.list
}

// Find finds a module by name
func (obj *modules) Find(name string) (Module, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the module (name: %s) is undefined", name)
	return nil, errors.New(str)
}
