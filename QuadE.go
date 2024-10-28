package main

import "fmt"

func QuadA(x, y int) {
	var result rune
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				result = 'A'
			} else if i == 5 && j == 1 {
				result = 'C'
			} else if (i == 1 && j == x) || (i == y && j == 1) {
				result = 'C'
			} else if i == 1 || i == y {
				result = 'B'
			} else if j == 1 || j == x {
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
	QuadA(5, 3)
})

