package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func printCircle(radius int, padding int, isHead bool, topTrim int, bottomTrim int, hasHands bool) {
	diameter := 2 * radius
	for y := -radius + topTrim; y <= radius-bottomTrim; y++ {
		fmt.Print("\n")
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

					if (y == -radius/3 || y == radius/3) && (x == -1 || x == 1) { // buttons
						fmt.Print("o")
					} else {
						if hasHands { // hands
							if y == -radius+topTrim+6 && x == -24 || y == -radius+topTrim+7 && x == -23 {
								fmt.Print("\\")
							} else if y == -radius+topTrim+6 && x == 24 || y == -radius+topTrim+7 && x == 23 {
								fmt.Print("/")
							} else {
								fmt.Print("*")
							}
						} else {
							fmt.Print("*")
						}
					}
				}
			} else {
				if hasHands { // hands
					if (y == -radius+topTrim && x == -30) || (y == -radius+topTrim+1 && x == -29) || (y == -radius+topTrim+2 && x == -28) || (y == -radius+topTrim+3 && x == -27) || (y == -radius+topTrim+4 && x == -26) || (y == -radius+topTrim+5 && x == -25) {
						fmt.Print("\\")
					} else if (y == -radius+topTrim && x == 30) || (y == -radius+topTrim+1 && x == 29) || (y == -radius+topTrim+2 && x == 28) || (y == -radius+topTrim+3 && x == 27) || (y == -radius+topTrim+4 && x == 26) || (y == -radius+topTrim+5 && x == 25) {
						fmt.Print("/")
					} else {
						fmt.Print(" ")
					}
				} else {
					fmt.Print(" ")
				}
			}
		}
	}
}

func printHat(padding int, width int) {
	for i := 0; i < 7; i++ { // top of the hat
		for j := 0; j < padding-1; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= width; k++ {
			fmt.Print("#")
		}
		fmt.Println()

	}

	for i := 0; i < padding-(padding/5); i++ {
		fmt.Print(" ")
	}
	fmt.Print("============================================================") // brim of the hat

}

func snowmanLoader(skipToQuestion bool) bool {
	if skipToQuestion != true {
		messages := []string{
			"Lowering temperature",
			"Generating snow",
			"Rolling snow balls",
			"Stacking snow balls",
			"Installing eyes",
			"Installing hands",
			"Installing hat",
			"Installing buttons",
			"Carving a charming smile",
		}

		for i := 0; i < 135; i++ {
			if i%15 == 0 {
				fmt.Print("\n" + messages[i/15])
			}
			fmt.Print(".")
			time.Sleep(500 * time.Millisecond)
		}

		time.Sleep(500 * time.Millisecond)

		fmt.Println("\nSnowman ready!")
		fmt.Println()

		time.Sleep(1000 * time.Millisecond)
	}

	reader := bufio.NewReader(os.Stdin)
	loaderSuccessful := false
	for {
		fmt.Print("Do you want to see the Snowman? (this action cannot be undone) (yes/no): ")
		text, _ := reader.ReadString('\n')
		// convert input to lower case
		text = strings.ToLower(text)
		text = strings.TrimSpace(text)

		if text == "yes" {
			fmt.Println("Ok, here it is..")
			time.Sleep(500 * time.Millisecond)
			loaderSuccessful = true
			break
		} else if text == "no" {
			fmt.Println("Uhh, what a waste of snow.")
			time.Sleep(500 * time.Millisecond)
			loaderSuccessful = false
			break
		} else {
			fmt.Println("Please answer with yes or no.")
		}
	}
	if loaderSuccessful {
		return true
	} else {
		return false
	}
}

func showSnowman() {
	fmt.Println()
	fmt.Println()
	printHat(50, 43)                        // hat
	printCircle(10, 50, true, 7, 2, false)  // head with eyes, nose, and mouth
	printCircle(15, 40, false, 1, 5, true)  // middle with buttons and hands
	printCircle(20, 30, false, 3, 3, false) // bottom with buttons
}

func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default: // For Linux or any other Unix system
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	fmt.Println()
	fmt.Println(".*********************************.")
	fmt.Println("|                                 |")
	fmt.Println("|   ELLO Snowman Builder v2.0.1   |")
	fmt.Println("|   Build date: 2023/12/12        |")
	fmt.Println("|                                 |")
	fmt.Println("'*********************************'")
	fmt.Println()
	fmt.Println("Initialising, please wait..")
	time.Sleep(2000 * time.Millisecond)
	//clearScreen()
	if snowmanLoader(false) {
		clearScreen()
		showSnowman()
	} else {
		fmt.Println("Good bye.")
	}
}
