package lexers

type variable struct {
    module string
    kind string
    name string
}

func createVariable(
    module string,
    kind string,
    name string,
) Variable {
    out := variable{
        module: module,
        kind: kind,
        name: name,
    }

    return &out
}

// Module returns the module
func (obj *variable) Module() string {
    return obj.module
}

// Kind returns the kind
func (obj *variable) Kind() string {
    return obj.kind
}

// Name returns the name
func (obj *variable) Name() string {
    return obj.name
}
