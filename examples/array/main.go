package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

func main() {
	//数组会自动初始化零值
	var arr [3]int

	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)
	fmt.Println(reflect.TypeOf(arr)) //[1,2,0],零值填充
	var b [0]int
	var c []int
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c)) //c为切片
	//数组字面量
	var strs [2]string = [2]string{
		"abc",
		"def",
	}
	fmt.Println(strs)
	s := [2]int{3, 4}
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	for i, v := range s {
		fmt.Println(i, ":", v)
	}

	for i := range s {
		fmt.Println(i)
	}

	for _, v := range s {
		fmt.Println(v)
	}
	//简写
	fmt.Println([10]int{})
	input, err := GetFloats("/Users/techzealot/goProjects/go-playground/examples/array/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
}

func average(nums [3]float64) float64 {
	sum := 0.0
	for _, value := range nums {
		sum += value
	}
	return sum / float64(len(nums))
}

// GetFloats 使用数组无法处理不确定数量的情况
func GetFloats(fileName string) ([3]float64, error) {
	var nums [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		return nums, err
	}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		nums[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nums, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return [3]float64{}, err
	}
	if scanner.Err() != nil {
		return nums, scanner.Err()
	}
	return nums, nil
}
