package main

import (
	"io"
	"strings"
	"os"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (int, error) {
	rInfo := make([]byte, len("Lbh penpxrq gur pbqr!"))
	strLen, err := rot.r.Read(rInfo)
	if err != nil {
		return 0, err
	}
	for i, v := range rInfo {
		p[i] = v + 13
	}
	return strLen, nil
}

func main() {
	src := "Lbh penpxrq gur pbqr!"
	s := strings.NewReader(src)
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

