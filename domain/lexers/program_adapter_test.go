package lexers

import (
	"fmt"
	"testing"
)

func TestAdapter_withComment_withoutRemaining_withInstructionLineDelimiter_withParameterLineDelimiter_isSuccess(t *testing.T) {
	script := `
		// this is a comment;;
	`

	program, remaining, err := NewProgramAdapter().ToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsParameter() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

	if !ins.IsComment() {
		t.Errorf("the instruction was expected to be a coment")
		return
	}

	str := ins.Comment()
	if str != " this is a comment" {
		t.Errorf("the comment was expected to be '%s', '%s' returned", " this is a comment", str)
		return
	}

	if ins.IsModule() {
		t.Errorf("the instruction was expected to NOT be a module")
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

func TestAdapter_withModule_withParameters_withoutRemaining_withInstructionLineDelimiter_withParameterLineDelimiter_isSuccess(t *testing.T) {
	script := `
		-> myModule.myType $input;;
		<- myModule.myType $output;;
		module myModule;;
	`

	program, remaining, err := NewProgramAdapter().ToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 3 {
		t.Errorf("the program was expecting %d instruction, %d returned", 3, len(instructions))
		return
	}

	firstInstruction := instructions[0]
	if !firstInstruction.IsParameter() {
		t.Errorf("the first instruction was expected to be a parameter")
		return
	}

	firstParameter := instructions[0].Parameter()
	if !firstParameter.IsInput() {
		t.Errorf("the parameter (index: 0) was expected to be an input")
		return
	}

	if !firstParameter.IsInput() {
		t.Errorf("the parameter (index: 0) was expected to be an input")
		return
	}

	secondInstruction := instructions[1]
	if !secondInstruction.IsParameter() {
		t.Errorf("the second instruction was expected to be a parameter")
		return
	}

	secondParameter := instructions[1].Parameter()
	if secondParameter.IsInput() {
		t.Errorf("the parameter (index: 1) was NOT expected to be an input")
		return
	}

	ins := instructions[2]
	if ins.IsParameter() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

	if ins.IsComment() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

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
	script := "module myModule;;"
	program, remaining, err := NewProgramAdapter().ToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsParameter() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

	if ins.IsComment() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

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
	script := "myModule.myType $myVariable;;"
	program, remaining, err := NewProgramAdapter().ToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsParameter() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

	if ins.IsComment() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

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
	script := "type data myModule.myDataType;;"
	program, remaining, err := NewProgramAdapter().ToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsParameter() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

	if ins.IsComment() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

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
	if retFlag&KindData == 0 {
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
	script := "type application myModule.myApplicationType;;"
	program, remaining, err := NewProgramAdapter().ToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsParameter() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

	if ins.IsComment() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

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
	if retFlag&KindApplication == 0 {
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
	expectedContent := " this is an escaped: \\;; and other characters"
	script := fmt.Sprintf("myModule.myType $myVariable =%s;;", expectedContent)
	program, remaining, err := NewProgramAdapter().ToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsParameter() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

	if ins.IsComment() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

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

	assignee := assignment.Assignee()
	if !assignee.IsDeclaration() {
		t.Errorf("the assignee was expected to contain a variable declaration")
		return
	}

	if assignee.IsName() {
		t.Errorf("the assignee was expected to NOT contain a name")
		return
	}

	variable := assignee.Declaration()
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
	expectedContent := " this is an escaped: \\;; and other characters"
	script := fmt.Sprintf("$myVariable =%s;;", expectedContent)
	program, remaining, err := NewProgramAdapter().ToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsParameter() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

	if ins.IsComment() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

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

	assignee := assignment.Assignee()
	if assignee.IsDeclaration() {
		t.Errorf("the assignee was expected to NOT contain a variable declaration")
		return
	}

	if !assignee.IsName() {
		t.Errorf("the assignee was expected to contain a name")
		return
	}

	retName := assignee.Name()
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
	script := "attach myDataVariable:data @ myAppVariable;;"
	program, remaining, err := NewProgramAdapter().ToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsParameter() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

	if ins.IsComment() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

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
	script := "detach myDataVariable:data @ myAppVariable;;"
	program, remaining, err := NewProgramAdapter().ToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsParameter() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

	if ins.IsComment() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

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
	script := "execute myAppVariable;;"
	program, remaining, err := NewProgramAdapter().ToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsParameter() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

	if ins.IsComment() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

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
	if execution.HasAssignee() {
		t.Errorf("the execution was expected to NOT contain an assignee")
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
	script := "myModule.myDataType $myOutput = execute myAppVariable;;"
	program, remaining, err := NewProgramAdapter().ToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsParameter() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

	if ins.IsComment() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

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
	if !execution.HasAssignee() {
		t.Errorf("the execution was expected to contain an assignee")
		return
	}

	assignee := execution.Assignee()
	if assignee.IsName() {
		t.Errorf("the assignee was expected to NOT be a name")
		return
	}

	if !assignee.IsDeclaration() {
		t.Errorf("the assignee was expected to be a declaration")
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

func TestAdapter_withExecution_withAssignee_withName_withoutParameters_withoutRemaining_withLineDelimiter_isSuccess(t *testing.T) {
	script := "$myOutput = execute myAppVariable;;"
	program, remaining, err := NewProgramAdapter().ToProgram(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	instructions := program.Instructions()
	if len(instructions) != 1 {
		t.Errorf("the program was expecting %d instruction, %d returned", 1, len(instructions))
		return
	}

	ins := instructions[0]
	if ins.IsParameter() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

	if ins.IsComment() {
		t.Errorf("the instruction was expected to NOT be a coment")
		return
	}

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
	if !execution.HasAssignee() {
		t.Errorf("the execution was expected to contain an assignee")
		return
	}

	assignee := execution.Assignee()
	if !assignee.IsName() {
		t.Errorf("the assignee was expected to be a name")
		return
	}

	if assignee.Name() != "myOutput" {
		t.Errorf("the variable declaration name was expected to be %s, %s returned", "myOutput", assignee.Name())
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
