package parsers

type execution struct {
	application Application
	output      Variable
}

func createExecution(
	application Application,
) Execution {
	return createExecutionInternally(application, nil)
}

func createExecutionWithOutput(
	application Application,
	output Variable,
) Execution {
	return createExecutionInternally(application, output)
}

func createExecutionInternally(
	application Application,
	output Variable,
) Execution {
	out := execution{
		application: application,
		output:      output,
	}

	return &out
}

// Application returns the application
func (obj *execution) Application() Application {
	return obj.application
}

// HasOutput returns true if there is an output, false otherwise
func (obj *execution) HasOutput() bool {
	return obj.output != nil
}

// Output returns the output, if any
func (obj *execution) Output() Variable {
	return obj.output
}
