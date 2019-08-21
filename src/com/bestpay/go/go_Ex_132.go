package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main()  {

	targetPath, err := filepath.Abs(filepath.Dir("target.txt"))
	if err != nil{
		fmt.Println("获取目标文件路径异常！")
		return
	}
	targetPath = filepath.Join(targetPath,"/src/target.txt")

	sourcePath,err := filepath.Abs(filepath.Dir("source.txt"))
	if err != nil{
		fmt.Println("获取源文件路径异常！")
		return
	}
	sourcePath = filepath.Join(sourcePath,"/src/source.txt")

	CopyFile(targetPath,sourcePath)
	fmt.Println("Copy done!")
}

func CopyFile(dstName,srcName string) (written int64,err error) {

	src,err := os.Open(srcName)

	if err != nil{
		return
	}

	defer src.Close()

	dst,err := os.Create(dstName)

	if err != nil{
		return
	}

	defer dst.Close()

	return io.Copy(dst,src)
}
