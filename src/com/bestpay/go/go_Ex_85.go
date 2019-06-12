package main

import "fmt"

// 关键字defer 允许我们推迟到函数返回之前（或任意位置执行return语句之后）一刻才执行某个语句或函数
func main()  {

	function1()

	doDBOperations()

	b()
}

func function1()  {

	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")


}

func function2()  {

	fmt.Printf("Function2: Deferred until the end of the calling function!")
}

func doDBOperations()  {

	connectToDB()
	fmt.Println("Defering the database disconnect.")

	defer disconnectFromDB()

	fmt.Println("Doing some DB operations...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return
}

func connectToDB()  {

	fmt.Println("ok,connected to db")
}

func disconnectFromDB()  {
	fmt.Println("ok,disconnected from db")
}

func trace(s string)  {
	fmt.Println("entering:",s)
}

func untrace(s string)  {
	fmt.Println("leaving:",s)
}
func b()  {

	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}

func a()  {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}