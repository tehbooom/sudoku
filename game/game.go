package game

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	cols = 9
	rows = 9
	size = 9
)

type State struct {
	Board [cols][rows]int
}

// type positionOutOfBound struct {
// 	row    int
// 	column int
// }

// type valueNotValid struct {
// 	value int
// }

// type originalValue struct {
// 	value int
// }

// func (e *positionOutOfBound) Error() string {
// 	return fmt.Sprintf("posistion (%d,%d) is out of bounds...", e.row, e.column)
// }

// func (e *valueNotValid) Error() string {
// 	return fmt.Sprintf("value %d is not 1-9", e.value)
// }

// func (e *originalValue) Error() string {
// 	return fmt.Sprintf("cannot modify an original value generated by the puzzle...", e.value)
// }

func (s *State) Start() {
	s.generateEmpty()
	s.Board[0][0] = generateNum()
	s.generatePuzzle()
	fmt.Println(s.Board)
}

func (s *State) generateEmpty() {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			s.Board[i][j] = 0
		}
	}
}

func generateNum() int {
	number_list := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().UnixNano())
	rlength := len(number_list)
	randomIndex := rand.Intn(rlength - 1)
	num := number_list[randomIndex]
	return num
}

func (s *State) generatePuzzle() {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if s.Board[i][j] == 0 {
				for k := 1; k <= size; k++ {
					if s.isValid(i, j, k) {
						s.Board[i][j] = k
						s.generatePuzzle()
					}
				}
			}
		}
	}
}

func (s *State) isValid(row, col, num int) bool {
	// Check if num is already in the same row
	for i := 0; i < size; i++ {
		if s.Board[row][i] == num {
			return false
		}
	}

	// Check if num is already in the same column
	for i := 0; i < size; i++ {
		if s.Board[i][col] == num {
			return false
		}
	}

	// Check if num is already in the same 3x3 box
	boxRow := row - row%3
	boxCol := col - col%3
	for i := boxRow; i < boxRow+3; i++ {
		for j := boxCol; j < boxCol+3; j++ {

			if s.Board[i][j] == num {
				return false
			}
		}
	}

	// If num is not already in the same row, column, or box, it is valid
	return true
}

// func (s *State) placeNumber(row int, column int, value int) error {
// 	if row < 1 || column < 0 || row >= 10 || column >= 10 {
// 		return &positionOutOfBound{row, column}
// 	}
// 	if value < 1 || value >= 10 {
// 		return &valueNotValid{value}
// 	}
// 	// if s.Board[row][column] == FUNCTION THAT WILL BE SET AFTER WE REMOVE NUMBER {
// 	// 	return &originalValue{value}
// 	// }
// 	return nil
// }

// func (s *State) checkForWinner() string {
// 	// figure out a better way to check for winner
// }
