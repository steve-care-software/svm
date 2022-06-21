package interpreters

type watch struct {
	module string
	event  WatchEvent
}

func createWatch(
	module string,
	event WatchEvent,
) Watch {
	out := watch{
		module: "",
		event:  nil,
	}

	return &out
}

// Module returns the module
func (obj *watch) Module() string {
	return obj.module
}

// Event returns the event
func (obj *watch) Event() WatchEvent {
	return obj.event
}
