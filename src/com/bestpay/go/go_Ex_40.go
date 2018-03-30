package main

import "fmt"
//import "os"

type point struct {
	x,y int
}

func main()  {

	p := point{1,2}

	/**
	Go 为常规 Go 值的格式化设计提供了多种打印方式。例如，这里打印了 point 结构体的一个实例。
	 */
	fmt.Printf("%v\n",p)

	/**
	如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
	 */
	fmt.Printf("%+v\n",p)

	/**
	%#v 形式则输出这个值的 Go 语法表示。例如，值的运行源代码片段。
	 */
	fmt.Printf("%#v\n",p)

	/**
	需要打印值的类型，使用 %T。
	 */
	fmt.Printf("%T\n",p)

	/**
	格式化布尔值是简单的
	 */
	fmt.Printf("%t\n",true)

	/**
	格式化整形数有多种方式，使用 %d进行标准的十进制格式化。
	 */
	fmt.Printf("%d\n",123)

	/**
	这个输出二进制表示形式。
	 */
	fmt.Printf("%b\n",14)

	/**
	这个输出给定整数的对应字符。
	 */
	fmt.Printf("%c\n",33)

	/**
	%x 提供十六进制编码。
	 */
	fmt.Printf("%x\n",456)

	/**
	对于浮点型同样有很多的格式化选项。使用 %f 进行最基本的十进制格式化。
	 */
	fmt.Printf("%f\n",78.9)

	/**
	%e 和 %E 将浮点型格式化为（稍微有一点不同的）科学技科学记数法表示形式。
	 */
	fmt.Printf("%e\n",123400000.0)
	fmt.Printf("%E\n",123400000.0)

	/**
	使用 %s 进行基本的字符串输出。
	 */
	fmt.Printf("%s\n","\"string\"")

	/**
	像 Go 源代码中那样带有双引号的输出，使用 %q。
	 */
	fmt.Printf("%q\n", "\"string\"")
}
