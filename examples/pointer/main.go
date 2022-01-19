package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	//指针声明
	var ptr *int
	//取地址
	ptr = &a
	//指针解引用
	fmt.Println(*ptr) //10
	//使用指针更新值
	*ptr = 20
	fmt.Println(a) //20
	fmt.Println(reflect.TypeOf(ptr))

	doubleInner(a)
	fmt.Println(a) //20
	double(&a)
	fmt.Println(a) //40
}

func doubleInner(input int) {
	input *= 2
}

func double(input *int) {
	*input *= 2
}
