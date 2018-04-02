package main


import (
	"fmt"
	"regexp"
	"bytes"
)


/**

正则表达式
 */
func main()  {

	//这个测试一个字符串是否符合一个表达式。
	match,_ := regexp.MatchString("p([a-z]+)ch","peach")
	fmt.Println(match)

	//上面我们是直接使用字符串，但是对于一些其他的正则任务，你需要 Compile 一个优化的 Regexp 结构体。
	r,_:= regexp.Compile("p([a-z]+)ch")


	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))

	fmt.Println(r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch",-1))

	fmt.Println(r.FindAllStringIndex("peach punch pinch",-1))

	fmt.Println(r.FindAllString("peach punch pinch",2))

	fmt.Println(r.Match([]byte("peach")))


	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach","<fruit>"))

	in :=[]byte("a peach")

	out := r.ReplaceAllFunc(in,bytes.ToUpper)
	fmt.Println(string(out))

}
