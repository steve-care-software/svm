package parsers

type attachments struct {
	list []Attachment
}

func createAttachments(
	list []Attachment,
) Attachments {
	out := attachments{
		list: list,
	}

	return &out
}

// List returns the list of attachments
func (obj *attachments) List() []Attachment {
	return obj.list
}
