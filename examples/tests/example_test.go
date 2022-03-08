package tests

import (
	"fmt"
)

// 参考io包 io/example_test.go
// T M 代表真正的方法 suffix代表细分场景,无可不写
// 包级别
func Example() {
	fmt.Println("example test")
	// Output: example test
}

func Example_x1() {
	fmt.Println("example test")
	// Output: example test
}

// 函数级别
func ExampleToUpperCase() {
	fmt.Println(ToUpperCase("abcd"))
	// Output:
	// ABCD
}

func ExampleToUpperCase_x1() {
	fmt.Println("example test")
	// Output:
	// example test
}

// 类型级别
func ExampleT_suffix() {
	fmt.Println("example test")
	// Output: example test
}

// 方法级别
func ExampleM_suffix() {
	fmt.Println("example test")
	// Output: example test
}

// 类型上方法级别
func ExampleT_M_suffix() {
	fmt.Println("example test")
	// Output: example test
}
