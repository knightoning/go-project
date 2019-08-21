package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main()  {

	absolPath,err := filepath.Abs(filepath.Dir("products2.txt"))
	if err != nil{
		fmt.Println("获取文件路径错误！")
		return
	}

	absolPath = filepath.Join(absolPath,"/src/products2.txt")
	file,err := os.Open(absolPath)
	if err != nil{
		panic(err)
	}

	defer file.Close()

	var col1,col2,col3 []string
	for{
		var v1,v2,v3 string
		_,err := fmt.Fscanln(file,&v1,&v2,&v3)

		if err != nil{
			break
		}

		col1 = append(col1,v1)
		col2 = append(col2,v2)
		col3 = append(col3,v3)
	}

	fmt.Printf("col1 := %s\n" ,col1)
	fmt.Printf("col2 := %s\n" ,col2)
	fmt.Printf("col3 := %s\n" ,col3)

}
