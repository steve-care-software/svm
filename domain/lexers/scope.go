package lexers

type scope struct {
    program string
    module string
}

func createScope(
    program string,
    module string,
) Scope {
    out := scope{
        program: program,
        module: module,
    }

    return &out
}

// Program returns the program
func (obj *scope) Program() string {
    return obj.program
}

// Module returns the module
func (obj *scope) Module() string {
    return obj.module
}
