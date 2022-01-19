package main

import (
	"fmt"
	"log"
)

func main() {
	needed, err := paintNeeded(10, 20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed \n", needed)
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0.0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", width)
	}
	if height < 0.0 {
		return 0, fmt.Errorf("a height of %.2f is invalid", height)
	}
	area := width * height

	return area / 10.0, nil
}
