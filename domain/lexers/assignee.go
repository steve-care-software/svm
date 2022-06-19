package lexers

type assignee struct {
    name string
    declaration Variable
}

func createAssigneeWithName(
    name string,
) Assignee {
    return createAssigneeInternally(name, nil)
}

func createAssigneeWithDeclaration(
    declaration Variable,
) Assignee {
    return createAssigneeInternally("", declaration)
}

func createAssigneeInternally(
    name string,
    declaration Variable,
) Assignee {
    out := assignee{
        name: name,
        declaration: declaration,
    }

    return &out
}

// IsName returns true if there is a name, false otherwise
func (obj *assignee) IsName() bool {
    return obj.name != ""
}

// Name returns the name, if any
func (obj *assignee) Name() string {
    return obj.name
}

// IsDeclaration returns true if there is a declaration, false otherwise
func (obj *assignee) IsDeclaration() bool {
    return obj.declaration != nil
}

// Declaration returns the declaration, if any
func (obj *assignee) Declaration() Variable {
    return obj.declaration
}
