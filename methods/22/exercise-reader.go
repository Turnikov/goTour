package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := 0; ; {
		b[i] = 'A'
		i++
		return i, nil
	}
}

func main() {
	reader.Validate(MyReader{})
}
