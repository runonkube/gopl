package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	lc "github.com/linehk/gopl/ch2/lengthconv"
	"github.com/linehk/gopl/ch2/tempconv"
	wc "github.com/linehk/gopl/ch2/weightconv"
)

func main() {
	cmdArgs := os.Args

	if len(cmdArgs) > 1 {
		for _, arg := range cmdArgs[1:] {
			if strings.Contains(arg, ",") {
				for _, num := range strings.Split(arg, ",") {
					if strings.TrimSpace(num) != "" {
						printConversions(num)
					}
				}
			} else {
				printConversions(arg)
			}
		}
	} else {
		fmt.Print("Enter numbers to convert separated by commas (e.g. 1,2,4):")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		for _, numStr := range strings.Split(input, ",") {
			printConversions(strings.TrimSpace(numStr))
		}
	}
}

func printConversions(numStr string) {
	if number, err := strconv.ParseFloat(numStr, 64); err != nil {
		fmt.Printf("Error converting '%s' to a number. Details: '%s'\n", numStr, err)
		return
	} else {
		doTempConversions(number)
		fmt.Println()
		doLengthConversions(number)
		fmt.Println()
		doWeightConversions(number)
		fmt.Println()
	}
}

func doTempConversions(number float64) {
	c := tempconv.Celsius(number)
	f := tempconv.Fahrenheit(number)
	fahrenheit := tempconv.CToF(c)
	celsius := tempconv.FToC(f)

	fmt.Printf("%s = %s\n", c, fahrenheit)
	fmt.Printf("%s = %s\n", f, celsius)
}

func doLengthConversions(number float64) {
	feet := lc.Feet(number)
	meter := lc.Meter(number)
	fmt.Printf("%s = %s\n", feet, lc.ToMeters(feet))
	fmt.Printf("%s = %s\n", meter, lc.ToFeet(meter))
}

func doWeightConversions(number float64) {
	kilo := wc.Kilo(number)
	pounds := wc.Pound(number)
	fmt.Printf("%s = %s\n", kilo, wc.ToPounds(kilo))
	fmt.Printf("%s = %s\n", pounds, wc.ToKilos(pounds))
}
