package parsers

import (
    "errors"
    "fmt"
)

type executions struct {
    list []Execution
    mp map[string]Execution
}

func createExecutions(
    list []Execution,
    mp map[string]Execution,
) Executions {
    out := executions{
        list: list,
        mp: mp,
    }

    return &out
}

// List returns the list of executions
func (obj *executions) List() []Execution {
    return obj.list
}

// Find finds a execution by name
func (obj *executions) Find(name string) (Execution, error) {
    if ins, ok := obj.mp[name]; ok {
        return ins, nil
    }

    str := fmt.Sprintf("the execution (name: %s) is undefined", name)
    return nil, errors.New(str)
}
