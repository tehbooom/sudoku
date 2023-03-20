package main

import (
	"fmt"
	"os"

	"github.com/inancgumus/screen"

	"github.com/tehbooom/sudoku/game"
	_ "github.com/tehbooom/sudoku/style"
)

func main() {
	var s = game.State{}
	var difficulty int
	arg_len := len(os.Args[1:])
	if arg_len != 1 {
		fmt.Println(" Please add arguements of either easy, medium, hard, expert")
		os.Exit(1)
	} else if os.Args[1] == "easy" {
		difficulty = 20
	} else if os.Args[1] == "medium" {
		difficulty = 40
	} else if os.Args[1] == "hard" {
		difficulty = 60
	} else if os.Args[1] == "expert" {
		difficulty = 70
	} else {
		fmt.Println(" Please add arguements of either easy, medium, hard, expert")
		os.Exit(1)
	}

	s.Start(difficulty)

	// start infinite game loop
	for {
		screen.Clear()
		screen.MoveTopLeft()
		s.PrintBoard()
		fmt.Println("Please enter the row column and value to put on the board: 1 2 9")
		for {
			var row, column, value int
			fmt.Scan(&row, &column, &value)
			valid := s.PlaceNumber(row, column, value)

			if valid == nil {
				break
			}
			fmt.Println(valid)
		}

		result := s.CheckForWinner()
		if result == "completeWrong" {
			fmt.Println("Somethings not right check again..")
		}
		if result == "winner" {
			s.PrintBoard()
			fmt.Println("Congrats!! You won suGOku")
			break
		}
		if result == "incomplete" {
			fmt.Println()
		}
	}

}
