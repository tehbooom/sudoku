package game

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tehbooom/sudoku/style"
)

// prints the sudoku board with all values from game
func (s *State) PrintBoard() {
	topBox := strings.Repeat(style.HorizontalBoxSeparator+style.TopColumnSeparator, 2) + style.HorizontalBoxSeparator + style.TopBoxSeparator
	var top string = style.TopLeft + strings.Repeat(topBox, 2) + strings.Repeat(style.HorizontalBoxSeparator+style.TopColumnSeparator, 2) + style.HorizontalBoxSeparator + style.TopRight
	// ┏━━━┯━━━┯━━━┳━━━┯━━━┯━━━┳━━━┯━━━┯━━━┓

	bottomBox := strings.Repeat(style.HorizontalBoxSeparator+style.BottomColumnSeparator, 2) + style.HorizontalBoxSeparator + style.BottomBoxSeparator
	var bottom string = style.BottomLeft + strings.Repeat(bottomBox, 2) + strings.Repeat(style.HorizontalBoxSeparator+style.BottomColumnSeparator, 2) + style.HorizontalBoxSeparator + style.BottomRight + style.N1
	// ┗━━━┷━━━┷━━━┻━━━┷━━━┷━━━┻━━━┷━━━┷━━━┛

	rowColumn := strings.Repeat(style.HorizontalRowSeparator+style.MiddleColumnSeparator, 2) + style.HorizontalRowSeparator + style.VerticalBoxColumnSeparator
	rowSeparator := style.LeftRowSeparator + strings.Repeat(rowColumn, 2) + strings.Repeat(style.HorizontalRowSeparator+style.MiddleColumnSeparator, 2) + style.HorizontalRowSeparator + style.RightRowSeparator
	// ┠───┼───┼───╂───┼───┼───╂───┼───┼───┨

	rowBox := strings.Repeat(style.HorizontalBoxSeparator+style.MiddleHybridSeparator, 2) + style.HorizontalBoxSeparator + style.MiddleBoxSeparator
	boxSeparator := style.LeftBoxSeparator + strings.Repeat(rowBox, 2) + strings.Repeat(style.HorizontalBoxSeparator+style.MiddleHybridSeparator, 2) + style.HorizontalBoxSeparator + style.RightBoxSeparator
	// ┣━━━┿━━━┿━━━╋━━━┿━━━┿━━━╋━━━┿━━━┿━━━┫

	// ┃ 1 │ 2 │ 3 ┃ 4 │ 5 │ 6 ┃ 7 │ 8 │ 9 ┃ * 9
	contents := make([]string, 17)
	for i := range contents {
		if i != 0 && i%2 != 0 { // all odd numbers for contents except 0
			if i == 5 || i == 11 { // the Box separator only appears twice in board
				contents[i] = boxSeparator
			} else {
				contents[i] = rowSeparator
			}
		}
	}
	for i := 0; i < Size; i++ {
		if i == 0 {
			var row string = s.getNum(i)
			contents[0] = row
		} else {
			newIndex := i * 2 // index is now multplied by 2 because contents is every other index
			var row string = s.getNum(i)
			contents[newIndex] = row
		}
	}

	var str string
	for index, i := range contents {
		if index == 0 {
			str += fmt.Sprintf("%s", i)
		} else {
			str += fmt.Sprintf("\n%s", i)
		}
	}
	fmt.Println(top)    // prints top
	fmt.Println(str)    // prints every element in slice contents line by line
	fmt.Println(bottom) // prints bottom

}

// grabs the number with the correpsonding position in the matrix from the game state
func (s *State) getNum(row int) string {
	var numbers [9]string
	for j := 0; j < Size; j++ {
		number := s.Board[row][j]
		numberString := strconv.Itoa(number)
		if numberString == "0" {
			numbers[j] = " "
		} else {
			numbers[j] = numberString
		}
	}
	fmt.Println()
	Box1 := " " + numbers[0] + " " + style.VerticalColumnSeparator + " " + numbers[1] + " " + style.VerticalColumnSeparator + " " + numbers[2] + " "
	Box2 := style.VerticalBoxSeparator + " " + numbers[3] + " " + style.VerticalColumnSeparator + " " + numbers[4] + " " + style.VerticalColumnSeparator + " " + numbers[5] + " "
	Box3 := style.VerticalBoxSeparator + " " + numbers[6] + " " + style.VerticalColumnSeparator + " " + numbers[7] + " " + style.VerticalColumnSeparator + " " + numbers[8] + " "
	var completerow string = style.VerticalBoxSeparator + Box1 + Box2 + Box3 + style.VerticalBoxSeparator
	return completerow
}
