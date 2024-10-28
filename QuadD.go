package main

import "fmt"

func QuadD(x, y int) {
	var result rune
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 && j == 1) || (i == y && j == 1) {
				result = 'A'
			} else if (i == 1 && j == x) || (i == y && j == x) {
				result = 'C'
			} else {
				result = 'B'
			}
			fmt.Printf("%c", result)
		}
		fmt.Println()
	}
}

func main() {
	QuadD(5, 3)
}
