package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/svm/domain/interpreters"
	"github.com/steve-care-software/svm/domain/lexers"
	"github.com/steve-care-software/svm/domain/parsers"
)

type application struct {
	lexerAdapter     lexers.ProgramAdapter
	parserAdapter    parsers.ProgramAdapter
	variableBuilder  parsers.VariableBuilder
	variablesBuilder parsers.VariablesBuilder
	modules          interpreters.Modules
}

func createApplication(
	lexerAdapter lexers.ProgramAdapter,
	parserAdapter parsers.ProgramAdapter,
	variableBuilder parsers.VariableBuilder,
	variablesBuilder parsers.VariablesBuilder,
	modules interpreters.Modules,
) Application {
	out := application{
		lexerAdapter:     lexerAdapter,
		parserAdapter:    parserAdapter,
		variableBuilder:  variableBuilder,
		variablesBuilder: variablesBuilder,
		modules:          modules,
	}

	return &out
}

// Compile compiles a script to a parsed program
func (app *application) Compile(script string) (parsers.Program, []byte, error) {
	lexedProgram, remaining, err := app.lexerAdapter.ToProgram(script)
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.parserAdapter.ToProgram(lexedProgram)
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

// Execute executes a parsed program and return the output values
func (app *application) Execute(params map[string]string, program parsers.Program) (parsers.Variables, error) {
	inputList := []parsers.Variable{}
	output := map[string]parsers.Variable{}
	if program.HasParameters() {
		parameters := program.Parameters().List()
		for _, oneParameter := range parameters {
			variable := oneParameter.Declaration()
			name := variable.Name()
			if oneParameter.IsInput() {
				if content, ok := params[name]; ok {
					kind := variable.Kind()
					ins, err := app.variableBuilder.Create().WithKind(kind).WithName(name).WithContent(content).Now()
					if err != nil {
						return nil, err
					}

					inputList = append(inputList, ins)
					continue
				}

				str := fmt.Sprintf("there is an input parameter variable (name: %s) declared in the script, but there is no input data of that name declared", name)
				return nil, errors.New(str)
			}

			output[name] = variable
		}
	}

	var input parsers.Variables
	if len(inputList) > 0 {
		ins, err := app.variablesBuilder.Create().WithList(inputList).Now()
		if err != nil {
			return nil, err
		}

		input = ins
	}

	// generate the execution params:
	attachments := map[int]map[string]string{}
	executionsList := program.Executions().List()
	for idx, oneExecution := range executionsList {
		application := oneExecution.Application()
		if application.HasAttachments() {
			attachments[idx] = map[string]string{}
			attachmentsList := application.Attachments().List()
			for _, oneAttachment := range attachmentsList {
				moduleVariable := oneAttachment.Module()
				nameInModule := moduleVariable.Name()
				attachments[idx][nameInModule] = ""
				if !moduleVariable.HasContent() && input != nil {
					program := oneAttachment.Program()
					moduleName := moduleVariable.Kind().Module()
					inputVariable, err := input.Find(moduleName, program)
					if err != nil {
						return nil, err
					}

					content := ""
					if inputVariable.HasContent() {
						pContent := inputVariable.Content()
						content = *pContent
					}

					attachments[idx][nameInModule] = content
					continue
				}

				if moduleVariable.HasContent() {
					pContent := moduleVariable.Content()
					attachments[idx][nameInModule] = *pContent
				}

			}
		}
	}

	executions := program.Executions().List()
	variables, err := app.executions(executions, input, attachments, output)
	if err != nil {
		return nil, err
	}

	outputList := []parsers.Variable{}
	for _, oneVariable := range output {
		name := oneVariable.Name()
		moduleName := oneVariable.Kind().Module()
		fetched, err := variables.Find(moduleName, name)
		if err != nil {
			return nil, err
		}

		outputList = append(outputList, fetched)
	}

	return app.variablesBuilder.Create().WithList(outputList).Now()
}

func (app *application) executions(executions []parsers.Execution, variables parsers.Variables, attachments map[int]map[string]string, output map[string]parsers.Variable) (parsers.Variables, error) {
	for idx, oneExecution := range executions {
		variablesAfterExec, err := app.execution(oneExecution, variables, attachments[idx])
		if err != nil {
			return nil, err
		}

		variables = variablesAfterExec
	}

	return variables, nil
}

func (app *application) execution(execution parsers.Execution, variables parsers.Variables, attachments map[string]string) (parsers.Variables, error) {
	execOutput, err := app.application(execution, variables, attachments)
	if err != nil {
		return nil, err
	}

	if execOutput != nil {
		list := variables.List()
		list = append(list, execOutput)
		return app.variablesBuilder.Create().WithList(list).Now()
	}

	return variables, nil
}

func (app *application) application(execution parsers.Execution, variables parsers.Variables, attachments map[string]string) (parsers.Variable, error) {
	application := execution.Application()
	appVar := application.Application()
	moduleName := appVar.Kind().Module()
	module, err := app.modules.Find(moduleName)
	if err != nil {
		return nil, err
	}

	if !module.HasEvent() {
		return nil, nil
	}

	eventFn := module.Event()
	appName := application.Application().Name()
	err = app.watch(moduleName, appName, attachments, nil, true)
	if err != nil {
		return nil, err
	}

	execOutput, err := eventFn(attachments, appName)
	if err != nil {
		return nil, err
	}

	var outputVariable parsers.Variable
	if execution.HasOutput() {
		output := execution.Output()
		kind := output.Kind()
		name := output.Name()
		builder := app.variableBuilder.Create().WithKind(kind).WithName(name)
		if execOutput != "" {
			builder.WithContent(execOutput)
		}

		ins, err := builder.Now()
		if err != nil {
			return nil, err
		}

		outputVariable = ins
	}

	err = app.watch(moduleName, appName, attachments, outputVariable, false)
	if err != nil {
		return nil, err
	}

	return outputVariable, nil
}

func (app *application) watch(moduleName string, appName string, attachments map[string]string, execOutput parsers.Variable, isEnter bool) error {
	modulesList := app.modules.List()
	for _, oneModule := range modulesList {
		if !oneModule.HasWatches() {
			continue
		}

		watchList, err := oneModule.Watches().Find(moduleName)
		if err != nil {
			continue
		}

		for _, oneWatch := range watchList {
			event := oneWatch.Event()
			err := app.watchEvent(appName, event, attachments, execOutput, isEnter)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (app *application) watchEvent(appName string, event interpreters.WatchEvent, attachments map[string]string, execOutput parsers.Variable, isEnter bool) error {
	if isEnter && event.HasEnter() {
		enterFn := event.Enter()
		err := enterFn(attachments, execOutput, appName)
		if err != nil {
			return err
		}
	}

	if !isEnter && event.HasExit() {
		exitFn := event.Exit()
		err := exitFn(attachments, appName)
		if err != nil {
			return err
		}
	}

	return nil
}
