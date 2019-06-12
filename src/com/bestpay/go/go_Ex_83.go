package main

import "fmt"

var num int = 10
var numx2,numx3 int

func main()  {

	numx2,numx3 = getX2AndX3(num)
	PrintValues()
	numx2,numx3 = getX2AndX3_2(num)
	PrintValues()

	sum, prod, diff := SumProductDiff(3,4)
	fmt.Println("Sum:",sum,"| Product:",prod,"| Diff:",diff)
	sum, prod, diff = SumProductDiffN(3,4)
	fmt.Println("Sum:",sum,"| Product:",prod,"| Diff:",diff)

}

func PrintValues(){
	fmt.Printf("num = %d, 2x num = %d, 3x num= %d\n",num,numx2,numx3)
}
func getX2AndX3(input int)(int,int)  {

	return 2 * input, 3* input
}

func getX2AndX3_2(input int)(x2 int,x3 int)  {

	x2 = 2 * input
	x3 = 3 * input

	return
}

func SumProductDiff(i,j int)(int,int,int){
	return i + j, i * j,i - j
}

func SumProductDiffN(i,j int)(s int, p int,d int)  {
	s, p, d = i + j, i * j, i - j
	return
}