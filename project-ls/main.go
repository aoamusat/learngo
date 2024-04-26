package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	output := strings.Join(listFiles("testdata"), " ")
	fmt.Print(output)
}

func listFiles(dirname string) []string {
	var dirs []string

	files, _ := os.ReadDir(dirname)

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
