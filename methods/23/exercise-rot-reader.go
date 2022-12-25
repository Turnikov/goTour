package main

import (
	"io"
	"os"
	"strings"
)

func byteRot13(x byte) byte {
	if 'a' <= x && x <= 'm' || 'A' <= x && x <= 'M' {
		x += 13
	} else if 'n' <= x && x <= 'z' || 'N' <= x && x <= 'Z' {
		x -= 13
	}
	return x
}

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	bytesCount, err := rot13.r.Read(b)
	if err == nil {
		for i, v := range b {
			b[i] = byteRot13(v)
		}
	}
	return bytesCount, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
