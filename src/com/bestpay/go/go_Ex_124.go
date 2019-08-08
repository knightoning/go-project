package main

/**
在 Go 语言的标准库中，前几节所述的反射的功能被大量地使用。
举个例子，fmt 包中的 Printf（以及其他格式化输出函数）
都会使用反射来分析它的 ... 参数。
 */

import "os"
import "strconv"

type Stringer interface {
	String() string
}

type Celsius float64

func (c Celsius)String() string  {
	return strconv.FormatFloat(float64(c),'f',1,64) + "  °C"
}

type Day int

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day)string()string  {
	return dayName[day]
}

func print(args ...interface{})  {

	for i,arg := range args{
		if i>0 {os.Stdout.WriteString(" ")}
		switch a := arg.(type) {
		case Stringer: os.Stdout.WriteString(a.String())
		case int:      os.Stdout.WriteString(strconv.Itoa(a))
		case string:   os.Stdout.WriteString(a)
		default:
			           os.Stdout.WriteString("???")
		
		}
	}
}

func main()  {
	print(Day(1),"was",Celsius(18.36))
}