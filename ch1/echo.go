// Go version of echo command
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// ex1
	fmt.Println(os.Args[1], " ")

	// ex2
	// for index, value := range os.Args[1:] {
	// 	fmt.Println(index, " ", value)
	// }

	var slice []string
	for i := 0; i < 1e5; i++ {
		slice = append(slice, "f ")
	}
	fmt.Println("completed slice")

	// ex3
	s, sep := "", ""
	start := time.Now()

	for _, value := range slice {
		s += sep + value
		sep = " "
	}
	fmt.Printf("%v %v\n", len(s), time.Since(start))

	s, sep = "", ""
	start = time.Now()
	sep = " "
	s = strings.Join(slice, sep)
	fmt.Printf("%v %v\n", len(s), time.Since(start))
}
