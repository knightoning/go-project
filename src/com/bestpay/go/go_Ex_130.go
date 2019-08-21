package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main()  {

	//execpath,err := os.Executable()
	//if err != nil{
	//	fmt.Println("获取程序路径异常")
	//}

	//absolPath := filepath.Join(filepath.Dir(execpath),"./products.txt")
	absolPath,err := filepath.Abs(filepath.Dir("products.txt"))

	if err != nil{
		fmt.Println("获取项目路径错误")
		return
	}

	inputFile := filepath.Join(absolPath,"/src/products.txt")
	outputFile := filepath.Join(absolPath,"/src/products_copy.txt")

	buf,err := ioutil.ReadFile(inputFile)

	if err != nil{
		fmt.Fprintf(os.Stderr,"File Error: %s\n",err)
	}

	fmt.Printf("%s\n",string(buf))

	err = ioutil.WriteFile(outputFile,buf,0644)
	if err != nil{
		panic(err.Error())
	}
}