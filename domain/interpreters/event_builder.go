package interpreters

import (
	"errors"
)

type eventBuilder struct {
	enter ExecuteFn
	exit  ExecuteFn
}

func createEventBuilder() EventBuilder {
	out := eventBuilder{
		enter: nil,
		exit:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *eventBuilder) Create() EventBuilder {
	return createEventBuilder()
}

// WithEnter adds an enter to the builder
func (app *eventBuilder) WithEnter(enter ExecuteFn) EventBuilder {
	app.enter = enter
	return app
}

// WithExit adds an exit to the builder
func (app *eventBuilder) WithExit(exit ExecuteFn) EventBuilder {
	app.exit = exit
	return app
}

// Now builds a new Event instance
func (app *eventBuilder) Now() (Event, error) {
	if app.enter != nil && app.exit != nil {
		return createEventWithEnterAndExit(app.enter, app.exit), nil
	}

	if app.enter != nil {
		return createEventWithEnter(app.enter), nil
	}

	if app.exit != nil {
		return createEventWithExit(app.exit), nil
	}

	return nil, errors.New("the Event is invalid")
}
