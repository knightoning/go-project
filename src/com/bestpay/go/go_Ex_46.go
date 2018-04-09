package main


import "fmt"
import (
	"math/rand"
	"time"
)

func main()  {

	 p := fmt.Println

	 p(rand.Intn(100),",")


	 p(rand.Intn(100))

	 p(rand.Float64())

	 p((rand.Float64()*5)+5)

	 s1 := rand.NewSource(time.Now().UnixNano())
	 p(s1)
	 r1 := rand.New(s1)

	 p(r1.Intn(100))
}
