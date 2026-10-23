package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintLine() {
	fmt.Println("--------------------------------------------------")
}

func PrintHeader(header string) {
	fmt.Println("--------------------------------------------------")
	color.Green(header)
	fmt.Println("--------------------------------------------------")
}

func PrintErrorMessage(message string) {
	fmt.Println("--------------------------------------------------")
	color.Red("Error: %s\n", message)
	fmt.Println("--------------------------------------------------")
}
