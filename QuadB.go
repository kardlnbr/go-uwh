package main

import "fmt"

func QuadB(x, y int) {
	var result rune
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			
			if (i == 1 && j == 1) || (i == y && j == x) {
				result = '/' 
			} else if (i == y && j == 1) || (i == 1 && j == x) {
				result = '\\' // Eğik slaş
			} else if j==1 || j==x || i==1 || i== y{
				result='*'
			} else{
				result=' '
			}

			fmt.Printf("%c", result) 
		}
		fmt.Println() 
	}
}

func main() {
	QuadB(5,3) 
}
