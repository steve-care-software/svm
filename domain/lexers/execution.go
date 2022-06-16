package lexers

type execution struct {
    application string
    declaration Variable
}

func createExecution(
    application string,
) Execution {
    return createExecutionInternally(application, nil)
}

func createExecutionWithDeclaration(
    application string,
    declaration Variable,
) Execution {
    return createExecutionInternally(application, declaration)
}

func createExecutionInternally(
    application string,
    declaration Variable,
) Execution {
    out := execution{
        application: application,
        declaration: declaration,
    }

    return &out
}

// Application returns the application
func (obj *execution) Application() string {
    return obj.application
}

// HasDeclaration returns true if there is a declaration, false otherwise
func (obj *execution) HasDeclaration() bool {
    return obj.declaration != nil
}

// Declaration returns the declaration, if any
func (obj *execution) Declaration() Variable {
    return obj.declaration
}
