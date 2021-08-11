package main

import (
	"fmt"
	go1 "go1/test1"
	"go1/test2"
	_ "go1/test2"
)

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		fmt.Println(result)
		return result
	}
	return 1
}

func main() {
	var a = true
	fmt.Println(a)
	if a {
		println("yes")
	}
	fmt.Println("hi")
	go1.Test001()
	var b float64 = 1.0
	println(b)
	circle := go1.Circle{Radius: 3}
	circle.Radius = 3.141592653
	var area = circle.GetArea()
	println(area)
	var test2_1 = test2.Student{Human: test2.Human{Age: 8, Name: "human"}, Name: "name", Value: "value", Score: 88, Id: "id"}
	test2_1.Name = ""
	test2_1.Human.Say()
	test2_1.Say()
	fmt.Println("test2_1 println")
	fmt.Println(test2_1)
	fmt.Println(test2_1.Human.Name)
	//test2_1.Println()
	var i int = 65
	//i = 100
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}
