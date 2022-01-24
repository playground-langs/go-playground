package main

import (
	"fmt"
	"go-playground/examples/errors"
	"log"
	"os"
)

func main() {
	numbers, err := errors.GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum : %0.2f \n\n", sum)
}
