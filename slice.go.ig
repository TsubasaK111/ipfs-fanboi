package main

import (
	"golang.org/x/tour/pic"
	//"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	var pic = make([][]uint8, dy)

	for x, _ := range pic {
		pic[x] = make([]uint8, dx)
		for y, _ := range pic[x] {
			pic[x][y] = (uint8(x) + uint8(y)) / 2
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
