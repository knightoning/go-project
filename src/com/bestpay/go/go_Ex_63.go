
//这个示例程序展示如何在程序里造成竞争状态
//实际上不希望出现这种情况
package main

import "fmt"
import "runtime"
import "sync"

var (
	//counter 是所有goroutine 都要增加其值的变量
	counter int

	// wg 用来等待程序结束
	wg sync.WaitGroup
)

func main()  {


	//计数加2，表示要等待两个goroutine
	wg.Add(2)

	//创建两个goroutine
	go incCounter(1)
	go incCounter(2)

	//等待goroutine 结束
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int)  {

	//在函数退出时调用Done 来通知main函数工作已经完成
	defer  wg.Done()

	for count := 0; count < 2; count++ {

		//捕获counter 的值
		value := counter

		//当前goroutine 从线程退出，并放回到队列
		runtime.Gosched()

		//增加本地value变量的值
		value ++

		//将该值保持回counter
		counter  = value
	}
}