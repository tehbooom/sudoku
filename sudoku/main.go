package main

import (
	"fmt"
	"os"

	"github.com/mbndr/figlet4go"
)

func main() {
	ascii := figlet4go.NewAsciiRender()
	renderStr, _ := ascii.Render("Sudoku")
	fmt.Print(renderStr)
	fmt.Printf("Would you like to play sudoku? (y/n)\n")

	var game string
	fmt.Scanln(&game)
	if game == "y" {

	} else if game == "n" {
		fmt.Printf("Goodbye!\n")
		os.Exit(0)
	} else {
		fmt.Printf("Please enter y or n\n")
		os.Exit(0)
	}

	var difficulty string
	fmt.Printf("What difficulty would you like to play? (easy/medium/hard)\n")
	fmt.Scanln(&difficulty)
	if difficulty == "easy" {
		fmt.Printf("Test\n")
	} else if difficulty == "medium" {
		fmt.Printf("Test\n")
	} else if difficulty == "hard" {
		fmt.Printf("Test\n")
	} else {
		fmt.Printf("Please enter easy medium or hard\n")
		os.Exit(0)
	}
}
