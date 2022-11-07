package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	t := time.Date(2022, 9, 11, 10, 10, 10, 0, time.UTC)
	t1 := time.Date(2022, 10, 11, 10, 11, 10, 0, time.UTC)

	fmt.Println(t)
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())

	fmt.Println(t.Format("2006-01-02 15:04:05"))
	diff := t1.Sub(t)

	fmt.Println(diff)
	fmt.Println(diff.Hours(), diff.Seconds(), now.Unix())

}
