package iohelpers

import (
	"bytes"
	"io"
)

func Read(reader io.Reader) string {
	var buffer bytes.Buffer
	b := make([]byte, 8)
	for {
		n, err := reader.Read(b)
		buffer.Write(b[:n])
		if err == io.EOF {
			break
		}
	}
	return buffer.String()
}
