package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, issue := range result.Items {
		fmt.Printf("#%-5d %s %.55s - %s\n",
			issue.Number, issue.Title, issue.HTMLURL, issue.CreatedAt.Local())
	}
}
