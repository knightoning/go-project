package main

import (
	"fmt"
)

var hello_word =""

func main()  {

	fmt.Println("收到 我马上执行")

	fmt.Println(hello_word)

	fmt.Println("我已经执行完成...")
}

/**
   init 初始化函数
 */
func init()  {

	fmt.Println("首先执行")
	fmt.Println("我已经执行了，该你了")

	hello_word = "Hello World!"

}