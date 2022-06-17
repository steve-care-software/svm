package parsers

import (
    "github.com/steve-care-software/svm/domain/lexers"
    "fmt"
    "errors"
)

type computer struct {
    variableBuilder VariableBuilder
    kinds map[string]map[string]lexers.Kind
    variables map[string]Variable
}

func createComputer(
    variableBuilder VariableBuilder,
    ) Computer {
    out :=computer{
        variableBuilder: variableBuilder,
        kinds: map[string]map[string]lexers.Kind{},
        variables: map[string]Variable{},
    }

    return &out
}

// Module declares a new module, returns an error if it already exists
func (app *computer) Module(name string) error {
    if _, ok := app.kinds[name]; ok {
        str := fmt.Sprintf("the module (name: %s) is already declared", name)
        return errors.New(str)
    }

    app.kinds[name] = map[string]lexers.Kind{}
    return nil
}

// Kind declares a new kind
func (app *computer) Kind(kind lexers.Kind) error {
    moduleName := kind.Module()
    if module, ok := app.kinds[moduleName]; ok {
        name := kind.Name()
        if _, ok := module[name]; ok {
            str := fmt.Sprintf("the type (module: %s, name: %s) is already declared", module, name)
            return errors.New(str)
        }

        app.kinds[moduleName][name] = kind
        return nil
    }

    str := fmt.Sprintf("the type (%s) is attached to an undeclared module (name: %s)", kind.Name(), moduleName)
    return errors.New(str)
}

// Variable declares a variable
func (app *computer) Variable(lexedVariable lexers.Variable) error {
    kind, name, err := app.validateVariable(lexedVariable)
    if err != nil {
        return err
    }

    variable, err := app.variableBuilder.Create().WithKind(kind).WithName(name).Now()
    if err != nil {
        return err
    }

    app.variables[name] = variable
    return nil
}

// Assignment declares an assignment
func (app *computer) Assignment(assignment lexers.Assignment) error {
    content := assignment.Content()
    if assignment.IsName() {
        name := assignment.Name()
        if variable, ok := app.variables[name]; ok {
            kind := variable.Kind()
            ins, err := app.variableBuilder.Create().WithContent(content).WithKind(kind).WithName(name).Now()
            if err != nil {
                return err
            }

            app.variables[name] = ins
        }

        str := fmt.Sprintf("the variable (%s) is undeclared and therefore cannot be used in an assignment by name", name)
        return errors.New(str)
    }

    lexedVariable := assignment.Declaration()
    kind, name, err := app.validateVariable(lexedVariable)
    if err != nil {
        return err
    }

    variable, err := app.variableBuilder.Create().WithContent(content).WithKind(kind).WithName(name).Now()
    if err != nil {
        return err
    }

    app.variables[name] = variable
    return nil
}

func (app *computer) validateVariable(variable lexers.Variable) (lexers.Kind, string, error) {
    name := variable.Name()
    moduleName := variable.Module()
    if module, ok := app.kinds[moduleName]; ok {
        kind := variable.Kind()
        if kind, ok := module[kind]; ok {
            return kind, name, nil
        }

        str := fmt.Sprintf("the variable (%s) is declared using an undeclared type (%s) in a declared module (%s)", name, kind, moduleName)
        return nil, "", errors.New(str)
    }

    str := fmt.Sprintf("the variable (%s) is attached to an undeclared module (name: %s)", name, moduleName)
    return nil, "", errors.New(str)
}
