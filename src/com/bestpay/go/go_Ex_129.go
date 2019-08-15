package main

import (
	"fmt"
	"os"
	"bufio"
	"io")

func main()  {

	inputFile,inputError := os.Open("//Users//hujie//Documents//workspace//private//go-project//src//input.bat")

	if inputError !=nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
				"Have you got acces to it?\n")
		return
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for{
		inputString,readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s",inputString)
		if readerError != io.EOF{
			return
		}
	}
}