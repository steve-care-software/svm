package lexers

import (
	"testing"
	"fmt"
)

func TestAdapter_withModule_withParameters_withoutRemaining_withInstructionLineDelimiter_withParameterLineDelimiter_isSuccess(t *testing.T) {
	script := `
		-> myModule.myType $input;
		<- myModule.myType $output;
		module myModule;
	`

	program, remaining, err := NewProgramAdapter().ScriptToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !program.HasParameters() {
		t.Errorf("the program was expecting parameters")
		return
	}

	parameters := program.Parameters()
	if len(parameters) != 2 {
		t.Errorf("%d parameters were expected, %d returned", 2, len(parameters))
		return
	}

	if !parameters[0].IsInput() {
		t.Errorf("the parameter (index: 0) was expected to be an input")
		return
	}

	if parameters[1].IsInput() {
		t.Errorf("the parameter (index: 1) was NOT expected to be an input")
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if !ins.IsModule() {
		t.Errorf("the instruction was expected to be a module")
		return
	}

	retModule := ins.Module()
	if retModule != "myModule" {
		t.Errorf("the module name was expected to be '%s', '%s' returned", "myModule", retModule)
		return
	}

	if ins.IsKind() {
		t.Errorf("the instruction was NOT expected to be a type declaration")
		return
	}

	if ins.IsVariable() {
		t.Errorf("the instruction was NOT expected to be a variable declaration")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the instruction was NOT expected to be an assignment")
		return
	}

	if ins.IsAction() {
		t.Errorf("the instruction was NOT expected to be an action")
		return
	}

	if ins.IsExecution() {
		t.Errorf("the instruction was NOT expected to be an execution")
		return
	}

	if len(remaining) > 0 {
		t.Errorf("the remaining was expected to be empty, %v returned", remaining)
		return
	}
}

func TestAdapter_withModule_withoutParameters_withoutRemaining_withLineDelimiter_isSuccess(t *testing.T) {
	script := "module myModule;"
	program, remaining, err := NewProgramAdapter().ScriptToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if program.HasParameters() {
		t.Errorf("the program was NOT expecting parameters")
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if !ins.IsModule() {
		t.Errorf("the instruction was expected to be a module")
		return
	}

	retModule := ins.Module()
	if retModule != "myModule" {
		t.Errorf("the module name was expected to be '%s', '%s' returned", "myModule", retModule)
		return
	}

	if ins.IsKind() {
		t.Errorf("the instruction was NOT expected to be a type declaration")
		return
	}

	if ins.IsVariable() {
		t.Errorf("the instruction was NOT expected to be a variable declaration")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the instruction was NOT expected to be an assignment")
		return
	}

	if ins.IsAction() {
		t.Errorf("the instruction was NOT expected to be an action")
		return
	}

	if ins.IsExecution() {
		t.Errorf("the instruction was NOT expected to be an execution")
		return
	}

	if len(remaining) > 0 {
		t.Errorf("the remaining was expected to be empty, %v returned", remaining)
		return
	}
}

func TestAdapter_withVariableDeclaration_withoutRemaining_withLineDelimiter_isSuccess(t *testing.T) {
	script := "myModule.myType $myVariable;"
	program, remaining, err := NewProgramAdapter().ScriptToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if program.HasParameters() {
		t.Errorf("the program was NOT expecting parameters")
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsModule() {
		t.Errorf("the instruction was expected to NOT be a module")
		return
	}

	if ins.IsKind() {
		t.Errorf("the instruction was NOT expected to be a type declaration")
		return
	}

	if !ins.IsVariable() {
		t.Errorf("the instruction was expected to be a variable declaration")
		return
	}

	variable := ins.Variable()
	retModule := variable.Module()
	if retModule != "myModule" {
		t.Errorf("the module was expected to be '%s', '%s' returned", "myModule", retModule)
		return
	}

	retKind := variable.Kind()
	if retKind != "myType" {
		t.Errorf("the type was expected to be '%s', '%s' returned", "myType", retKind)
		return
	}

	retName := variable.Name()
	if retName != "myVariable" {
		t.Errorf("the variable name was expected to be '%s', '%s' returned", "myVariable", retName)
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the instruction was NOT expected to be an assignment")
		return
	}

	if ins.IsAction() {
		t.Errorf("the instruction was NOT expected to be an action")
		return
	}

	if ins.IsExecution() {
		t.Errorf("the instruction was NOT expected to be an execution")
		return
	}

	if len(remaining) > 0 {
		t.Errorf("the remaining was expected to be empty, %v returned", remaining)
		return
	}
}

func TestAdapter_withTypeDeclaration_isData_withoutParameters_withoutRemaining_withLineDelimiter_isSuccess(t *testing.T) {
	script := "type data myModule.myDataType;"
	program, remaining, err := NewProgramAdapter().ScriptToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if program.HasParameters() {
		t.Errorf("the program was NOT expecting parameters")
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsModule() {
		t.Errorf("the instruction was NOT expected to be a module")
		return
	}

	if !ins.IsKind() {
		t.Errorf("the instruction was expected to be a type declaration")
		return
	}

	kind := ins.Kind()
	retModule := kind.Module()
	if retModule != "myModule" {
		t.Errorf("the module name was expected to be '%s', '%s' returned", "myModule", retModule)
		return
	}

	retName := kind.Name()
	if retName != "myDataType" {
		t.Errorf("the type name was expected to be '%s', '%s' returned", "myDataType", retName)
		return
	}

	retFlag := kind.Flag()
	if retFlag & KindData == 0 {
		t.Errorf("the flag (%d) was expected to contain %d (data)", retFlag, KindData)
		return
	}

	if ins.IsVariable() {
		t.Errorf("the instruction was NOT expected to be a variable declaration")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the instruction was NOT expected to be an assignment")
		return
	}

	if ins.IsAction() {
		t.Errorf("the instruction was NOT expected to be an action")
		return
	}

	if ins.IsExecution() {
		t.Errorf("the instruction was NOT expected to be an execution")
		return
	}

	if len(remaining) > 0 {
		t.Errorf("the remaining was expected to be empty, %v returned", remaining)
		return
	}
}


func TestAdapter_withTypeDeclaration_isApplication_withoutParameters_withoutRemaining_withLineDelimiter_isSuccess(t *testing.T) {
	script := "type application myModule.myApplicationType;"
	program, remaining, err := NewProgramAdapter().ScriptToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if program.HasParameters() {
		t.Errorf("the program was NOT expecting parameters")
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsModule() {
		t.Errorf("the instruction was NOT expected to be a module")
		return
	}

	if !ins.IsKind() {
		t.Errorf("the instruction was expected to be a type declaration")
		return
	}

	kind := ins.Kind()
	retModule := kind.Module()
	if retModule != "myModule" {
		t.Errorf("the module name was expected to be '%s', '%s' returned", "myModule", retModule)
		return
	}

	retName := kind.Name()
	if retName != "myApplicationType" {
		t.Errorf("the type name was expected to be '%s', '%s' returned", "myApplicationType", retName)
		return
	}

	retFlag := kind.Flag()
	if retFlag & KindApplication == 0 {
		t.Errorf("the flag (%d) was expected to contain %d (data)", retFlag, KindApplication)
		return
	}

	if ins.IsVariable() {
		t.Errorf("the instruction was NOT expected to be a variable declaration")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the instruction was NOT expected to be an assignment")
		return
	}

	if ins.IsAction() {
		t.Errorf("the instruction was NOT expected to be an action")
		return
	}

	if ins.IsExecution() {
		t.Errorf("the instruction was NOT expected to be an execution")
		return
	}

	if len(remaining) > 0 {
		t.Errorf("the remaining was expected to be empty, %v returned", remaining)
		return
	}
}

func TestAdapter_withAssignment_withVariableDeclaration_withoutParameters_withoutRemaining_withLineDelimiter_isSuccess(t *testing.T) {
	expectedContent := " this is an escaped: \\; and other characters"
	script := fmt.Sprintf("myModule.myType $myVariable =%s;", expectedContent)
	program, remaining, err := NewProgramAdapter().ScriptToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if program.HasParameters() {
		t.Errorf("the program was NOT expecting parameters")
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsModule() {
		t.Errorf("the instruction was NOT expected to be a module")
		return
	}

	if ins.IsKind() {
		t.Errorf("the instruction was expected to NOT be a type declaration")
		return
	}

	if ins.IsVariable() {
		t.Errorf("the instruction was expected to NOT be a variable declaration")
		return
	}

	if !ins.IsAssignment() {
		t.Errorf("the instruction was expected to be an assignment")
		return
	}

	assignment := ins.Assignment()
	retContent := assignment.Content()
	if retContent != expectedContent {
		t.Errorf("the content was expected to be '%s', '%s' returned", expectedContent, retContent)
		return
	}

	if !assignment.IsDeclaration() {
		t.Errorf("the assignment was expected to contain a variable declaration")
		return
	}

	if assignment.IsName() {
		t.Errorf("the assignment was expected to NOT contain a name")
		return
	}

	variable := assignment.Declaration()
	retModule := variable.Module()
	if retModule != "myModule" {
		t.Errorf("the module was expected to be '%s', '%s' returned", "myModule", retModule)
		return
	}

	retKind := variable.Kind()
	if retKind != "myType" {
		t.Errorf("the type was expected to be '%s', '%s' returned", "myType", retKind)
		return
	}

	retName := variable.Name()
	if retName != "myVariable" {
		t.Errorf("the variable name was expected to be '%s', '%s' returned", "myVariable", retName)
		return
	}

	if ins.IsAction() {
		t.Errorf("the instruction was NOT expected to be an action")
		return
	}

	if ins.IsExecution() {
		t.Errorf("the instruction was NOT expected to be an execution")
		return
	}

	if len(remaining) > 0 {
		t.Errorf("the remaining was expected to be empty, %v returned", remaining)
		return
	}
}


func TestAdapter_withAssignment_withName_withoutParameters_withoutRemaining_withLineDelimiter_isSuccess(t *testing.T) {
	expectedContent := " this is an escaped: \\; and other characters"
	script := fmt.Sprintf("$myVariable =%s;", expectedContent)
	program, remaining, err := NewProgramAdapter().ScriptToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if program.HasParameters() {
		t.Errorf("the program was NOT expecting parameters")
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsModule() {
		t.Errorf("the instruction was NOT expected to be a module")
		return
	}

	if ins.IsKind() {
		t.Errorf("the instruction was expected to NOT be a type declaration")
		return
	}

	if ins.IsVariable() {
		t.Errorf("the instruction was expected to NOT be a variable declaration")
		return
	}

	if !ins.IsAssignment() {
		t.Errorf("the instruction was expected to be an assignment")
		return
	}

	assignment := ins.Assignment()
	retContent := assignment.Content()
	if retContent != expectedContent {
		t.Errorf("the content was expected to be '%s', '%s' returned", expectedContent, retContent)
		return
	}

	if assignment.IsDeclaration() {
		t.Errorf("the assignment was expected to NOT contain a variable declaration")
		return
	}

	if !assignment.IsName() {
		t.Errorf("the assignment was expected to contain a name")
		return
	}

	retName := assignment.Name()
	if retName != "myVariable" {
		t.Errorf("the variable name was expected to be '%s', '%s' returned", "myVariable", retName)
		return
	}

	if ins.IsAction() {
		t.Errorf("the instruction was NOT expected to be an action")
		return
	}

	if ins.IsExecution() {
		t.Errorf("the instruction was NOT expected to be an execution")
		return
	}

	if len(remaining) > 0 {
		t.Errorf("the remaining was expected to be empty, %v returned", remaining)
		return
	}
}


func TestAdapter_withAction_isAttach_withoutParameters_withoutRemaining_withLineDelimiter_isSuccess(t *testing.T) {
	script := "attach myDataVariable:data @ myAppVariable;"
	program, remaining, err := NewProgramAdapter().ScriptToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if program.HasParameters() {
		t.Errorf("the program was NOT expecting parameters")
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsModule() {
		t.Errorf("the instruction was NOT expected to be a module")
		return
	}

	if ins.IsKind() {
		t.Errorf("the instruction was expected to NOT be a type declaration")
		return
	}

	if ins.IsVariable() {
		t.Errorf("the instruction was expected to NOT be a variable declaration")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the instruction was expected to NOT be an assignment")
		return
	}

	if !ins.IsAction() {
		t.Errorf("the instruction was expected to be an action")
		return
	}

	action := ins.Action()
	reApplication := action.Application()
	if reApplication != "myAppVariable" {
		t.Errorf("the application variable was expected to be %s, %s returned", "myAppVariable", reApplication)
		return
	}

	if !action.IsAttach() {
		t.Errorf("the action was expected to be attach")
		return
	}

	scope := action.Scope()
	retModule := scope.Module()
	if retModule != "data" {
		t.Errorf("the scoped module variable was expected to be %s, %s returned", "data", retModule)
		return
	}

	retProgram := scope.Program()
	if retProgram != "myDataVariable" {
		t.Errorf("the scoped program variable was expected to be %s, %s returned", "myDataVariable", retProgram)
		return
	}

	if ins.IsExecution() {
		t.Errorf("the instruction was NOT expected to be an execution")
		return
	}

	if len(remaining) > 0 {
		t.Errorf("the remaining was expected to be empty, %v returned", remaining)
		return
	}
}


func TestAdapter_withAction_isDetach_withoutParameters_withoutRemaining_withLineDelimiter_isSuccess(t *testing.T) {
	script := "detach myDataVariable:data @ myAppVariable;"
	program, remaining, err := NewProgramAdapter().ScriptToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if program.HasParameters() {
		t.Errorf("the program was NOT expecting parameters")
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsModule() {
		t.Errorf("the instruction was NOT expected to be a module")
		return
	}

	if ins.IsKind() {
		t.Errorf("the instruction was expected to NOT be a type declaration")
		return
	}

	if ins.IsVariable() {
		t.Errorf("the instruction was expected to NOT be a variable declaration")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the instruction was expected to NOT be an assignment")
		return
	}

	if !ins.IsAction() {
		t.Errorf("the instruction was expected to be an action")
		return
	}

	action := ins.Action()
	reApplication := action.Application()
	if reApplication != "myAppVariable" {
		t.Errorf("the application variable was expected to be %s, %s returned", "myAppVariable", reApplication)
		return
	}

	if action.IsAttach() {
		t.Errorf("the action was expected to NOT be attach")
		return
	}

	scope := action.Scope()
	retModule := scope.Module()
	if retModule != "data" {
		t.Errorf("the scoped module variable was expected to be %s, %s returned", "data", retModule)
		return
	}

	retProgram := scope.Program()
	if retProgram != "myDataVariable" {
		t.Errorf("the scoped program variable was expected to be %s, %s returned", "myDataVariable", retProgram)
		return
	}

	if ins.IsExecution() {
		t.Errorf("the instruction was NOT expected to be an execution")
		return
	}

	if len(remaining) > 0 {
		t.Errorf("the remaining was expected to be empty, %v returned", remaining)
		return
	}
}


func TestAdapter_withExecution_withoutDeclaration_withoutParameters_withoutRemaining_withLineDelimiter_isSuccess(t *testing.T) {
	script := "execute myAppVariable;"
	program, remaining, err := NewProgramAdapter().ScriptToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if program.HasParameters() {
		t.Errorf("the program was NOT expecting parameters")
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsModule() {
		t.Errorf("the instruction was NOT expected to be a module")
		return
	}

	if ins.IsKind() {
		t.Errorf("the instruction was expected to NOT be a type declaration")
		return
	}

	if ins.IsVariable() {
		t.Errorf("the instruction was expected to NOT be a variable declaration")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the instruction was expected to NOT be an assignment")
		return
	}

	if ins.IsAction() {
		t.Errorf("the instruction was expected to NOT be an action")
		return
	}

	if !ins.IsExecution() {
		t.Errorf("the instruction was expected to be an execution")
		return
	}

	execution := ins.Execution()
	if execution.HasDeclaration() {
		t.Errorf("the execution was expected to NOT contain a variable declaration")
		return
	}

	retApplication := execution.Application()
	if retApplication != "myAppVariable" {
		t.Errorf("the application was expected to be %s, %s returned", "myAppVariable", retApplication)
		return
	}

	if len(remaining) > 0 {
		t.Errorf("the remaining was expected to be empty, %v returned", remaining)
		return
	}
}

func TestAdapter_withExecution_withDeclaration_withoutParameters_withoutRemaining_withLineDelimiter_isSuccess(t *testing.T) {
	script := "myModule.myDataType $myOutput = execute myAppVariable;"
	program, remaining, err := NewProgramAdapter().ScriptToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if program.HasParameters() {
		t.Errorf("the program was NOT expecting parameters")
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsModule() {
		t.Errorf("the instruction was NOT expected to be a module")
		return
	}

	if ins.IsKind() {
		t.Errorf("the instruction was expected to NOT be a type declaration")
		return
	}

	if ins.IsVariable() {
		t.Errorf("the instruction was expected to NOT be a variable declaration")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the instruction was expected to NOT be an assignment")
		return
	}

	if ins.IsAction() {
		t.Errorf("the instruction was expected to NOT be an action")
		return
	}

	if !ins.IsExecution() {
		t.Errorf("the instruction was expected to be an execution")
		return
	}

	execution := ins.Execution()
	if !execution.HasDeclaration() {
		t.Errorf("the execution was expected to contain a variable declaration")
		return
	}

	declaration := execution.Declaration()
	if declaration.Name() != "myOutput" {
		t.Errorf("the variable declaration name was expected to be %s, %s returned", "myOutput", declaration.Name())
		return
	}

	retApplication := execution.Application()
	if retApplication != "myAppVariable" {
		t.Errorf("the application was expected to be %s, %s returned", "myAppVariable", retApplication)
		return
	}

	if len(remaining) > 0 {
		t.Errorf("the remaining was expected to be empty, %v returned", remaining)
		return
	}
}
