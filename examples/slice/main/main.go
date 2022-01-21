package main

import (
	"embed"
	"fmt"
	"go-playground/examples/slice"
	"log"
)

//go:embed data.txt
var dataFs embed.FS

func main() {
	//切片不会自动初始化元素零值,切片的初始值为nil;len,append都把nil看做空切片
	var s1 []int
	fmt.Println(s1 == nil)
	fmt.Printf("%#v\n", s1)
	fmt.Println(len(s1))
	s1 = make([]int, 7)
	s2 := make([]int, 7)
	s2[2] = 1
	fmt.Println(s1)
	fmt.Println(s2)
	for i, v := range s2 {
		fmt.Println(i, v)
	}
	//切片字面量
	s3 := []int{1, 2, 3}
	fmt.Println(s3)
	//从数组创建切片,左闭右开 [m:n]~[m,n);[:n]~[0,n);[m:]~[m,len);[:]~[0,len)
	//切片运算符可以用于切片
	//对数组的修改会反映在切片上，且切片在可见范围内也可以修改数组
	//通常使用make和切片字面量创建切片
	arr := [5]int{1, 2, 3, 4, 5}
	s4 := arr[0:2]
	fmt.Println(s4)
	//切片增加数据,注意必须接受返回的切片作为新值，可能底层数组已经发生了改变
	//切片类似于java ArrayList中指向底层数组的指针
	s4 = append(s4, 16)
	fmt.Println(s4)
	file, err := dataFs.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	input, err := slice.GetFloats(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
}
