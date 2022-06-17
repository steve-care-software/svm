package lexers

type assignment struct {
    content string
    name string
    declaration Variable
}

func createAssignmentWithName(
    content string,
    name string,
) Assignment {
    return createAssignmentInternally(content, name, nil)
}

func createAssignmentWithDeclaration(
    content string,
    declaration Variable,
) Assignment {
    return createAssignmentInternally(content,"", declaration)
}

func createAssignmentInternally(
    content string,
    name string,
    declaration Variable,
) Assignment {
    out := assignment{
        content: content,
        name: name,
        declaration: declaration,
    }

    return &out
}

// Content returns the content
func (obj *assignment) Content() string {
    return obj.content
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
