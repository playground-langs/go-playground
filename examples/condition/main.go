package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("enter a grade:")

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	text = strings.TrimSpace(text)
	grade, err := strconv.ParseFloat(text, 64)
	if err != nil {
		log.Fatal(err)
	}
	var result string
	if grade >= 60 {
		result = "及格"
	} else {
		result = "不及格"
	}
	fmt.Println(result)
}

func ifInit() {
	if f, err := strconv.ParseFloat("1.234", 64); err != nil {
		fmt.Println(f)
	}
	if f, err := strconv.ParseFloat("23.234", 64); err != nil {
		fmt.Println(f)
	}
}

func switchTest() {
	switch i := rand.Intn(10); i {
	case 0:
		fmt.Println(i)
	case 1:
		fmt.Println(i)
	case 2:
		fmt.Println(i)
	default:
		fmt.Println(i)
	}
}
