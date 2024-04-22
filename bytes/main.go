package main

import (
	"bytes"
	"fmt"
)

func main() {
	var bs bytes.Buffer

	c, err := bs.WriteString("hello world")

	if err != nil {
		fmt.Println("error writing string: ", err)
	}

	fmt.Println(c)
}
