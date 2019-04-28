//这个示例程序展示如何创建goroutine
//以及goruntine 调度器的行为
package main

import "fmt"
import "runtime"
import "sync"

//main 是所有Go程序的入口
func main()  {

	//分配2个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(2)

	//wg用来等待程序完成
	//计数加2,表示要等待两个gorountine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	//声明一个匿名函数，并创建一个goroutine
	go func() {

		//在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		//显示字母3次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a' + 26; char++ {
				fmt.Printf("%c ",char)
			}
		}
	}()

	//声明一个匿名函数，并创建一个goroutine
	go func() {

		//在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		//显示字母3次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A' +26; char++ {
				fmt.Printf("%c ",char)
			}
		}

	}()

	//等待goroutine 结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

