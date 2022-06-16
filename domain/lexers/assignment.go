package lexers

type assignment struct {
    name string
    declaration Variable
}

func createAssignmentWithName(
    name string,
) Assignment {
    return createAssignmentInternally(name, nil)
}

func createAssignmentWithDeclaration(
    declaration Variable,
) Assignment {
    return createAssignmentInternally("", declaration)
}

func createAssignmentInternally(
    name string,
    declaration Variable,
) Assignment {
    out := assignment{
        name: name,
        declaration: declaration,
    }

    return &out
}

// IsName returns true if there is a name, false otherwise
func (obj *assignment) IsName() bool {
    return obj.name != ""
}

// Name returns the name, if any
func (obj *assignment) Name() string {
    return obj.name
}

// IsDeclaration returns true if there is a declaration, false otherwise
func (obj *assignment) IsDeclaration() bool {
    return obj.declaration != nil
}

// Declaration returns the declaration, if any
func (obj *assignment) Declaration() Variable {
    return obj.declaration
}
