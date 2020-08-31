package main

import "golang.org/x/tour/pic"


func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	pivot := 50
	p1, p2 := 18, 64
	pmax := 2000
	for i := range slice {
		slice[i] = make([]uint8, dx)
		for i2 := range slice[i] {
			if pivot > pmax { pivot /= p1 } else { pivot *= p2 }
			slice[i][i2] = uint8(pivot)
		}
	}
	return slice
}

func main() {
	pic.Show(Pic)
}