package main
/**
现在接口变量包含一个指向 Square 变量的引用，通过它可以调用 Square 上的方法 Area()。当然也可以直接在 Square 的实例上调用此方法，但是在接口实例上调用此方法更令人兴奋，它使此方法更具有一般性。接口变量里包含了接收者实例的值和指向对应方法表的指针。

这是 多态 的 Go 版本，多态是面向对象编程中一个广为人知的概念：根据当前的类型选择正确的方法，或者说：同一种类型在不同的实例上似乎表现出不同的行为。
 */

import "fmt"

type Shaper interface {

	Area() float32
}

type Square struct {

	side float32
}

func (sq *Square) Area() float32  {

	return sq.side * sq.side
}

type Rectangle struct {

	length,width float32
}

func (r Rectangle) Area() float32  {

	return r.length * r.width
}

func main()  {


	r := Rectangle{5,3}
	q := &Square{5}

	shapes := []Shaper{r,q}
	fmt.Println("Looping through shapes for area...")

	for n,_ := range shapes{
		fmt.Println("Shape details: ",shapes[n])
		fmt.Println("Area of this shape is:",shapes[n].Area())
	}
}
