package board

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tehbooom/sudoku/game"
	_ "github.com/tehbooom/sudoku/style"
)

const (
	cols                    = 9
	rows                    = 9
	n1                      = "\n"
	topRight                = "┓"
	topBoxSeparator         = "┳"
	topColumnSeparator      = "┯"
	topLeft                 = "┏"
	verticalBoxSeparator    = "┃"
	verticalColumnSeparator = "│"
	bottomRight             = "┛"
	bottomBoxSeparator      = "┻"
	bottomColumnSeparator   = "┷"
	bottomLeft              = "┗"
	horizontalBoxSeparator  = "━━━"
	horizontalRowSeparator  = "───"
	leftBoxSeparator        = "┣"
	leftRowSeparator        = "┠"
	rightBoxSeparator       = "┫"
	rightRowSeparator       = "┫"
	middleBoxSeparator      = "╋"
	middleColumnSeparator   = "┼"
	middleHybridSeparator   = "┿"
)

// prints the sudoku board with all values from game
func PrintBoard() {
	topBox := strings.Repeat(horizontalBoxSeparator+topColumnSeparator, 2) + horizontalBoxSeparator + topBoxSeparator
	var top string = topLeft + strings.Repeat(topBox, 2) + strings.Repeat(horizontalBoxSeparator+topColumnSeparator, 2) + horizontalBoxSeparator + topRight + n1
	// ┏━━━┯━━━┯━━━┳━━━┯━━━┯━━━┳━━━┯━━━┯━━━┓

	bottomBox := strings.Repeat(horizontalBoxSeparator+bottomColumnSeparator, 2) + horizontalBoxSeparator + bottomBoxSeparator
	var bottom string = bottomLeft + strings.Repeat(bottomBox, 2) + strings.Repeat(horizontalBoxSeparator+bottomColumnSeparator, 2) + horizontalBoxSeparator + bottomRight + n1
	// ┗━━━┷━━━┷━━━┻━━━┷━━━┷━━━┻━━━┷━━━┷━━━┛

	rowColumn := strings.Repeat(horizontalRowSeparator+middleColumnSeparator, 2) + horizontalRowSeparator + verticalBoxSeparator
	rowSeparator := leftRowSeparator + strings.Repeat(rowColumn, 2) + strings.Repeat(horizontalRowSeparator+middleHybridSeparator, 2) + horizontalRowSeparator + rightRowSeparator
	// ┠───┼───┼───╂───┼───┼───╂───┼───┼───┨

	rowBox := strings.Repeat(horizontalBoxSeparator+middleHybridSeparator, 2) + horizontalBoxSeparator + verticalBoxSeparator
	boxSeparator := leftBoxSeparator + strings.Repeat(rowBox, 2) + strings.Repeat(horizontalBoxSeparator+middleBoxSeparator, 2) + horizontalBoxSeparator + rightBoxSeparator
	// ┣━━━┿━━━┿━━━╋━━━┿━━━┿━━━╋━━━┿━━━┿━━━┫

	// Creates the middle
	// ┃ 1 │ 2 │ 3 ┃ 4 │ 5 │ 6 ┃ 7 │ 8 │ 9 ┃ * 9
	contents := make([]string, 17)
	for i := range contents {
		if i != 0 && i%2 != 0 { // all odd numbers for contents except 0
			if i == 5 || i == 11 { // the Box separator only appears twice in board
				contents[i] = boxSeparator
			}
			contents[i] = rowSeparator
		}
	}
	for i := 0; i < rows; i++ {
		if i > 0 { // because they are the same index
			newIndex := i * 2 // index is now multplied by 2 because contents is every other index
			var row string = getNum(i)
			contents[newIndex] = row
		}
		var row string = getNum(i)
		contents[0] = row
	}

	fmt.Println(top)      // prints top
	fmt.Println(contents) // prints every element in slice contents line by line
	fmt.Println(bottom)   // prints bottom

}

// grabs the number with the correpsonding position in the matrix from the game state
func getNum(row int) string {
	var s = game.State{}
	var numbers []string
	i := row
	for j := 0; j < cols; j++ {
		if s.Board[i][j] == 0 {
			numbers[j] = " "
		}
		numbers[j] = strconv.Itoa(s.Board[i][j])
	}
	Box1 := " " + numbers[0] + " " + verticalColumnSeparator + " " + numbers[1] + " " + verticalColumnSeparator + " " + numbers[2] + " "
	Box2 := verticalBoxSeparator + " " + numbers[3] + " " + verticalColumnSeparator + " " + numbers[4] + " " + verticalColumnSeparator + " " + numbers[5] + " "
	Box3 := verticalBoxSeparator + " " + numbers[6] + " " + verticalColumnSeparator + " " + numbers[7] + " " + verticalColumnSeparator + " " + numbers[8] + " "
	var completerow string = verticalBoxSeparator + Box1 + Box2 + Box3 + verticalBoxSeparator
	return completerow
}
