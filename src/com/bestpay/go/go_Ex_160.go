package main

import (
	"flag"
	"fmt"
	)

var ngroutine = flag.Int("n",100000,"how many goroutines")

func f(left,right chan int)  {
	 left <- 1 + <- right
}

func main()  {

	flag.Parse()
	leftmost := make(chan int)
	var left,right chan int = nil,leftmost
	for i := 0; i < *ngroutine; i++ {
		left,right = right,make(chan int)
		go f(left,right)
	}

	right <- 0
	x := <- leftmost
	fmt.Println(x)
}