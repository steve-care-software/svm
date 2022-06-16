package parsers

import (
    "errors"
    "fmt"
)

type variables struct {
    list []Variable
    mp map[string]Variable
}

func createVariables(
    list []Variable,
    mp map[string]Variable,
) Variables {
    out := variables{
        list: list,
        mp: mp,
    }

    return &out
}

// List returns the list of variables
func (obj *variables) List() []Variable {
    return obj.list
}

// Find finds a variable by name
func (obj *variables) Find(module string, variable string) (Variable, error) {
    keyname := fmt.Sprintf("%s%s%s", module, moduleVariableDelimiter, variable)
    if ins, ok := obj.mp[keyname]; ok {
        return ins, nil
    }

    str := fmt.Sprintf("the variable (module: %s, variable: %s) is undefined", module, variable)
    return nil, errors.New(str)
}
