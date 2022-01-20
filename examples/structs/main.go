// Package struct具有零值，作为参数传递时使用值传递，若要在方法内修改，必须传递指针
package main

import (
	"fmt"
	"go-playground/examples/structs/person"
)

func main() {
	var myStruct struct {
		name string
		age  int
		on   bool
	}
	fmt.Println(myStruct)
	myStruct.name = "tom"
	fmt.Println(myStruct.name)
	//struct字面量赋值
	p := person.Person{
		Name: "cat",
		Age:  20,
	}
	fmt.Println(p)
	errorModifyAge(p)
	fmt.Println("have no effect", p)
	SetAge(&p)
	fmt.Println(p)
	em := person.Employee{
		Name:    "a",
		Age:     10,
		Address: person.Address{Street: "street a", City: "c"},
	}
	fmt.Println(em)
	fmt.Println(em.Address.Street, ":", em.Address.City)
	fmt.Println(em.Street, ":", em.City)
}

func errorModifyAge(person person.Person) {
	person.Age = 100
}

func SetAge(person *person.Person) {
	//使用点运算符在struct指针和struct上都可访问字段
	//推荐
	person.Age = 100
	//ok
	(*person).Age = 100
}
