package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())

	secret := rand.Intn(maxNum)

	fmt.Println("the secret is ", secret)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(" an err", err)
		return
	}

	input = strings.TrimSuffix(input, "\n")

	guess, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("invalid input . please enter an integer value")
		return
	}

	fmt.Println("You guess is", guess)

	if secret == guess {
		fmt.Println("winner")
		return
	}
	fmt.Println("device")

}
