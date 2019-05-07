package main

import (
	"log"
	"net/http"
	"time"
)

const(
	// Poller Go 线程的启动数
	numPollers     = 2
	// 轮询每一个URL的频率
	pollInterval   = 60 * time.Second
	// 将状态记录到标准输出的频率
	statusInterval = 10 * time.Second
	//退回超时的错误
	errTimeout     = 10 * time.Second
)


var urls = []string{
	"http://www.google.com/",
	"http://golang.org/",
	"http://blog.golang.org",
}

// State 表示一个URL最后的已知状态。
type State struct {
	url    string
	status string
}

//StateMonitor 维护了一个映射，它存储了URL被轮询的状态，并每隔 updateInterval
//纳秒打印出其当前的状态。它向资源状态的接收者返回一个 chan State。
func StateMonitor(updateInterval time.Duration) chan <- State {

	updates := make(chan State)
	urlStatus := make(map[string]string)
	ticker := time.NewTicker(updateInterval)

	go func() {
		for {
			select {
			   case <- ticker.C:
			   	     logState(urlStatus)
			   case s := <- updates:
			   	     urlStatus[s.url] =  s.status
			}
		}
	}()

	return  updates
}

// logState 打印一个状态映射。
func logState(s map[string]string)  {

	log.Println("Current state:")
	for k,v := range s{
		log.Printf("%s %s",k,v)
	}
}

// Resource 表示一个被此程序轮询的HTTP UTL。
type Resource struct {
	url      string
	errCount int
}

// Poll 为url执行一个HTTP HEAD 请求，并返回HTTP的状态字符串或者一个错误字符串
func (r *Resource) Poll() string {

	resp,err := http.Head(r.url)

	if err != nil {
		log.Println("Error",r.url,err)
		r.errCount ++
		return err.Error()
	}

	r.errCount = 0
	return resp.Status
}

// Sleep 在将Resource 发送到 done之前的休眠一段适当的时间（取决于错误状态）
func (r *Resource) Sleep(done chan<- *Resource)  {

	time.Sleep(pollInterval + errTimeout*time.Duration(r.errCount))
	done <- r
}

func Poller(in <- chan *Resource,out chan<- *Resource,status chan <- State)  {

	for r := range in {
		s := r.Poll()
		status <- State{r.url,s}
		out <- r
	}
}

func main()  {

	//创建我们的输入和输出信道
	pending, complete := make(chan *Resource),make(chan *Resource)

	//启动 StateMonitor
	status := StateMonitor(statusInterval)

	//启动一些 Poller Go 线程
	for i := 0; i < numPollers; i++{
		go Poller(pending,complete,status)
	}

	//将 一些 Resource 发送至 pending 序列
	go func() {
		for _,url := range urls{
			pending <- &Resource{url: url}
		}
	}()

	for r := range complete{
		go r.Sleep(pending)
	}
}
