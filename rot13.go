package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(p []byte) (n int, err error) {
	len, err := rot13.r.Read(p)
	for i := 0; i < len; i++ {
		c := p[i]
		if c >= 65 && c < 91 {
			p[i] = ((c-65+13)%26 + 65)
		} else if c >= 97 && c < 123 {
			p[i] = ((c-97+13)%26 + 97)
		} else {
			// Do nothing
		}
	}
	return len, nil
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
