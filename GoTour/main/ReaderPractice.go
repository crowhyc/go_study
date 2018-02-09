package main

import "golang.org/x/tour/reader"

type MyReader struct {
}

func (my MyReader) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = byte('A')

	}
	return 1, nil
}
func main() {
	reader.Validate(MyReader{})
}
