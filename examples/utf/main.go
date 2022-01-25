package main

import "fmt"

func main() {
	asciiString := "abcd"
	for i, c := range asciiString {
		fmt.Println(i, ":", c)
	}
	utf8String := "测试字符串"
	//索引i代表字节索引，不是rune索引
	for i, s := range utf8String {
		fmt.Println(i, ":", s)
	}
	//错误处理方式
	for i, s := range []byte(utf8String) {
		fmt.Println(i, ":", s)
	}
}
