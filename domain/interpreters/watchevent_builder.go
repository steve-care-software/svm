package interpreters

import (
	"errors"
)

type watchEventBuilder struct {
	enter EnterWatchFn
	exit  ExitWatchFn
}

func createWatchEventBuilder() WatchEventBuilder {
	out := watchEventBuilder{
		enter: nil,
		exit:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *watchEventBuilder) Create() WatchEventBuilder {
	return createWatchEventBuilder()
}

// WithEnter adds an enter to the builder
func (app *watchEventBuilder) WithEnter(enter EnterWatchFn) WatchEventBuilder {
	app.enter = enter
	return app
}

// WithExit adds an exit to the builder
func (app *watchEventBuilder) WithExit(exit ExitWatchFn) WatchEventBuilder {
	app.exit = exit
	return app
}

// Now builds a new WatchEvent instance
func (app *watchEventBuilder) Now() (WatchEvent, error) {
	if app.enter != nil && app.exit != nil {
		return createWatchEventWithEnterAndExit(app.enter, app.exit), nil
	}

	if app.enter != nil {
		return createWatchEventWithEnter(app.enter), nil
	}

	if app.exit != nil {
		return createWatchEventWithExit(app.exit), nil
	}

	return nil, errors.New("the WatchEvent is invalid")
}
