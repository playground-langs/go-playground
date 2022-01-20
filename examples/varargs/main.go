package main

import "fmt"

func main() {
	fmt.Println(twoSum(1, 1))
	fmt.Println(nsum(1, 1, 1, 1))
	fmt.Println(nsum()) //接收到空切片
	fmt.Println(InRange(1, 10, 1, 2, 3, 4, 12, 40))
	fmt.Printf("%#v\n", InRange(1, 10))
	//可变参数接收切片
	var s []int = []int{1, 2, 3, 4, 5, 6, 7, 22, 33}
	fmt.Println(InRange(1, 20, s...))
}

func twoSum(a int, b int) int {
	return a + b
}

func nsum(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

// InRange 可变参数只能作为最后一个参数
func InRange(min int, max int, numbers ...int) []int {
	var in []int
	for _, v := range numbers {
		if min <= v && v <= max {
			in = append(in, v)
		}
	}
	return in
}
