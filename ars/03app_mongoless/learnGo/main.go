package main

import (
	"fmt"
)

type Car struct {
	Id int 
	Number string 
	Model string 
	Type string 
}

func main() {
	//car1 := Car {Id:101, Number:"TN48 Z7878", Model:"Ambasister", Type:"Sedan"}
	var car1 Car = Car {Id:101, Number:"TN48 Z7878", Model:"Ambasister", Type:"Sedan"}
	fmt.Println(car1)
	fmt.Println(car1.Type)

	var cars[] Car = [] Car {
		{Id:101, Number:"TN48 Z7878", Model:"Ambasister", Type:"Sedan"},
		{Id:102, Number:"KA09 C8756", Model:"Toyota Innova", Type:"SUM"},
	}
	fmt.Println(cars)

	var car2* Car = &car1
	fmt.Println(car2.Model)
	/*
	fmt.Println("Hello World!!!");
	var first int = 10
	var second int = 20
	var sum int = first + second 
	

	//var salaries[] float32 = [] float32 {1.0, 2.0}
	salaries := [] float32 {1.0, 2.0}
	fmt.Println(sum)
	fmt.Println(salaries)*/
}