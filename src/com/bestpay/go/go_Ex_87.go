package main

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs[LIM] uint64

// 通过内存缓存来提升性能
func main()  {

	var result uint64 = 0
	start := time.Now()

	for i := 0; i < LIM; i++{
		result =  fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n",i,result)
	}

	end := time.Now()
	delta := end.Sub(start)

	fmt.Printf("longCalculation took this amount of time: %s\n",delta)

}

func fibonacci(n int) (res uint64) {

	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	if n <= 1{
		res = 1
	}else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}

	// 临时缓存第n个计算的值
	fibs[n] = res
	return

}