package main

import (
	_ "fmt"
	_ "math"
	"fmt"
	"math"
)


const (
	name   string = "sunlong"
	age    int    = 10
	height int    = 180
)

func control_if()  {

	if age >= 10 && name != "sunlong" {
		fmt.Printf("my name is %s,my age is %d",name,age)
	}

	if(age >= 10 && name == "sunlong2") || (height >=180){
		fmt.Printf("my name is %s,my age is %d",name,age)
	}
}

func control_switch(score float64) string {

	var res float64 = math.Floor( score / 10)
	var resstr string
	switch res {
	case 7:
		resstr ="合格"
	case 9:
		resstr = "优秀"
	default:
		resstr = "不合格"
	}

	return resstr
}

func control_for()  {
	var i int = 1
	for i <= 10 {
		fmt.Printf("\r\n 第%d个",i)
		i ++
	}
}

func main()  {

	fmt.Printf(name)

	control_if()

	var score = control_switch(20)
	fmt.Printf("\r\n 你的结果是：%s",score)

	control_for()
}

