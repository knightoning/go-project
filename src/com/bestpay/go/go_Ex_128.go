package main

import (
	"fmt"
	"os"
	"bufio"
	)

func main()  {

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")

	input,err := inputReader.ReadString('\n')

	if err != nil{
		fmt.Println("There were errors reading,exiting program.")
		return
	}

	fmt.Println("Your name is %s",input)

	switch input {
	case "Philip\r\n": fmt.Println("Welcome Philip!")
	case "Chris\r\n": fmt.Println("Welcome Chris!")
	case "Ivo\r\n": fmt.Println("Welcome Ivo!")
	default:
		fmt.Println("You are not welcome here! Goodbye!")
	}

	switch input {
	case "Philip\r\n": fallthrough
	case "Ivo\r\n": fallthrough
	case "Chris\r\n": fmt.Printf("Welcome %s\n",input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	switch input {
	case "Philip\r\n","Ivo\r\n": fmt.Printf("Welcome %s\n",input)
	default:
		fmt.Printf("You are not welcome here! Goodbye\n")

	}
}
