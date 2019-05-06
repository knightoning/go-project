package main

import (
	"log"
	"net/http"
	"time"
)

const(
	numPollers     = 2
	pollInterval   = 60 * time.Second
	statusInterval = 10 * time.Second
	errTimeout     = 10 * time.Second
)


var urls = []string{
	"http://www.google.com/",
	"http://golang.org/",
	"http://blog.golang.org",
}

type State struct {
	url    string
	status string
}

//
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

func logState(s map[string]string)  {

	log.Println("Current state:")
	for k,v := range s{
		log.Printf("%s %s",k,v)
	}
}

type Resource struct {
	url      string
	errCount int
}

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
