package main

import (
	"fmt"
)

func QuadE(x, y int) {
	var result rune
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 && j == x) || (i == y && j == 1) {
				result = 'C'
			} else if (i == 1 && j == 1) || (i == y && j == x) {
				result = 'A'
			} else if j == 1 || j == x || i == 1 || i == y {
				result = 'B'
			} else {
				result = ' '
			}
			fmt.Printf("%c", result)
		}
		fmt.Println()
	}
}

func main() {
	QuadE(5, 3)
}

