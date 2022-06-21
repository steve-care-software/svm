package lexers

type execution struct {
	application string
	assignee    Assignee
}

func createExecution(
	application string,
) Execution {
	return createExecutionInternally(application, nil)
}

func createExecutionWithAssignee(
	application string,
	assignee Assignee,
) Execution {
	return createExecutionInternally(application, assignee)
}

func createExecutionInternally(
	application string,
	assignee Assignee,
) Execution {
	out := execution{
		application: application,
		assignee:    assignee,
	}

	return &out
}

// Application returns the application
func (obj *execution) Application() string {
	return obj.application
}

// HasAssignee returns true if there is a assignee, false otherwise
func (obj *execution) HasAssignee() bool {
	return obj.assignee != nil
}

// Assignee returns the assignee, if any
func (obj *execution) Assignee() Assignee {
	return obj.assignee
}
