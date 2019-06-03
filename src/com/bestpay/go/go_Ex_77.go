package main

import (
	"fmt"
	"strings"
)

func main()  {

	//var str string = "This is an example of a string"
	//fmt.Printf("T/F? Does the string \"%s\" have prefix %s?",str,"Th")
	//fmt.Printf("%t\n",strings.HasPrefix(str,"Th"))

	var str string = "Hi, I'm Marc, Hi."

	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n",strings.Index(str,"Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n",strings.Index(str,"Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n",strings.LastIndex(str,"Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n",strings.Index(str,"Burger"))

	fmt.Printf("Number of H's in %s: ",str)
	fmt.Printf("%d\n",strings.Count(str,"H"))

	var manyG = "gggggggggg"
	fmt.Printf("Number of double g's in %s is:",manyG)
	fmt.Printf("%d\n",strings.Count(manyG,"gg"))

	var origS string = "Hi there!"
	var newS string

	newS =  strings.Repeat(origS,3)
	fmt.Printf("The new repeated string is: %s\n",newS)

	var orig string = "Hey, how are you George?"
	var lower string
	var upper string

	fmt.Printf("The original string is: %s\n",orig)
	lower = strings.ToLower(orig)
	fmt.Printf("The lowercase string is; %s\n",lower)
	upper = strings.ToUpper(orig)
	fmt.Printf("The uppercase string is: %s\n",upper)

	strF := "The quick brown fox jumps over the lazy dog"
	s1 := strings.Fields(strF)
	fmt.Printf("Splitted in slice: %v\n",s1)
	for _,val := range s1{
		fmt.Printf("%s -",val)
	}

	fmt.Println()

	str2 := "G01|The ABC of Go|25"
	sl2 := strings.Split(str2,"|")
	fmt.Printf("Splitted in slice: %v\n",sl2)
	for _,val := range sl2{
		fmt.Printf("%s -",val)
	}

	fmt.Println()

	str3 := strings.Join(sl2,";")
	fmt.Printf("sl2 joined by ;: %s\n",str3)
}
