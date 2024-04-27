package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: xkcd <command>\n\nAvailable commands: \n   xkcd init - Build up the comic index\n   xkcd find <search term> - Search for a Comic")
		return
	}

	if os.Args[1] == "init" {
		XkcdInit()
	} else if os.Args[1] == "find" {
		if len(os.Args) < 3 {
			fmt.Printf("Usage: xkcd <command>\n\nAvailable commands: \n   xkcd init - Build up the comic index\n   xkcd find <search term> - Search for a Comic")
			return
		}

		XkcdSearch(strings.ToLower(os.Args[2]))
	}
}
