package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"

	fmt.Println(strings.Contains(a, "2321"))
	fmt.Println(strings.Count(a, "l"))
	fmt.Println(strings.HasPrefix(a, "he"))
	fmt.Println(strings.HasSuffix(a, "lo"))
	fmt.Println(strings.Index(a, "ll"))
	fmt.Println(strings.Join([]string{"123", "ll"}, "l"))
	fmt.Println(strings.Count(a, "l"))
	fmt.Println(strings.Count(a, "l"))
}
