// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

package main

import (
	"fmt"
	"io"
	"strconv"
)

type MyReader struct{}

func (reader *MyReader) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

type Rot13Reader struct {
	r io.Reader
}

func (rr *Rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rr.r.Read(p)
	for i := 0; i < n; i++ {
		c := p[i]
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			// Rotate uppercase letters by 13 positions
			if c >= 'A' && c <= 'Z' {
				c = 'A' + ((c-'A')+13)%26
			} else { // Rotate lowercase letters by 13 positions
				c = 'a' + ((c-'a')+13)%26
			}
			p[i] = c
		}
	}
	return n, err
}

func NewRot13Reader(r io.Reader) io.Reader {
	return &Rot13Reader{r}
}

func sumDigits() {
	str := "31122054131411225414314222231222313132233333322333333313322223321261223221211232322222"
	sum := 0

	for _, char := range str {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("Error converting character to integer:", err)
			return
		}
		sum += digit
	}

	fmt.Println("Sum of digits in the string:", sum)
}
