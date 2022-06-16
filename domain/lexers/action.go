package lexers

type action struct {
    scope Scope
    application string
    isAttach bool
}

func createAction(
    scope Scope,
    application string,
    isAttach bool,
) Action {
    out := action {
        scope: scope,
        application: application,
        isAttach: isAttach,
    }

    return &out
}

// Scope returns the scope
func (obj *action) Scope() Scope {
    return obj.scope
}

// Application returns the application
func (obj *action) Application() string {
    return obj.application
}

// IsAttach returns true if attach, false otherwise
func (obj *action) IsAttach() bool {
    return obj.isAttach
}
