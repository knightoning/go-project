package main

import (
	"os"
	"bufio"
	"fmt"
	"path/filepath"
)

func main()  {

	absolPath,err := filepath.Abs(filepath.Dir("output.bat"))
	if err != nil{
		fmt.Println("获取目标文件路径错误")
		return
	}
	absolPath = filepath.Join(absolPath,"/src/output.bat")

	outputFile,outputError := os.OpenFile(absolPath,os.O_WRONLY|os.O_CREATE,0666)

	if outputError != nil{
		fmt.Println("An error occurred with file opening or creation\n")
		return
	}

	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i:=0; i<10; i++{
		outputWriter.WriteString(outputString)
	}

	outputWriter.Flush();
}