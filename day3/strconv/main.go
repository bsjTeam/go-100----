package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)

	fmt.Println(f)

	n, _ := strconv.ParseInt("111", 10, 64)

	fmt.Println(n)

}
