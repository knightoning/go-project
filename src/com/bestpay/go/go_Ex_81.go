package main

import "fmt"

func main()  {

	str := "Go is a beautiful language!"

	fmt.Printf("The length of str is: %d\n",len(str))

	for ix := 0;ix < len(str); ix++{
		fmt.Printf("The length of str is : %d is: %c \n",ix,str[ix])
	}

	str2 := "日本语"
	fmt.Printf("The length of str2 is: %d\n",len(str2))
	for ix := 0; ix <len(str2); ix++{
		fmt.Printf("Character on position %d is: %c \n",ix,str2[ix])
	}

	fmt.Println()

	for pos,char := range str{
		fmt.Printf("Character on position %d is: %c\n",pos,char)
	}

	str3 := "Chinese: 日本语"
	fmt.Printf("The length of str3 is: %d\n",len(str3))
	for pos, char := range str3{
		fmt.Printf("character %c starts at byte position %d\n",char,pos)
	}
	fmt.Println()

	fmt.Printf("index int(rune)rune  char bytes")
	for index,rune := range str3{
		fmt.Printf("%-2d     %d   %U '%c'  % X\n",index,rune,rune,rune,[]byte(string(rune)))
	}
}