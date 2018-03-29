package main

import "fmt"

import "time"

/**
速率限制
 */

 /**
 速率限制(英) 是一个重要的控制服务资源利用和质量的途径。
 Go 通过 Go 协程、通道和打点器优美的支持了速率限制。
  */
func main()  {

	/**
	首先我们将看一下基本的速率限制。假设我们想限制我们接收请求的处理，
	我们将这些请求发送给一个相同的通道。
	 */
	requests := make(chan int ,5)

	for i:=1;i <=5;i++{
		requests <- i
	}

	close(requests)

	/**
	这个 limiter 通道将每 200ms 接收一个值。这个是速率限制任务中的管理器。
	 */
	limiter := time.Tick(time.Millisecond *200)

	/**
	通过在每次请求前阻塞 limiter 通道的一个接收，我们限制自己每 200ms 执行一次请求。
	 */
	for req := range requests{
		<- limiter
		fmt.Println("request",req,time.Now())
	}

	burstyLimiter := make(chan time.Time,3)

	for i:=0; i<3;i++{
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond *200){
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int,5)

	for i := 1;i <= 5 ; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)

	for req := range burstyRequests{
		<-burstyLimiter
		fmt.Println("request",req,time.Now())
	}
}

