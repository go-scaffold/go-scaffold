package processors

import (
	"io"
)

type mockProcessor struct {
	err       error
	processed bool
}

func (p *mockProcessor) ProcessFile(_ string, _ io.Reader) error {
	p.processed = true
	return p.err
}
