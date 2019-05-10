package main

import (
	"fmt"
	"sync"
	"time"
)

// wg 用来等待程序结束
var wg sync.WaitGroup

// main是所以Go程序的入口
func main()  {

	// 创建一个无缓冲的通道
	baton := make(chan int)

	// 为最后一位跑步者将计数加1
	wg.Add(1)

	// 第一位跑步者持有接力棒
	go Runner(baton)

}

// Runner模拟接力比赛中的一位跑步者
func Runner(baton chan int)  {

	var newRunner int

	// 等待接力棒
	runner := <- baton

	// 开始绕着跑到跑步
	fmt.Printf("Runner %d Running With Baton\n",runner)

	// 创建下一位跑步者

}
