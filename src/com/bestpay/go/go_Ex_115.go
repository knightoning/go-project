package main
/**
一个接口类型的变量 varI 中可以包含任何类型的值，必须有一种方式来检测它的 动态 类型，即运行时在变量中存储的值的实际类型。
在执行过程中动态类型可能会有所不同，但是它总是可以分配给接口变量本身的类型。通常我们可以使用 类型断言 来测试在某个时刻 varI 是否包含类型 T 的值：
v := varI.(T)       // unchecked type assertion
varI 必须是一个接口变量
 */

import (
	"fmt"
	"math"
)

type Square struct {

	side float32
}

type Circle struct {

	radius float32
}

type Shaper interface {
	Area() float32
}

func (sq *Square)Area() float32 {
	return sq.side * sq.side
}

func (c *Circle)Area() float32  {
	return c.radius * c.radius * math.Pi
}
func main()  {

	var areaIntf Shaper
	sq1 := new (Square)
	sq1.side = 5

	areaIntf = sq1

	if t,ok := areaIntf.(*Square); ok{
		fmt.Printf("The type of areaIntf is: %T\n",t)
	}

	if u,ok := areaIntf.(*Circle); ok{
		fmt.Printf("The type of areaIntf is: %T\n",u)
	}else{
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}


}
