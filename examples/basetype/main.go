package main

import "fmt"

func main() {
	l := Liters(10.0)
	g := Gallon(240.0)
	fmt.Println(l, g)
	fmt.Println(l + 10)
	fmt.Println(g + Gallon(30))
	fmt.Println(l.ToGallon())
	l.Double()
	fmt.Println(l)
	//你只能获取保存在变量中的指针,只有保存在变量中，go才会帮忙自动转换
	//Liters(120.0).Double() error
}

// Liters 为了保护类型定义，方法和类型必须定义在同一个包中
//为了一致性，你所有的类型函数要么都接受值类型，要么都接受指针类型，你应该避免混用的情况
type Liters float64

//ToGallon 值接收器
func (l Liters) ToGallon() Gallon {
	return Gallon(l * 0.264)
}

//Double 若要修改接收器，必须使用 指针接收器
//当你用一个非指针的变量调用一个需要指针的接收器的方法的时候，Go会自动为你将非指针类型转换为指针类型。
//同样指针类型也会自动转换为非指针类型，如果你调用一个要求值类型的接收器，Go会自动帮你获取指针指向的值，然后传递给方法
func (l *Liters) Double() {
	*l *= 2
}

type Gallon float64

type MilliLiters float64

func (m MilliLiters) ToGallon() Gallon {
	return Gallon(m * 0.264 * 0.001)
}

// LitersToGallons seems bad
func LitersToGallons(l Liters) Gallon {
	return Gallon(l * 0.264)
}

func MilliLitersToGallons(m MilliLiters) Gallon {
	return Gallon(m * 0.264 * 0.001)
}

func ToLiters(g Gallon) Liters {
	return Liters(g * 3.785)
}
