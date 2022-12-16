package main

import (
	"fmt"

	_ "github.com/tehbooom/sudoku/board"
	_ "github.com/tehbooom/sudoku/game"

	"github.com/mbndr/figlet4go"
)

func main() {

	// generates the sudoku banner
	// 				     #         #
	//  mmm   m   m   mmm#   mmm   #   m  m   m
	// #   "  #   #  #" "#  #" "#  # m"   #   #
	//  """m  #   #  #   #  #   #  #"#    #   #
	// "mmm"  "mm"#  "#m##  "#m#"  #  "m  "mm"#
	ascii := figlet4go.NewAsciiRender()
	renderStr, _ := ascii.Render("suGOku")
	fmt.Print(renderStr)

	game.Start()

	// start infinite game loop
	for {

		board.PrintBoard()
		fmt.Println("Please enter the row column and value to put on the board..\n")
		fmt.Println("For example, r1 c2 9")
		for {
			var row, column, value int
			fmt.Scan(&row, &column, &value)
			valid := game.placeNumber(row, column, value)

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
