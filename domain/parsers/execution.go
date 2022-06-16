package parsers

type execution struct {
    module string
    application Variable
    attachments Variables
}

func createExecution(
    module string,
    application Variable,
) Execution {
    return createExecutionInternally(module, application, nil)
}

func createExecutionWithAttachments(
    module string,
    application Variable,
    attachments Variables,
) Execution {
    return createExecutionInternally(module, application, attachments)
}

func createExecutionInternally(
    module string,
    application Variable,
    attachments Variables,
) Execution {
    out := execution{
        module: module,
        application: application,
        attachments: attachments,
    }

    return &out
}

// Module returns the module
func (obj *execution) Module() string {
    return obj.module
}

// Application returns the application
func (obj *execution) Application() Variable {
    return obj.application
}

// HasAttachments returns true if there is attachments, false otherwise
func (obj *execution) HasAttachments() bool {
    return obj.attachments != nil
}

// Attachments returns the attachments, if any
func (obj *execution) Attachments() Variables {
    return obj.attachments
}
