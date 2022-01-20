package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//切片运算符可以用于切片,去除第一个参数(命令本身)，得到所有参数
	args := os.Args[1:]
	var input []float64
	for _, v := range args {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, f)
	}
	fmt.Println(average(input))
}

func average(input []float64) float64 {
	sum := 0.0
	for _, v := range input {
		sum += v
	}
	return sum / float64(len(input))
}
