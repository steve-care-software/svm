package interpreters

type watch struct {
    module string
    event Event
}

func createWatch(
    module string,
    event Event,
) Watch {
    out:= watch{
        module: "",
        event: nil,
    }

    return &out
}

// Module returns the module
func (obj *watch) Module() string {
    return obj.module
}

// Event returns the event
func (obj *watch) Event() Event {
    return obj.event
}
