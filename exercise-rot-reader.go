package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(bytes []byte) (int, error) {
	length, err := reader.r.Read(bytes)

	for i := 0; i < length; i++ {
		bytes[i] = doRot13(bytes[i])
	}
	return length, err
}

func doRot13(value byte) byte {
	if value < 'N' || value < 'n' {
		return value + 13
	}
	return value - 13
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
