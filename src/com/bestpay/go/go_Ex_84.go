package main

import (
	"errors"
	"fmt"
	"math"
)

func main()  {

	fmt.Printf("First example with -1:")

	ret1, err1 := MySqrt(-1)

	if err1 != nil{
		fmt.Printf("Error! Return values are: ",ret1,err1)
	}
}

func MySqrt(f float64)(float64,error)  {

	if f < 0 {
		return float64(math.NaN()),errors.New("I won't be able to do a sqrt" +
			"")
	}
}
