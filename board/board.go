package board

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"github.com/tehbooom/sudoku/style"
	"github.com/tehbooom/sudoku/game"
)

func PrintBoard() {
	topBox := strings.Repeat(style.horizontalBoxSeparator+style.topColumnSeparator, 2) + style.horizontalBoxSeparator + style.topBoxSeparator
	var top string = style.topLeft + strings.Repeat(topBox, 2) + strings.Repeat(style.horizontalBoxSeparator + style.topColumnSeparator, 2) + style.horizontalBoxSeparator + style.topRight + style.n1
	// ┏━━━┯━━━┯━━━┳━━━┯━━━┯━━━┳━━━┯━━━┯━━━┓

	bottomBox := strings.Repeat(style.horizontalBoxSeparator+style.bottomColumnSeparator, 2) + style.horizontalBoxSeparator + style.bottomBoxSeparator
	var bottom string = style.bottomLeft + strings.Repeat(bottomBox, 2) + strings.Repeat(style.horizontalBoxSeparator + style.bottomColumnSeparator, 2) + style.horizontalBoxSeparator + style.bottomRight + style.n1
	// ┗━━━┷━━━┷━━━┻━━━┷━━━┷━━━┻━━━┷━━━┷━━━┛


	rowColumn := strings.Repeat(style.horizontalRowSeparator + style.middleColumnSeparator, 2) + style.horizontalRowSeparator + style.verticalBoxSeparator
	rowSeparator := style.leftRowSeparator + strings.Repeat(rowColumn, 2) + strings.Repeat(style.horizontalRowSeparator + style.middleHybridSeparator, 2)  + style.horizontalRowSeparator + style.rightRowSeparator
	// ┠───┼───┼───╂───┼───┼───╂───┼───┼───┨

	rowBox := strings.Repeat(style.horizontalBoxSeparator + style.middleHybridSeparator, 2) + style.horizontalBoxSeparator + style.verticalBoxSeparator
	boxSeparator := style.leftBoxSeparator + strings.Repeat(rowBox, 2) + strings.Repeat(style.horizontalBoxSeparator + style.middleBoxSeparator, 2) + style.horizontalBoxSeparator + style.rightBoxSeparator
	// ┣━━━┿━━━┿━━━╋━━━┿━━━┿━━━╋━━━┿━━━┿━━━┫

	// Creates the middle
	// ┃ 1 │ 2 │ 3 ┃ 4 │ 5 │ 6 ┃ 7 │ 8 │ 9 ┃ * 9
	contents := make([]string, 17)
	for _, i := range contents{
		if i != 0 && i % 2 != 0 { // all odd numbers for contents except 0
			if i == 5 || 11 { // the Box separator only appears twice in board
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
		var row string = s.getNum(i)
		contents[0] = row
	}

	fmt.Println(top) // prints top
	printSlice(contents) // prints every element in slice contents line by line
	fmt.Println(bottom) // prints bottom

}

// grabs the number with the correpsonding position in the matrix from the game state
func (s *game.State) getNum(row int) string {
	var numbers []string
	i := row
	for j := 0; j < cols; j++ {
		if s.board[i][j] == 0 {
			numbers[j] == [" "]
		}
		numbers[j] = strconv.Itoa(s.board[i][j])
	}
	Box1 := " " + numbers[0] + " " + style.verticalColumnSeparator + " " + numbers[1] + " " + style.verticalColumnSeparator + " " + numbers[2] + " "
	Box2 := style.verticalBoxSeparator + " " + numbers[3] + " " + style.verticalColumnSeparator + " " + numbers[4] + " " + style.verticalColumnSeparator + " " + numbers[5] + " " 
	Box3 := style.verticalBoxSeparator + " " + numbers[6] + " " + style.verticalColumnSeparator + " " + numbers[7] + " " + style.verticalColumnSeparator + " " + numbers[8] + " "
	var completerow string = style.verticalBoxSeparator + Box1 + Box2 + Box3 + style.verticalBoxSeparator
	return completerow
}

// gets all elements in slice through recursion
func printSlice(s []int) {
    fmt.Println(s[0])
    printSlice(s[1:])
}
