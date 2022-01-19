//Package a guess number game
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1
	fmt.Println("I've choose a number between 1 and 100.")
	fmt.Println("can you guess it?")
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		fmt.Println("you have ", 10-i, "left")
		fmt.Println("make a guess")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess > target {
			fmt.Println("high")
		} else if guess < target {
			fmt.Println("low")
		} else {
			fmt.Println("you get it")
			return
		}
	}
	fmt.Println("sorry,you didn't guess my number,it was ", target)

}
