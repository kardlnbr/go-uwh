package main

import "fmt"

var result rune

func QuadC(x, y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 && j == 1) || (i == 1 && j == x) {
				result = 'A'
			} else if (j == 1 && i == y) || (j == x && i == y) {
				result = 'C'
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
	QuadC(5, 3)
}
