package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)

	fmt.Println("The sectet number is ", secretNumber)
}
