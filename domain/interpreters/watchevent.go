package interpreters

type watchEvent struct {
    enter EnterWatchFn
    exit ExitWatchFn
}

func createWatchEventWithEnter(
    enter EnterWatchFn,
) WatchEvent {
    return createWatchEventInternally(enter, nil)
}

func createWatchEventWithExit(
    exit ExitWatchFn,
) WatchEvent {
    return createWatchEventInternally(nil, exit)
}

func createWatchEventWithEnterAndExit(
    enter EnterWatchFn,
    exit ExitWatchFn,
) WatchEvent {
    return createWatchEventInternally(enter, exit)
}

func createWatchEventInternally(
    enter EnterWatchFn,
    exit ExitWatchFn,
) WatchEvent {
    out := watchEvent{
        enter: enter,
        exit: exit,
    }

    return &out
}

// HasEnter returns true if there is an enter, false otherwise
func (obj *watchEvent) HasEnter() bool {
    return obj.enter != nil
}

// Enter returns the enter, if any
func (obj *watchEvent) Enter() EnterWatchFn {
    return obj.enter
}

// HasExit returns true if there is an exit, false otherwise
func (obj *watchEvent) HasExit() bool {
    return obj.exit != nil
}

// Exit returns the exit, if any
func (obj *watchEvent) Exit() ExitWatchFn {
    return obj.exit
}
