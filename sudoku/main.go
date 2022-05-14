package main

import (
	"fmt"
	"os"

	_ "github.com/tehbooom/sudoku/board"

	"github.com/mbndr/figlet4go"
)

func main() {
	state := board.State{}
	ascii := figlet4go.NewAsciiRender()
	renderStr, _ := ascii.Render("Sudoku")
	fmt.Print(renderStr)

	// check if the player wants to play the game
	fmt.Printf("Would you like to play sudoku? (y/n)\n")
	var game string
	fmt.Scanln(&game)
	if game == "y" {
		board.generateEmpty()
	} else if game == "n" {
		fmt.Printf("Goodbye!\n")
		os.Exit(0)
	} else {
		fmt.Printf("Please enter y or n\n")
		os.Exit(0)
	}
	fmt.Printf("Generating Puzzle...\n")
	board.start()

	// start infinite game loop
	for {
		fmt.Print(renderStr)

		board.PrintBoard()
		fmt.Println("Please enter the row column and value to put on the board..\n")
		fmt.Println("For example, 1 2 9")
		for {
			var row, column, value int
			fmt.Scan(&row, &column, &value)
			valid := board.placeNumber(row, column, value)

			if valid == nil {
				break
			}
			fmt.Println(valid)
			fmt.Printf("Please re-enter a correct position and value...\n")
		}

		result = board.checkForWinner()
		if result == completeWrong {
			fmt.Println("Somethings not right check again..")
		}
		if result == winner {
			break
		}
		if result == incomplete {
			fmt.Println()
		}
	}

}
