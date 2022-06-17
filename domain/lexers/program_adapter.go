package lexers

import (
    "bytes"
    "fmt"
    "errors"
)

type programAdapter struct {
    builder ProgramBuilder
    parameterBuilder ParameterBuilder
    instructionBuilder InstructionBuilder
    executionBuilder ExecutionBuilder
    actionBuilder ActionBuilder
    scopeBuilder ScopeBuilder
    assignmentBuilder AssignmentBuilder
    variableBuilder VariableBuilder
    kindBuilder KindBuilder
    moduleKeyname string
    typeKeyname string
    dataKeyname string
    inputKeyname string
    outputKeyname string
    applicationKeyname string
    attachKeyname string
    detachKeyname string
    toKeyname string
    fromKeyname string
    executeKeyname string
    moduleNameCharacters []byte
    typeCharacters []byte
    variableCharacters []byte
    channelCharacters []byte
    scopeDelimiter byte
    lineDelimiter byte
    escapeDelimiter byte
    assignmentDelimiter byte
    moduleTypeDelimiter byte
    variableNameUsage byte
}

func createProgramAdapter(
    builder ProgramBuilder,
    parameterBuilder ParameterBuilder,
    instructionBuilder InstructionBuilder,
    executionBuilder ExecutionBuilder,
    actionBuilder ActionBuilder,
    scopeBuilder ScopeBuilder,
    assignmentBuilder AssignmentBuilder,
    variableBuilder VariableBuilder,
    kindBuilder KindBuilder,
    moduleKeyname string,
    typeKeyname string,
    dataKeyname string,
    inputKeyname string,
    outputKeyname string,
    applicationKeyname string,
    attachKeyname string,
    detachKeyname string,
    toKeyname string,
    fromKeyname string,
    executeKeyname string,
    moduleNameCharacters []byte,
    typeCharacters []byte,
    variableCharacters []byte,
    channelCharacters []byte,
    scopeDelimiter byte,
    lineDelimiter byte,
    escapeDelimiter byte,
    assignmentDelimiter byte,
    moduleTypeDelimiter byte,
    variableNameUsage byte,
    ) ProgramAdapter {
    out := programAdapter {
        builder: builder,
        parameterBuilder: parameterBuilder,
        instructionBuilder: instructionBuilder,
        executionBuilder: executionBuilder,
        actionBuilder: actionBuilder,
        scopeBuilder: scopeBuilder,
        assignmentBuilder: assignmentBuilder,
        variableBuilder: variableBuilder,
        kindBuilder: kindBuilder,
        moduleKeyname: moduleKeyname,
        typeKeyname: typeKeyname,
        dataKeyname: dataKeyname,
        inputKeyname: inputKeyname,
        outputKeyname: outputKeyname,
        applicationKeyname: applicationKeyname,
        attachKeyname: attachKeyname,
        detachKeyname: detachKeyname,
        toKeyname: toKeyname,
        fromKeyname: fromKeyname,
        executeKeyname: executeKeyname,
        moduleNameCharacters: moduleNameCharacters,
        typeCharacters: typeCharacters,
        variableCharacters: variableCharacters,
        channelCharacters: channelCharacters,
        scopeDelimiter: scopeDelimiter,
        lineDelimiter: lineDelimiter,
        escapeDelimiter: escapeDelimiter,
        assignmentDelimiter: assignmentDelimiter,
        moduleTypeDelimiter: moduleTypeDelimiter,
        variableNameUsage: variableNameUsage,
    }

    return &out
}

// ScriptToProgram converts a script to a Program instance
func (app *programAdapter) ScriptToProgram(script string) (Program, []byte, error) {
    // convert to bytes:
	bytes := []byte(script)

    // return the program:
    return app.program(bytes)
}

func (app *programAdapter) program(input []byte) (Program, []byte, error) {
    // retrieve the parameters:
    parameters, remainingafterParameters, err := app.parameterDeclarations(input)
    if err != nil {
        return nil, nil, err
    }

    if parameters == nil {
        remainingafterParameters = input
    }

    // retrieve the instructions:
    instructions, remaining, err := app.instructions(remainingafterParameters)
    if err != nil {
        return nil, nil, err
    }

    // build the program:
    builder := app.builder.Create().WithInstructions(instructions)
    if parameters != nil {
        builder.WithParameters(parameters)
    }

    ins, err := builder.Now()
    if err != nil {
        return nil, nil, err
    }

    return ins, app.removeChannelCharactersPrefix(remaining), nil
}

