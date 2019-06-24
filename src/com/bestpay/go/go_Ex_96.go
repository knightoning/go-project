package main

import "fmt"

func main()  {

	map1 := make(map[int] float64)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	for key, value := range map1{
		fmt.Printf("Key is: %d -  value is: %f\n",key,value)
	}
}