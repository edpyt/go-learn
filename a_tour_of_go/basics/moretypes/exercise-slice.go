package main

import 	"golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	dy_ss := make([][]uint8, dy)
	for y:= 0; y < len(dy_ss); y++ {
		dy_ss[y] = make([]uint8, dx)
	}
	return dy_ss
}

func main() {
	pic.Show(Pic)
}

