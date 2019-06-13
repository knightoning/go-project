package main

import "fmt"

func main()  {

	var arr1 [5]int

	for i:=0; i < len(arr1); i++{
		arr1[i] = i*2
	}

	//for i:=0; i < len(arr1); i++{
	//	fmt.Printf("Array at index %d is %d\n",i,arr1[i])
	//}

	for i,_:= range arr1{
		fmt.Printf("Array at index %d is %d\n",i,arr1[i])
	}

	var ar [3]int
	f(ar)
	fp(&ar)

	var arrKeyvalue = [5]string{3: "Chris",4:"ROn"}
	for i:=0; i < len(arrKeyvalue) ;i++{
		fmt.Printf("Person at %d is %s\n",i,arrKeyvalue[i])
	}
}

func f(a [3]int)  {
	fmt.Println(a)
}

func fp(a *[3]int)  {
	fmt.Println(a)

}
