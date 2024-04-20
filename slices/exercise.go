package main

func Pic(dx, dy int) [][]uint8 {
	retVal := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			row[x] = uint8((x + y) / 2)
		}
		retVal[y] = row
	}
	return retVal
}
