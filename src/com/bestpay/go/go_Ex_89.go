package main

const(
	WIDTH = 1920
	HEIGHT = 1080
)

type piexl int
var screen [WIDTH][HEIGHT] piexl

func main()  {

	for y := 0; y < HEIGHT; y++{
		for x := 0; x < WIDTH; x++{
			screen[x][y] = 0
		}
	}
}
