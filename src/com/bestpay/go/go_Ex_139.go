package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"path/filepath"
)

func main()  {

	absoltInputFile,err := filepath.Abs(filepath.Dir("goprogram.txt"))
	if err != nil{
		fmt.Println("获取目标文件路径错误")
		return
	}

	inputFile,_ := os.Open(filepath.Join(absoltInputFile,"/src/goprogram.txt"))

	absoltOutputFile,err := filepath.Abs(filepath.Dir("goprogramT.txt"))
	if err != nil{
		fmt.Println("获取目标文件路径错误")
		return
	}

	outputFile,_ := os.OpenFile(filepath.Join(absoltOutputFile,"/src/goprogramT.txt"),os.O_WRONLY|os.O_CREATE,0666)

	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)

	for {
		inputString,_,readerError := inputReader.ReadLine()
		if readerError == io.EOF{
			fmt.Println("EOF")
			return
		}

		outputString := string(inputString[2:5]) + "\r\n"
		_,err := outputWriter.WriteString(outputString)
		if err != nil{
			fmt.Println(err)
			return
		}

		//刷新数据，写入文件
		outputWriter.Flush()
	}

	fmt.Println("Conversion done")
}
