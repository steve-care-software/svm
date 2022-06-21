package parsers

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/svm/domain/lexers"
)

type computer struct {
	programBuilder     ProgramBuilder
	executionsBuilder  ExecutionsBuilder
	parametersBuilder  ParametersBuilder
	parameterBuilder   ParameterBuilder
	executionBuilder   ExecutionBuilder
	applicationBuilder ApplicationBuilder
	variablesBuilder   VariablesBuilder
	variableBuilder    VariableBuilder
	kinds              map[string]map[string]lexers.Kind
	variables          map[string]Variable
	attachments        map[string]map[string]Variable
	executions         []Execution
	parameters         map[string]Parameter
}

func createComputer(
	programBuilder ProgramBuilder,
	executionsBuilder ExecutionsBuilder,
	parametersBuilder ParametersBuilder,
	parameterBuilder ParameterBuilder,
	executionBuilder ExecutionBuilder,
	applicationBuilder ApplicationBuilder,
	variablesBuilder VariablesBuilder,
	variableBuilder VariableBuilder,
) Computer {
	out := computer{
		programBuilder:     programBuilder,
		executionsBuilder:  executionsBuilder,
		parametersBuilder:  parametersBuilder,
		parameterBuilder:   parameterBuilder,
		executionBuilder:   executionBuilder,
		applicationBuilder: applicationBuilder,
		variablesBuilder:   variablesBuilder,
		variableBuilder:    variableBuilder,
		kinds:              map[string]map[string]lexers.Kind{},
		variables:          map[string]Variable{},
		attachments:        map[string]map[string]Variable{},
		executions:         []Execution{},
		parameters:         map[string]Parameter{},
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
	assignee := assignment.Assignee()
	kind, name, err := app.assignee(assignee)
	if err != nil {
		return err
	}

	content := assignment.Content()
	ins, err := app.variableBuilder.Create().WithContent(content).WithKind(kind).WithName(name).Now()
	if err != nil {
		return err
	}

	app.variables[name] = ins
	return nil
}

// Assignee declares an assignee and returns the variable name
func (app *computer) assignee(assignee lexers.Assignee) (lexers.Kind, string, error) {
	if assignee.IsName() {
		name := assignee.Name()
		if variable, ok := app.variables[name]; ok {
			kind := variable.Kind()
			return kind, name, nil
		}

		str := fmt.Sprintf("the variable (%s) is undeclared and therefore cannot be used in an assignment by name", name)
		return nil, "", errors.New(str)
	}

	lexedVariable := assignee.Declaration()
	return app.validateVariable(lexedVariable)
}

// Action declares an action
func (app *computer) Action(action lexers.Action) error {
	application := action.Application()
	if applicationVariable, ok := app.variables[application]; ok {
		if applicationVariable.Kind().Flag()&lexers.KindApplication == 0 {
			str := fmt.Sprintf("the variable (%s) was expected to be of application type in order to execute the requested Action", application)
			return errors.New(str)
		}

		scope := action.Scope()
		program := scope.Program()
		if programVariable, ok := app.variables[program]; ok {
			if programVariable.Kind().Flag()&lexers.KindData == 0 {
				str := fmt.Sprintf("the variable (%s) was expected to be of data type in order to execute the requested Action", program)
				return errors.New(str)
			}

			if _, ok := app.attachments[application]; !ok {
				app.attachments[application] = map[string]Variable{}
			}

			name := scope.Module()
			if action.IsAttach() {
				// verify that there is not already an attached variable of the same name:
				if _, ok := app.attachments[application][name]; ok {
					str := fmt.Sprintf("there is already a data variable (name: %s) attached to the application variable (%s) and therefore the requested attach Action cannot be executed", name, application)
					return errors.New(str)
				}

				kind := programVariable.Kind()
				moduleVariableBuilder := app.variableBuilder.Create().WithKind(kind).WithName(name)
				if programVariable.HasContent() {
					content := programVariable.Content()
					moduleVariableBuilder.WithContent(*content)
				}

				moduleVariable, err := moduleVariableBuilder.Now()
				if err != nil {
					return err
				}

				app.attachments[application][name] = moduleVariable
				return nil
			}

			// verify if the detach exists, if not return an error:
			if _, ok := app.attachments[application][name]; !ok {
				str := fmt.Sprintf("there is no data variable (name: %s) attached to the application variable (%s) and therefore the requested detach Action cannot be executed", name, application)
				return errors.New(str)
			}

			// detach:
			delete(app.attachments[application], name)
			return nil
		}

		str := fmt.Sprintf("the data variable (%s) is undeclared and therefore cannot be used in the requested Action", program)
		return errors.New(str)
	}

	str := fmt.Sprintf("the application variable (%s) is undeclared and therefore cannot be used in the requested Action", application)
	return errors.New(str)
}

// Execute executes an execution
func (app *computer) Execute(execution lexers.Execution) error {
	application := execution.Application()
	if applicationVariable, ok := app.variables[application]; ok {
		if applicationVariable.Kind().Flag()&lexers.KindApplication == 0 {
			str := fmt.Sprintf("the variable (%s) was expected to be of application type in order to execute the requested Execution", application)
			return errors.New(str)
		}

		attachmentsList := []Variable{}
		if variables, ok := app.attachments[application]; ok {
			for _, oneVariable := range variables {
				attachmentsList = append(attachmentsList, oneVariable)
			}
		}

		applicationBuilder := app.applicationBuilder.Create().WithApplication(applicationVariable)
		if len(attachmentsList) > 0 {
			attachments, err := app.variablesBuilder.Create().WithList(attachmentsList).Now()
			if err != nil {
				return err
			}

			applicationBuilder.WithAttachments(attachments)
		}

		appIns, err := applicationBuilder.Now()
		if err != nil {
			return err
		}

		builder := app.executionBuilder.Create().WithApplication(appIns)
		if execution.HasAssignee() {
			assignee := execution.Assignee()
			_, declName, err := app.assignee(assignee)
			if err != nil {
				return err
			}

			if _, ok := app.variables[declName]; !ok {
				str := fmt.Sprintf("the output variable (name: %s) is undeclared, therefore the requested Execution (application name: %s) cannot be executed", declName, application)
				return errors.New(str)
			}

			outputVariable := app.variables[declName]
			if outputVariable.Kind().Flag()&lexers.KindData == 0 {
				str := fmt.Sprintf("the variable (%s) was expected to be of data type in order to receive the output of the requested Execution", application)
				return errors.New(str)
			}

			builder.WithOutput(outputVariable)
		}

		ins, err := builder.Now()
		if err != nil {
			return err
		}

		app.executions = append(app.executions, ins)
		return nil
	}

	str := fmt.Sprintf("the application variable (name: %s) is undeclared, therefore the requested Execution cannot be executed", application)
	return errors.New(str)
}

// Parameter adds a parameter to the builder
func (app *computer) Parameter(lexedParameter lexers.Parameter) error {
	lexedVariable := lexedParameter.Declaration()
	name := lexedVariable.Name()
	if _, ok := app.parameters[name]; ok {
		str := fmt.Sprintf("the parameter (variable name: %s) is already declared", name)
		return errors.New(str)
	}

	err := app.Variable(lexedVariable)
	if err != nil {
		return err
	}

	builder := app.parameterBuilder.Create().WithDeclaration(app.variables[name])
	if lexedParameter.IsInput() {
		builder.IsInput()
	}

	ins, err := builder.Now()
	if err != nil {
		return err
	}

	app.parameters[name] = ins
	return nil
}

// Program builds the program
func (app *computer) Program() (Program, error) {
	executions, err := app.executionsBuilder.Create().WithList(app.executions).Now()
	if err != nil {
		return nil, err
	}

	builder := app.programBuilder.Create().WithExecutions(executions)
	if len(app.parameters) > 0 {
		list := []Parameter{}
		for _, oneParameter := range app.parameters {
			list = append(list, oneParameter)
		}

		parameters, err := app.parametersBuilder.Create().WithList(list).Now()
		if err != nil {
			return nil, err
		}

		builder.WithParameters(parameters)
	}

	return builder.Now()
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
