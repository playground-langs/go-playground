package main

import (
	"bufio"
	"fmt"
	"log"
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
