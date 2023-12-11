package main

import (
	"fmt"
)

func printCircle(radius int, padding int, isHead bool, topTrim int, bottomTrim int, hasHands bool) {
	diameter := 2 * radius
	for y := -radius + topTrim; y <= radius-bottomTrim; y++ {
		for i := 0; i < padding; i++ {
			fmt.Print(" ")
		}
		for x := -diameter; x <= diameter; x++ {
			if x*x+4*y*y <= diameter*diameter {
				if isHead {
					if y == -1 { // eyes
						if x == -3 || x == -1 || x == 3 || x == 5 {
							fmt.Print(" ")
						} else if x == -2 || x == 4 {
							fmt.Print("@")
						} else {
							fmt.Print("*")
						}
					} else if y == 2 && x >= -1 && x <= 3 { // mouth
						fmt.Print("-")
					} else {
						fmt.Print("*")
					}
				} else {
					if (y == 0 || y == 3) && (x == -1 || x == 1) { // buttons
						fmt.Print("o")
					} else {
						fmt.Print("*")
					}
				}
			} else {
				if hasHands && y == 0 { // hands
					if x < -diameter {
						fmt.Print("/")
					} else if x > diameter {
						fmt.Print("\\")
					}
				} else {
					fmt.Print(" ")
				}
			}
		}
		if y != radius-bottomTrim {
			fmt.Println()
		}
	}
}

func printHat(padding int) {
	for i := 0; i < 7; i++ { // top of the hat
		for j := 0; j < padding-2; j++ {
			fmt.Print(" ")
		}
		fmt.Println("############################################")
	}

	for i := 0; i < padding-(padding/5); i++ {
		fmt.Print(" ")
	}
	fmt.Println("============================================================") // brim of the hat

}

func main() {
	fmt.Println()
	fmt.Println()
	printHat(50)                            // hat
	printCircle(10, 50, true, 7, 1, false)  // head with eyes, nose, and mouth
	printCircle(15, 40, false, 1, 4, true)  // middle with buttons and hands
	printCircle(20, 30, false, 3, 3, false) // bottom
}
