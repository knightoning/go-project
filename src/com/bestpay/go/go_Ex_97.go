package main

import (
	"fmt"
)

func main()  {

	items := make([]map[int]int,5)

	for i := range items{
		// 获取的只是items 的索引
		items[i] = make(map[int]int,1)
		items[i][1] = 2
	}

	fmt.Printf("Version A: Value of items: %v\n",items)

	items2 := make([]map[int]int,5)

	for _,item := range items2{
		// 获取的只是items2 的拷贝
		item = make(map[int]int,1)
		item[1] = 2
	}

	fmt.Printf("Version B; Value of items: %v\n",items2)
}
