package parsers

import (
	"fmt"
)

type screenWriter struct {
}

func createScreenWriterForTests() *screenWriter {
	return &screenWriter{}
}

func (p *screenWriter) Write(data []byte) (n int, err error) {
	fmt.Printf("%s\n", data)
	return len(data), nil
}
