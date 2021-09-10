package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(bytes []byte) (int, error) {
	n, err := rot13.r.Read(bytes)
	for i := 0; i < n; i++ {
		bytes[i] = rot13.rot13(bytes[i])
	}
	return n, err
}

func (rot13 *rot13Reader) rot13(b byte) byte {
	// ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
	// NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm
	if (b >= 'A' && b <= 'M') || (b >= 'a' && b <= 'm') {
		return b + 13
	} else if (b >= 'N' && b <= 'Z') || (b >= 'n' && b <= 'z') {
		return b - 13
	} else {
		return b
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
