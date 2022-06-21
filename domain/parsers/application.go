package parsers

type application struct {
	application Variable
	attachments Attachments
}

func createApplication(
	application Variable,
) Application {
	return createApplicationInternally(application, nil)
}

func createApplicationWithAttachments(
	application Variable,
	attachments Attachments,
) Application {
	return createApplicationInternally(application, attachments)
}

func createApplicationInternally(
	app Variable,
	attachments Attachments,
) Application {
	out := application{
		application: app,
		attachments: attachments,
	}

	return &out
}

// Application returns the application
func (obj *application) Application() Variable {
	return obj.application
}

// HasAttachments returns true if there is attachments, false otherwise
func (obj *application) HasAttachments() bool {
	return obj.attachments != nil
}

// Attachments returns the attachments, if any
func (obj *application) Attachments() Attachments {
	return obj.attachments
}
