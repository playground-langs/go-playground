package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Year(), now.Month(), now.Day())

	replacer := strings.NewReplacer("#", "o")

	fixed := replacer.Replace("G# r#cks")

	fmt.Println(fixed)

}
