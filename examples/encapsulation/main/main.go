package main

import (
	"fmt"
	"go-playground/examples/encapsulation"
	"log"
)

func main() {
	d := encapsulation.Date{}
	fmt.Println(d)
	err := d.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}
	err = d.SetMonth(1)
	if err != nil {
		log.Fatal(err)
	}
	err = d.SetDay(22)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)
	fmt.Println(d.Year(), d.Month(), d.Day())
	l := encapsulation.Coordinates{}
	l.SetLongitude(11.0)
	l.SetLatitude(120.0)
	fmt.Println(l)
	e := encapsulation.Event{}
	_ = e.SetTitle("test")
	_ = e.SetYear(2022)
	_ = e.SetMonth(1)
	_ = e.SetDay(24)
	fmt.Println(e)
	fmt.Println(e.Title())
}