func (app *programAdapter) parameterDeclarations(input []byte) ([]Parameter, []byte, error) {
    cpt := uint(0)
    list := []Parameter{}
    remaining := input
    for {
        parameter, remainingAfterRemaining, err := app.parameterDeclaration(remaining, cpt)
        if err != nil {
            return nil, nil, err
        }

        if parameter == nil {
            break
        }

        remaining = remainingAfterRemaining
        list = append(list, parameter)
        cpt++
    }

    return list, remaining, nil
}

func (app *programAdapter) parameterDeclaration(input []byte, index uint) (Parameter, []byte, error) {
    builder := app.parameterBuilder.Create()
    hasInput, remaining := app.hasPrefix(input, app.inputKeyname)
    if !hasInput {
        hasOutput, remainingAfterOutput := app.hasPrefix(remaining, app.outputKeyname)
        if !hasOutput {
            return nil, nil, nil
        }

        remaining = remainingAfterOutput
    }

    if hasInput {
        builder.IsInput()
    }

    variable, remainingAfterVariable := app.variableDeclaration(remaining)
    if variable != nil {
        builder.WithDeclaration(variable)
    }

    ins, err := builder.Now()
    if err != nil {
        return nil, nil, nil
    }

    hasLineDelimiter, remainingAfterLineDelimiter := app.hasPrefix(remainingAfterVariable, string(app.lineDelimiter))
    if !hasLineDelimiter {
        str := fmt.Sprintf("the line delimiter (%s) was expected after the parameter (index: %d)", string(app.lineDelimiter), index)
        return nil, nil, errors.New(str)
    }

    return ins, remainingAfterLineDelimiter, nil
}

func (app *programAdapter) instructions(input []byte) ([]Instruction, []byte, error) {
    cpt := uint(0)
    remaining := input
    list := []Instruction{}
    for {
        if len(remaining) <= 0 {
            break
        }

        ins, remainingAfterIns, err := app.instruction(remaining, cpt)
        if err != nil {
            break
        }

        remaining = remainingAfterIns
        list = append(list, ins)
        cpt++;
    }

    return list, remaining, nil
}

func (app *programAdapter) instruction(input []byte, index uint) (Instruction, []byte, error) {
    found := false
    remaining := input
    builder := app.instructionBuilder.Create()
    moduleName, remainingAfterModule := app.moduleName(input)
    if moduleName != "" {
        found = true
        remaining = remainingAfterModule
        builder.WithModule(moduleName)
    }

    if !found {
        kind, remainingAfterKind := app.typeDeclaration(input)
        if kind != nil {
            found = true
            remaining = remainingAfterKind
            builder.WithKind(kind)
        }
    }

    if !found {
        execution, remainingAfterExecution := app.execution(input)
        if execution != nil {
            found = true
            remaining = remainingAfterExecution
            builder.WithExecution(execution)
        }
    }

    // this block must be after the execution block:
    if !found {
        assignment, remainingAfterAssignment := app.assignment(input)
        if assignment != nil {
            found = true
            remaining = remainingAfterAssignment
            builder.WithAssignment(assignment)
        }
    }

    // this block must be after the assignment block:
    if !found {
        variable, remainingAfterVariable := app.variableDeclaration(input)
        if variable != nil {
            found = true
            remaining = remainingAfterVariable
            builder.WithVariable(variable)
        }
    }

    if !found {
        action, remainingAfterAction := app.action(input)
        if action != nil {
            remaining = remainingAfterAction
            builder.WithAction(action)
        }
    }

    ins, err := builder.Now()
    if err != nil {
        return nil, nil, err
    }

    hasLineDelimiter, remainingAfterLineDelimiter := app.hasPrefix(remaining, string(app.lineDelimiter))
    if !hasLineDelimiter {
        str := fmt.Sprintf("the line delimiter (%s) was expected after the instruction (index: %d)", string(app.lineDelimiter),  index)
        return nil, nil, errors.New(str)
    }

    return ins, remainingAfterLineDelimiter, nil
}

func (app *programAdapter) moduleName(input []byte) (string, []byte) {
    hasModule, remaining := app.hasPrefix(input, app.moduleKeyname)
    if !hasModule {
        return "", nil
    }

    return app.fetchName(remaining, app.moduleNameCharacters)
}

func (app *programAdapter) typeDeclaration(input []byte) (Kind, []byte) {
    hasType, remainingAfterType := app.hasPrefix(input, app.typeKeyname)
    if !hasType {
        return nil, nil
    }

    kind, remaining := app.applicationTypeDeclaration(remainingAfterType)
    if kind != nil {
        return kind, remaining
    }

    return app.dataTypeDeclaration(remainingAfterType)
}

