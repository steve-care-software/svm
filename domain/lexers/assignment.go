package lexers

type assignment struct {
	content  string
	assignee Assignee
}

func createAssignment(
	content string,
	assignee Assignee,
) Assignment {
	out := assignment{
		content:  content,
		assignee: assignee,
	}

	return &out
}

// Content returns the content
func (obj *assignment) Content() string {
	return obj.content
}

// Assignee returns the assignee
func (obj *assignment) Assignee() Assignee {
	return obj.assignee
}
