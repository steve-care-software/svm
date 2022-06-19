package interpreters

type event struct {
    enter ExecuteFn
    exit ExecuteFn
}

func createEventWithEnter(
    enter ExecuteFn,
) Event {
    return createEventInternally(enter, nil)
}

func createEventWithExit(
    exit ExecuteFn,
) Event {
    return createEventInternally(nil, exit)
}

func createEventWithEnterAndExit(
    enter ExecuteFn,
    exit ExecuteFn,
) Event {
    return createEventInternally(enter, exit)
}

func createEventInternally(
    enter ExecuteFn,
    exit ExecuteFn,
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
func (obj *event) Enter() ExecuteFn {
    return obj.enter
}

// HasExit returns true if there is an exit, false otherwise
func (obj *event) HasExit() bool {
    return obj.exit != nil
}

// Exit returns the exit, if any
func (obj *event) Exit() ExecuteFn {
    return obj.exit
}
