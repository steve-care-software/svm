package lexers

type kind struct {
	flag   uint8
	module string
	name   string
}

func createKind(
	flag uint8,
	module string,
	name string,
) Kind {
	out := kind{
		flag:   flag,
		module: module,
		name:   name,
	}

	return &out
}

// Flag returns the flag
func (obj *kind) Flag() uint8 {
	return obj.flag
}

// Module returns the module
func (obj *kind) Module() string {
	return obj.module
}

// Name returns the name
func (obj *kind) Name() string {
	return obj.name
}
