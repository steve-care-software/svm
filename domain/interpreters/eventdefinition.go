package interpreters

type eventDefinition struct {
    execute ExecuteFn
    assign AssignFn
}

func createEventDefinition(
    execute ExecuteFn,
) EventDefinition {
    return createEventDefinitionInternally(execute, nil)
}

func createEventDefinitionWithAssign(
    execute ExecuteFn,
    assign AssignFn,
) EventDefinition {
    return createEventDefinitionInternally(execute, assign)
}

func createEventDefinitionInternally(
    execute ExecuteFn,
    assign AssignFn,
) EventDefinition {
    out := eventDefinition{
        execute: execute,
        assign: assign,
    }

    return &out
}

// Execute returns the execute func
func (obj *eventDefinition) Execute() ExecuteFn {
    return obj.execute
}

// HasAssign returns true if there is an assign, false otherwise
func (obj *eventDefinition) HasAssign() bool {
    return obj.assign != nil
}

// Assign returns the assign func
func (obj *eventDefinition) Assign() AssignFn {
    return obj.assign
}
