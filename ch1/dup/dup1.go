// Program that prints all the duplicate lines and the files in which those lines appear
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	count := make(map[string]int)

	filesReferences := make(map[string]map[string]bool)
	files := os.Args[1:]
	for _, path := range files {
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not open %s. Err: %v\n", path, err)
			continue
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			count[line]++
			if filesReferences[line] == nil {
				filesReferences[line] = make(map[string]bool)
			}
			filesReferences[line][path] = true
		}
	}
	for line, count := range count {
		if count > 1 {
			fmt.Printf("%s\nrepeated: %v times\n", line, count)
			var fileSlice []string
			for file := range filesReferences[line] {
				fileSlice = append(fileSlice, file)
			}
			fmt.Printf("%s\n", strings.Join(fileSlice, ", "))
		}
	}
}
