package applications

import (
    "github.com/steve-care-software/svm/domain/lexers"
    "github.com/steve-care-software/svm/domain/parsers"
    "github.com/steve-care-software/svm/domain/interpreters"
)

type application struct {
    lexerAdapter lexers.ProgramAdapter
    parserAdapter parsers.ProgramAdapter
    variableBuilder parsers.VariableBuilder
    variablesBuilder parsers.VariablesBuilder
}

func createApplication(
    lexerAdapter lexers.ProgramAdapter,
    parserAdapter parsers.ProgramAdapter,
    variableBuilder parsers.VariableBuilder,
    variablesBuilder parsers.VariablesBuilder,
    ) Application {
    out := application{
        lexerAdapter: lexerAdapter,
        parserAdapter: parserAdapter,
        variableBuilder: variableBuilder,
        variablesBuilder: variablesBuilder,
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
func (app *application) Execute(params map[string]string, modules interpreters.Modules, program parsers.Program) (parsers.Variables, error) {
    inputList := []parsers.Variable{}
    output := map[string]parsers.Variable{}
    if program.HasParameters() {
        parameters := program.Parameters().List()
        for _, oneParameter := range(parameters) {
            if oneParameter.IsInput() {
                variable := oneParameter.Declaration()
                name := variable.Name()
                if content, ok := params[name]; ok {
                    kind := variable.Kind()
                    ins, err := app.variableBuilder.Create().WithKind(kind).WithName(name).WithContent(content).Now()
                    if err != nil {
                        return nil, err
                    }

                    inputList = append(inputList, ins)
                    continue
                }

                output[name] = variable
            }
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

    executions := program.Executions().List()
    variables, err := app.executions(executions, modules, input, output)
    if err != nil {
        return nil, err
    }

    outputList := []parsers.Variable{}
    for _, oneVariable := range(output) {
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

func (app *application) executions(executions []parsers.Execution, modules interpreters.Modules, variables parsers.Variables, output map[string]parsers.Variable) (parsers.Variables, error) {
    for _, oneExecution := range(executions) {
        variablesAfterExec, err := app.execution(oneExecution, modules, variables)
        if err != nil {
            return nil, err
        }

        variables = variablesAfterExec
    }

    return variables, nil
}

func (app *application) execution(execution parsers.Execution, modules interpreters.Modules, variables parsers.Variables) (parsers.Variables, error) {
    execOutput, err := app.application(execution, modules, variables)
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

func (app *application) application(execution parsers.Execution, modules interpreters.Modules, variables parsers.Variables) (parsers.Variable, error) {
    application := execution.Application()
    appVar := application.Application()
    moduleName := appVar.Kind().Module()
    module, err := modules.Find(moduleName)
    if err != nil {
        return nil, err
    }

    if !module.HasEvent() {
        return nil, nil
    }

    eventFn := module.Event()
    appName := application.Application().Name()
    attachments := application.Attachments()
    err = app.watch(moduleName, appName, modules, attachments, nil, true)
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

    err = app.watch(moduleName, appName, modules, attachments, outputVariable, false)
    if err != nil {
        return nil, err
    }

    return outputVariable, nil
}

func (app *application) watch(moduleName string, appName  string, modules interpreters.Modules, attachments parsers.Variables, execOutput parsers.Variable, isEnter bool) error {
    modulesList := modules.List()
    for _, oneModule := range(modulesList) {
        if !oneModule.HasWatches() {
            continue
        }

        watchList, err := oneModule.Watches().Find(moduleName)
        if err != nil {
            continue
        }

        for _, oneWatch := range(watchList) {
            event := oneWatch.Event()
            err := app.watchEvent(appName, event, attachments, execOutput, isEnter)
            if err != nil {
                return err
            }
        }
    }

    return nil
}

func (app *application) watchEvent(appName string, event interpreters.WatchEvent, attachments parsers.Variables, execOutput parsers.Variable, isEnter bool) error {
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
