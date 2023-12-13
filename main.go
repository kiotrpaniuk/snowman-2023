package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
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
						//if hasHands { // hands
						//	if y == -radius+topTrim+6 && x == -diameter+5 || y == -radius+topTrim+7 && x == -diameter+6 {
						//		fmt.Print("\\")
						//	} else if y == -radius+topTrim+6 && x == 24 || y == -radius+topTrim+7 && x == 23 {
						//		fmt.Print("/")
						//	} else {
						//		fmt.Print("*")
						//	}
						//} else {
						fmt.Print("*")
						//}
					}
				}
			} else {
				if hasHands { // hands
					if y == -radius+topTrim && x == -diameter || y == -radius+topTrim+1 && x == -diameter+1 || y == -radius+topTrim+2 && x == -diameter+2 || y == -radius+topTrim+3 && x == -diameter+3 || y == -radius+topTrim+4 && x == -diameter+4 {
						fmt.Print("\\")
					} else if y == -radius+topTrim && x == diameter || y == -radius+topTrim+1 && x == diameter-1 || y == -radius+topTrim+2 && x == diameter-2 || y == -radius+topTrim+3 && x == diameter-3 || y == -radius+topTrim+4 && x == diameter-4 {
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
	brimWidth := width + 10
	hatCenterPosition := padding + (width / 2)
	brimPadding := hatCenterPosition - (brimWidth / 2)
	//fmt.Println(brimWidth, hatCenterPosition, brimPadding)

	for i := 0; i < 7; i++ { // top of the hat
		for j := 0; j < padding-1; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= width; k++ {
			fmt.Print("#")
		}
		fmt.Println()

	}

	for i := 0; i < brimPadding; i++ {
		fmt.Print(" ")
	}
	for i := 0; i < brimWidth; i++ {
		fmt.Print("=") // brim of the hat
	}

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

	}

	reader := bufio.NewReader(os.Stdin)
	loaderSuccessful := false
	for {
		time.Sleep(500 * time.Millisecond)

		fmt.Println("\nSnowman ready!")
		fmt.Println()

		time.Sleep(1000 * time.Millisecond)
		fmt.Print("Do you want to see the Snowman? (this action cannot be undone) (yes/no): ")
		text, _ := reader.ReadString('\n')
		// convert input to lower case
		text = strings.ToLower(text)
		text = strings.TrimSpace(text)

		if text == "yes" {
			fmt.Println("Ok, here it is!")
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

func showSnowman(windowWidth int, windowHeight int) {

	// calculate scaling factor
	var scalingFactor float32
	if windowHeight > 80 {
		scalingFactor = 1
	} else {
		scalingFactor = float32(windowHeight) / 100
	}

	//fmt.Println("scalingFactor:", scalingFactor)

	fmt.Println()
	fmt.Println()

	// calculate Snowman dimensions based on window height
	// when height is around 85-90 then the following dimensions work well

	var snowmanDimensions = map[string]int{
		//"widthHat": 43,
		//"paddingHat":  50,
		"radiusHead":  10,
		"radiusTorso": 15,
		"radiusBase":  20,
		//"paddingHead":     50,
		//"paddingTorso":    40,
		"paddingBase":     30,
		"topTrimHead":     7,
		"topTrimTorso":    3,
		"topTrimBase":     3,
		"bottomTrimHead":  3,
		"bottomTrimTorso": 5,
		"bottomTrimBase":  3,
	}

	// Randomise the dimensions a bit
	rand.Seed(time.Now().UnixNano())
	for key, value := range snowmanDimensions {
		if key == "radiusHead" || key == "radiusTorso" || key == "radiusBase" || key == "topTrimHead" || key == "bottomTrimHead" || key == "topTrimTorso" || key == "bottomTrimTorso" || key == "topTrimBase" || key == "bottomTrimBase" {
			// Calculate the 10% margin
			margin := int(float64(value) * 0.11)
			// Generate a random number within the range [value - margin, value + margin]
			randomValue := value - margin + rand.Intn(2*margin+1)
			// Update the map
			snowmanDimensions[key] = randomValue
		}
	}

	snowmanDimensions["widthHat"] = snowmanDimensions["radiusHead"] * 4

	centerPosition := snowmanDimensions["paddingBase"] + snowmanDimensions["radiusBase"]
	snowmanDimensions["paddingHead"] = centerPosition - snowmanDimensions["radiusHead"] + 10
	snowmanDimensions["paddingTorso"] = centerPosition - (snowmanDimensions["radiusTorso"]) + 5
	snowmanDimensions["paddingHat"] = centerPosition - (snowmanDimensions["widthHat/2"]) + 1

	// apply scaling to all snowmanDimensions
	for key, value := range snowmanDimensions {
		//if key == "radiusHead" || key == "radiusTorso" || key == "radiusBase" || key == "topTrimHead" || key == "bottomTrimHead" || key == "topTrimTorso" || key == "bottomTrimTorso" || key == "topTrimBase" || key == "bottomTrimBase" {
		snowmanDimensions[key] = int(scalingFactor * float32(value))
		//}
	}

	//fmt.Println(snowmanDimensions)

	// output the Snowman
	printHat(snowmanDimensions["paddingHat"], snowmanDimensions["widthHat"])                                                                                               // hat
	printCircle(snowmanDimensions["radiusHead"], snowmanDimensions["paddingHead"], true, snowmanDimensions["topTrimHead"], snowmanDimensions["bottomTrimHead"], false)     // head with eyes, nose, and mouth
	printCircle(snowmanDimensions["radiusTorso"], snowmanDimensions["paddingTorso"], false, snowmanDimensions["topTrimTorso"], snowmanDimensions["bottomTrimTorso"], true) // middle with buttons and hands
	printCircle(snowmanDimensions["radiusBase"], snowmanDimensions["paddingBase"], false, snowmanDimensions["topTrimBase"], snowmanDimensions["bottomTrimBase"], false)    // bottom with buttons
	fmt.Println()
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

func getConsoleSize() (int, int, error) {
	if runtime.GOOS == "windows" {
		return getConsoleSizeWindows()
	}
	return getConsoleSizeUnix()
}

func getConsoleSizeUnix() (int, int, error) {
	return terminal.GetSize(int(os.Stdin.Fd()))
}

func getConsoleSizeWindows() (int, int, error) {
	cmd := exec.Command("powershell", "-Command", "$Host.UI.RawUI.WindowSize.Height")
	output, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	height, err := strconv.Atoi(strings.TrimSpace(string(output)))
	if err != nil {
		return 0, 0, err
	}

	cmd = exec.Command("powershell", "-Command", "$Host.UI.RawUI.WindowSize.Width")
	output, err = cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	width, err := strconv.Atoi(strings.TrimSpace(string(output)))
	if err != nil {
		return 0, 0, err
	}

	return width, height, nil
}

func main() {

	fastPtr := flag.Bool("fast", false, "a bool")
	flag.Parse()
	fmt.Println()
	fmt.Println(".***********************************************************************.")
	fmt.Println("|                                                                       |")
	fmt.Println("|    ELLO Snowman Builder v2.0.2                                        |")
	fmt.Println("|    Build date: 2023/12/13                                             |")
	fmt.Println("|                                                                       |")
	fmt.Println("|    Release notes:                                                     |")
	fmt.Println("|    * Every Snowman is unique, dimensions are randomised               |")
	fmt.Println("|    * Quality of the Snowman depends on size of the terminal window    |")
	fmt.Println("|    * Use --fast for fast mode (WARNING: Uses more snow)               |")
	fmt.Println("|                                                                       |")
	fmt.Println("'***********************************************************************'")
	fmt.Println()
	fmt.Println("Initialising, please wait")
	if *fastPtr {
		fmt.Println("Fast mode is enabled")
	}
	//time.Sleep(2000 * time.Millisecond)

	windowWidth, windowHeight, err := getConsoleSize()
	if err != nil {
		fmt.Println("Error getting size:", err)
	} else {
		fmt.Printf("Window dimensions: %dx%d\n", windowWidth, windowHeight)
	}

	//clearScreen()
	if snowmanLoader(*fastPtr) {
		clearScreen()
		showSnowman(windowWidth, windowHeight)
		fmt.Println()
	} else {
		fmt.Println("Good bye.")
	}
}
