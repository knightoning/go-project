package main
/**
有两个类型 stockPosition 和 car，它们都有一个 getValue() 方法，我们可以定义一个具有此方法的接口 valuable。
接着定义一个使用 valuable 类型作为参数的函数 showValue()，所有实现了 valuable 接口的类型都可以用这个函数。
 */

import "fmt"

type stockPostition struct {

	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPostition) getValue()float32 {

	return s.sharePrice * s.count
}

type car struct {

	make  string
	model string
	price float32
}

func (c car) getValue() float32  {

	return c.price
}

type valuable interface {

	getValue() float32
}

func showValue(asset valuable)  {

	fmt.Printf("Value of the asset is %f\n",asset.getValue())
}

func main()  {

	var o valuable = stockPostition{"GOOG",577.20,4}
	showValue(o)
	o = car{"BMW","M3",66500}
	showValue(o)
}