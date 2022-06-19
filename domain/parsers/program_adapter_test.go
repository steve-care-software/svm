package parsers

import (
	"testing"
    "github.com/steve-care-software/svm/domain/lexers"
)

func TestAdapter_isSuccess(t *testing.T) {
	script := `
        // declare the module and its types;
        module data;
        type data data.byte;
        type data data.bytes;
        type application data.leftShift;

        // declare the parameters;
		-> data.bytes $input;
		<- data.bytes $output;

        // declare the application;
        data.leftShift $leftShiftApp;

        // declare the amount to shift;
        data.byte $amount = $34;

        // atachthe data;
        attach input:data @ leftShiftApp;

        // attach the amount;
        attach amount:amount @ leftShiftApp;

        // execute;
        $output = execute leftShiftApp;
	`

    lexedProgram, _, err := lexers.NewProgramAdapter().ToProgram(script)
    if err != nil {
        t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
        return
    }

    program, err := NewProgramAdapter(createScreenWriterForTests()).ToProgram(lexedProgram)
    if err != nil {
        t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
        return
    }

	if !program.HasParameters() {
		t.Errorf("the program was expecting parameters")
		return
	}

	parameters := program.Parameters().List()
	if len(parameters) != 2 {
		t.Errorf("%d parameters were expected, %d returned", 2, len(parameters))
		return
	}

	executions := program.Executions().List()
	if len(executions) != 1 {
		t.Errorf("%d executions were expected, %d returned", 1, len(executions))
		return
	}

}


func TestAdapter_withoutCommentLogger_isSuccess(t *testing.T) {
	script := `
        // declare the module and its types;
        module data;
        type data data.byte;
        type data data.bytes;
        type application data.leftShift;

        // declare the parameters;
		-> data.bytes $input;
		<- data.bytes $output;

        // declare the application;
        data.leftShift $leftShiftApp;

        // declare the amount to shift;
        data.byte $amount = $34;

        // atachthe data;
        attach input:data @ leftShiftApp;

        // attach the amount;
        attach amount:amount @ leftShiftApp;

        // execute;
        $output = execute leftShiftApp;
	`

    lexedProgram, _, err := lexers.NewProgramAdapter().ToProgram(script)
    if err != nil {
        t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
        return
    }

    program, err := NewProgramAdapter(nil).ToProgram(lexedProgram)
    if err != nil {
        t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
        return
    }

	if !program.HasParameters() {
		t.Errorf("the program was expecting parameters")
		return
	}

	parameters := program.Parameters().List()
	if len(parameters) != 2 {
		t.Errorf("%d parameters were expected, %d returned", 2, len(parameters))
		return
	}

	executions := program.Executions().List()
	if len(executions) != 1 {
		t.Errorf("%d executions were expected, %d returned", 1, len(executions))
		return
	}

}
