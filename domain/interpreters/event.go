package interpreters

type event struct {
    enter EventDefinition
    exit EventDefinition
}

func createEventWithEnter(
    enter EventDefinition,
) Event {
    return createEventInternally(enter, nil)
}

func createEventWithExit(
    exit EventDefinition,
) Event {
    return createEventInternally(nil, exit)
}

func createEventWithEnterAndExit(
    enter EventDefinition,
    exit EventDefinition,
) Event {
    return createEventInternally(enter, exit)
}

func createEventInternally(
    enter EventDefinition,
    exit EventDefinition,
) Event {
    out := event{
        enter: enter,
        exit: exit,
    }

    return &out
}

// HasEnter returns true if there is an enter, false otherwise
func (obj *event) HasEnter() bool {
    return obj.enter != nil
}

// Enter returns the enter, if any
func (obj *event) Enter() EventDefinition {
    return obj.enter
}

// HasExit returns true if there is an exit, false otherwise
func (obj *event) HasExit() bool {
    return obj.exit != nil
}

// Exit returns the exit, if any
func (obj *event) Exit() EventDefinition {
    return obj.exit
}
