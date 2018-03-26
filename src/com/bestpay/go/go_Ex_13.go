package main

import "fmt"

/**
结构体
 */

 /**
 这里的 person 结构体包含了 name 和 age 两个字段。
  */
type person struct{
	name string
	age int
}

func main(){

	fmt.Println(person{"BOb",20})

	fmt.Println(person{name:"Alice",age:30})

	fmt.Println(person{name:"Fred"})

	/**
	& 前缀生成一个结构体指针。
	 */
	fmt.Println(&person{name:"Ann",age:40})

	s := person{name:"Sean",age:50}
	fmt.Println(s.name)

	/**
	也可以对结构体指针使用. - 指针会被自动解引用。
	 */
	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
