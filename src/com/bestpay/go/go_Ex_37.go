package main

import "fmt"
import "os"

/**

Defer 被用来确保一个函数调用在程序执行结束前执行。同样用来执行一些清理工作。
defer 用在像其他语言中的ensure 和 finally用到的地方。
 */
func createFile(p string)  *os.File {

	fmt.Println("creating")

	f,err := os.Create(p)

	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File)  {
	fmt.Println("wring")
	fmt.Fprintln(f,"data")
}

/**
在 closeFile 后得到一个文件对象，我们使用 defer通过 closeFile
来关闭这个文件。这会在封闭函数（main）结束时执行，就是 writeFile 结束后。
 */
func closeFile( f *os.File)  {
	fmt.Println("closing")
	f.Close()
}
func main()  {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}
