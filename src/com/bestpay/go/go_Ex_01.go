package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	var age int
	fmt.Println(age)

	var istrue = true
	var istrue2 bool = true
	fmt.Println(istrue)
	fmt.Println(istrue2)

	age = 1000
	fmt.Println(age)

	var app string
	var str1 string ="Hello "
	var str2 string ="World"
	app = str1 + str2
	fmt.Println(app)

	var a = 1.0
	fmt.Println(a)

	fmt.Println("go" +"lang")
	fmt.Println("1+1=",1+1)
	fmt.Println("7.0/3.0 =",7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// var f string = "short" 简写
	f := "short"
    fmt.Println(f)


}
