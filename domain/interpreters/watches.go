package interpreters

import (
	"errors"
	"fmt"
)

type watches struct {
	list []Watch
	mp   map[string][]Watch
}

func createWatches(
	list []Watch,
	mp map[string][]Watch,
) Watches {
	out := watches{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list of watches
func (obj *watches) List() []Watch {
	return obj.list
}

// Find finds a watch by name
func (obj *watches) Find(module string) ([]Watch, error) {
	if list, ok := obj.mp[module]; ok {
		return list, nil
	}

	str := fmt.Sprintf("there is no watch for module (name: %s)", module)
	return nil, errors.New(str)
}
