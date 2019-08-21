package main

import (
	"flag"
	"os"
)

var NewLine = flag.Bool("n",false,"print newLine")

const(
	Space =" "
	MewLine = "\n"
)

func main()  {

	flag.PrintDefaults()
	flag.Parse()
	var s string = ""
	for i:=0; i < flag.NArg(); i ++{
		if i > 0{
			s += " "
			if *NewLine {
				//s += NewLine
			}
		}
		s += flag.Arg(i)
	}

	os.Stdout.WriteString(s)
}