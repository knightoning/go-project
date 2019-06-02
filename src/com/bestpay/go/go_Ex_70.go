package main

// 这个示例程序展示如何使用
// 有缓冲的通道和固定数目的
// goroutine 来处理一堆工作
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (

	// 要使用的 goroutine 的数量
	numberGoroutines = 4
	// 要处理的工作的数量
	taskLoad         = 10
)

// wg 用来等待程序完成
var wg sync.WaitGroup

// init 初始化包，Go 语言运行时会在其他代码执行之前
// 优先执行这个函数
func init()  {

	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
}

// main 是所有Go 程序的入口
func main()  {

	// 创建一个有缓冲的通道来管理工作
	tasks := make(chan string,taskLoad)

	// 启动 goroutine 来处理工作
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {

		go worker(tasks,gr)
	}

	// 增加一组要完成的工作
	for post := 1; post <= taskLoad; post ++{
		tasks <- fmt.Sprintf("Task: %d",post)
	}

	// 当所有工作都处理完成时关闭通道
	// 以便所有 goroutine 退出
	close(tasks)

	// 等待所有工作完成
	wg.Wait()
}


// worker 作为 goroutine 启动来处理
// 从有缓冲的通道传入的工作
func worker(tasks chan string, worker int)  {

	// 通知函数已经返回
	defer wg.Done()

	for{
		// 等待分配工作
		task, ok := <- tasks
		if !ok {

			// 这意味着通道已经空了，并且已被关闭
			fmt.Printf("Worker: %d : Shutting Down\n",worker)
			return
		}

		// 显示我们开始工作了
		fmt.Printf("Worker: %d : Started %s\n",worker,tasks)

		//随机等一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 显示我们完成工作
		fmt.Printf("Worker: %d : completed %s\n",worker,task)
	}
}