func (app *programAdapter) applicationTypeDeclaration(input []byte) (Kind, []byte) {
    hasApplication, remaining := app.hasPrefix(input, app.applicationKeyname)
    if !hasApplication {
        return nil, nil
    }

    return app.moduleNameWithTypeUsingFlag(remaining, KindApplication)
}

func (app *programAdapter) dataTypeDeclaration(input []byte) (Kind, []byte) {
    hasData, remaining := app.hasPrefix(input, app.dataKeyname)
    if !hasData {
        return nil, nil
    }

    return app.moduleNameWithTypeUsingFlag(remaining, KindData)
}

func (app *programAdapter) moduleNameWithTypeUsingFlag(input []byte, flag uint8) (Kind, []byte) {
    moduleName, typeName, remaining := app.moduleNameWithType(input)
    ins, err := app.kindBuilder.Create().WithFlag(flag).WithModule(moduleName).WithName(typeName).Now()
    if err != nil {
        return nil, nil
    }

    return ins, remaining
}

func (app *programAdapter) moduleNameWithType(input []byte) (string, string, []byte) {
    moduleName, remainingAfterModule := app.fetchName(input, app.moduleNameCharacters)
    if moduleName == "" {
        return "", "", nil
    }

    hasDelimiter, remainingAfterDelimiter := app.hasPrefix(remainingAfterModule, string(app.moduleTypeDelimiter))
    if !hasDelimiter {
        return "", "", nil
    }

    typeName, remainingAfterType := app.fetchName(remainingAfterDelimiter, app.typeCharacters)
    if typeName == "" {
        return "", "", nil
    }

    return moduleName, typeName, remainingAfterType
}

func (app *programAdapter) variableDeclaration(input []byte) (Variable, []byte) {
    moduleName, typeName, remaining := app.moduleNameWithType(input)
    variableName, remainingAfterVariable := app.fetchVariableNameUsage(remaining)
    if variableName == "" {
        return nil, nil
    }

    ins, err := app.variableBuilder.Create().WithModule(moduleName).WithKind(typeName).WithName(variableName).Now()
    if err != nil {
        return nil, nil
    }

    return ins, remainingAfterVariable
}

func (app *programAdapter) fetchVariableNameUsage(input []byte) (string, []byte) {
    hasPrefix, remainingAfterPrefix := app.hasPrefix(input, string(app.variableNameUsage))
    if !hasPrefix {
        return "", nil
    }

    return app.fetchName(remainingAfterPrefix, app.variableCharacters)
}

func (app *programAdapter) assignment(input []byte) (Assignment, []byte) {
    variable, remainingAfterVariable := app.variableDeclaration(input)
    if variable == nil {
        remainingAfterVariable = input
    }

    typeName := ""
    remaining := remainingAfterVariable
    if variable == nil {
        name, remainingAfterName := app.fetchVariableNameUsage(remaining)
        if name == "" {
            return nil, nil
        }

        typeName = name
        remaining = remainingAfterName
    }

    hasPrefix, remainingAfterPrefix := app.hasPrefix(remaining, string(app.assignmentDelimiter))
    if !hasPrefix {
        return nil, nil
    }

    content, remaining := app.fetchAssignmentContent(remainingAfterPrefix)
    if content == "" {
        return nil, nil
    }

    builder := app.assignmentBuilder.Create().WithContent(content)
    if variable != nil {
        builder.WithDeclaration(variable)
    }

    if typeName != "" {
        builder.WithName(typeName)
    }

    ins, err := builder.Now()
    if err != nil {
        return nil, nil
    }

    return ins, remaining
}

func (app *programAdapter) fetchAssignmentContent(input []byte) (string, []byte) {
    content := []byte{}
    skipNext := false
    for idx, oneByte := range(input) {
        content = append(content, oneByte)
        if skipNext {
            skipNext = false
            continue
        }

        if oneByte == app.escapeDelimiter {
            if len(input) -1 < (idx + 1) {
                continue
            }

            if input[idx + 1] == app.lineDelimiter {
                skipNext = true
                continue
            }
        }

        if oneByte == app.lineDelimiter {
            break
        }
    }

    if len(content) < 1 {
        return "", nil
    }

    found := content[:len(content) - 1]
    return string(found), input[len(found):]
}

func (app *programAdapter) action(input []byte) (Action, []byte) {
    attach, remainingAfterAttach :=app.actionAttach(input)
    if attach != nil {
        return attach, remainingAfterAttach
    }

    return app.actionDetach(input)
}

