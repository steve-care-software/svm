package parsers

import "errors"

type attachmentBuilder struct {
	program string
	module  Variable
}

func createAttachmentBuilder() AttachmentBuilder {
	out := attachmentBuilder{
		program: "",
		module:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *attachmentBuilder) Create() AttachmentBuilder {
	return createAttachmentBuilder()
}

// WithProgram adds a program to the builder
func (app *attachmentBuilder) WithProgram(program string) AttachmentBuilder {
	app.program = program
	return app
}

// WithModule adds a module to the builder
func (app *attachmentBuilder) WithModule(module Variable) AttachmentBuilder {
	app.module = module
	return app
}

// Now builds a new Attachment instance
func (app *attachmentBuilder) Now() (Attachment, error) {
	if app.program == "" {
		return nil, errors.New("the program variable name is mandatory in order to build an Attachment instance")
	}

	if app.module == nil {
		return nil, errors.New("the module variable is mandatory in order to build an Attachment instance")
	}

	return createAttachment(app.program, app.module), nil
}
