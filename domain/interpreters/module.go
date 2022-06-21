package interpreters

type module struct {
	name    string
	event   ExecuteFn
	watches Watches
}

func createModuleWithEvent(
	name string,
	event ExecuteFn,
) Module {
	return createModuleInternally(name, event, nil)
}

func createModuleWithWatches(
	name string,
	watches Watches,
) Module {
	return createModuleInternally(name, nil, watches)
}

func createModuleWithEventAndWatches(
	name string,
	event ExecuteFn,
	watches Watches,
) Module {
	return createModuleInternally(name, event, watches)
}

func createModuleInternally(
	name string,
	event ExecuteFn,
	watches Watches,
) Module {
	out := module{
		name:    name,
		event:   event,
		watches: watches,
	}

	return &out
}

// Name returns the name
func (obj *module) Name() string {
	return obj.name
}

// HasEvent returns true if there is an event, false otherwise
func (obj *module) HasEvent() bool {
	return obj.event != nil
}

// Event returns the event
func (obj *module) Event() ExecuteFn {
	return obj.event
}

// HasWatches returns true if there is watches, false otherwise
func (obj *module) HasWatches() bool {
	return obj.watches != nil
}

// Watches returns the watches, if any
func (obj *module) Watches() Watches {
	return obj.watches
}
