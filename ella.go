package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convert(hex string) (int64, error) {
	return strconv.ParseInt(hex, 16, 64)

}

func good(bin string) (int64, error) {
	return strconv.ParseInt(bin, 2, 64)
}

func ella(dec int64) string {
	return strconv.FormatInt(dec, 16)
}

func milo(dec int64) string {
	return strconv.FormatInt(dec, 2)
}
func main() {
	for {
		var input string
		fmt.Println("Enter input:")
		fmt.Scanln(&input)

		var operation string
		fmt.Println("hex", "bin", "dec", "Quit")
		fmt.Scanln(&operation)
		switch operation {
		case "hex":
			fmt.Println(convert(input))

		case "bin":
			fmt.Println(good(input))
		case "dec":
			i, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				fmt.Println("invalid decimal")
			}
			fmt.Println(strings.ToUpper(ella(i)))

			i2, _ := strconv.ParseInt(input, 10, 64)
			fmt.Println(milo(i2))
		case "Quit":
			fmt.Println("Goodbye")
			return
		}
	}

}
