package main

import (
	"io"
	"strings"
	"os"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	ret := make([]byte, 1024)
	i, err := rot.r.Read(ret)
	if err != nil {
		return 0, err
	}
	for i, v := range ret[:i] {
		b[i] = v + 13
	}
	return i, nil

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