func (app *programAdapter) actionAttach(input []byte) (Action, []byte) {
    hasAttach, remainingAfterAttach := app.hasPrefix(input, app.attachKeyname)
    if !hasAttach {
        return nil, nil
    }

    scope, remainingAfterScope := app.scope(remainingAfterAttach)
    if scope == nil {
        return nil, nil
    }

    hasTo, remainingAfterTo := app.hasPrefix(remainingAfterScope, app.toKeyname)
    if !hasTo {
        return nil, nil
    }

    variableName, remainingAfterVariable := app.fetchName(remainingAfterTo, []byte(app.variableCharacters))
    if variableName == "" {
        return nil, nil
    }

    ins, err := app.actionBuilder.Create().IsAttach().WithApplication(variableName).WithScope(scope).Now()
    if err != nil {
        return nil, nil
    }

    return ins, remainingAfterVariable
}

func (app *programAdapter) actionDetach(input []byte) (Action, []byte) {
    hasDetach, remainingAfterDetach := app.hasPrefix(input, app.detachKeyname)
    if !hasDetach {
        return nil, nil
    }

    scope, remainingAfterScope := app.scope(remainingAfterDetach)
    if scope == nil {
        return nil, nil
    }

    hasFrom, remainingAfterFrom := app.hasPrefix(remainingAfterScope, app.fromKeyname)
    if !hasFrom {
        return nil, nil
    }

    variableName, remainingAfterVariable := app.fetchName(remainingAfterFrom, []byte(app.variableCharacters))
    if variableName == "" {
        return nil, nil
    }

    ins, err := app.actionBuilder.Create().WithApplication(variableName).WithScope(scope).Now()
    if err != nil {
        return nil, nil
    }

    return ins, remainingAfterVariable
}

func (app *programAdapter) scope(input []byte) (Scope, []byte) {
    programName, remainingAfterProgram := app.fetchName(input, []byte(app.variableCharacters))
    if programName == "" {
        return nil, nil
    }

    hasDelimiter, remainingAfterDelimiter := app.hasPrefix(remainingAfterProgram, string(app.scopeDelimiter))
    if !hasDelimiter {
        return nil, nil
    }

    moduleName, remainingAfterModule := app.fetchName(remainingAfterDelimiter, []byte(app.variableCharacters))
    if moduleName == "" {
        return nil, nil
    }

    ins, err := app.scopeBuilder.Create().WithModule(moduleName).WithProgram(programName).Now()
    if err != nil {
        return nil, nil
    }

    return ins, remainingAfterModule
}

func (app *programAdapter) execution(input []byte) (Execution, []byte) {
    variable, remainingAfterVariable := app.variableDeclaration(input)
    if variable == nil {
        remainingAfterVariable = input
    }

    remaining := remainingAfterVariable
    if variable != nil {
        hasEqual, remainingAfterEqual := app.hasPrefix(remainingAfterVariable, string(app.assignmentDelimiter))
        if !hasEqual {
            return nil, nil
        }

        remaining = remainingAfterEqual
    }

    hasExecute, remainingAfterExecute := app.hasPrefix(remaining, app.executeKeyname)
    if !hasExecute {
        return nil, nil
    }

    variableName, remainingAfterVariable := app.fetchName(remainingAfterExecute, app.variableCharacters)
    if variableName == "" {
        return nil, nil
    }

    builder := app.executionBuilder.Create().WithApplication(variableName)
    if variable != nil {
        builder.WithDeclaration(variable)
    }

    ins, err := builder.Now()
    if err != nil {
        return nil, nil
    }

    return ins, remainingAfterVariable
}

func (app *programAdapter) hasPrefix(input []byte, prefix string) (bool, []byte) {
    retInput := app.removeChannelCharactersPrefix(input)
    if !bytes.HasPrefix(retInput, []byte(prefix)) {
        return false, retInput
    }

    length := len(prefix)
    return true, retInput[length:]
}

func (app *programAdapter) fetchName(input []byte, characters []byte) (string, []byte) {
	nameBytes := []byte{}
    retInput := app.removeChannelCharactersPrefix(input)
	for _, oneInputByte := range retInput {
		if !app.isBytePresent(oneInputByte, characters) {
			break
		}

		nameBytes = append(nameBytes, oneInputByte)
	}

	if len(nameBytes) <= 0 {
		return "", nil
	}

	return string(nameBytes), retInput[len(nameBytes):]
}

func (app *programAdapter) isBytePresent(value byte, data []byte) bool {
	isPresent := false
	for _, oneChanByte := range data {
		if value == oneChanByte {
			isPresent = true
			break
		}
	}

	return isPresent
}

func (app *programAdapter) removeChannelCharactersPrefix(input []byte) []byte {
	cpt := 0
	for _, oneInputByte := range input {
		if app.isBytePresent(oneInputByte, app.channelCharacters) {
			cpt++
            continue
		}

		break
	}

	return input[cpt:]
}
