package interpreters

import (
    "errors"
    "fmt"
)

type watches struct {
    list []Watch
    mp map[string]Watch
}

func createWatches(
    list []Watch,
    mp map[string]Watch,
) Watches {
    out := watches{
        list: list,
        mp: mp,
    }

    return &out
}

// List returns the list of watches
func (obj *watches) List() []Watch {
    return obj.list
}

// Find finds a watch by name
func (obj *watches) Find(module string) (Watch, error) {
    if ins, ok := obj.mp[module]; ok {
        return ins, nil
    }

    str := fmt.Sprintf("the watch (module: %s) is undefined", module)
    return nil, errors.New(str)
}
