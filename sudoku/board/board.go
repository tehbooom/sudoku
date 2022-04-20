package board

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	cols = 9
	rows = 9
	n1   = "\n"
)

type Board struct {
	TopRight                string // ┓
	TopBoxSeparator         string // ┳
	TopColumnSeparator      string // ┯
	TopLeft                 string // ┏
	VerticalBoxSeparator    string // ┃
	VerticalColumnSeparator string // │
	BottomRight             string // ┛
	BottomBoxSeparator      string // ┻
	BottonColumnSeparator   string // ┷
	BottomLeft              string // ┗
	HorizontalBoxSeparator  string // ━━━
	HorizontalRowSeparator  string // ─
	LeftBoxSeparator        string // ┣
	LeftRowSeparator        string // ┠
	RightBoxSeparator       string // ┫
	RightRowSeparator       string // ┫
	MiddleBoxSeparator      string // ╋
	MiddleColumnSeparator   string // "┼"
	MiddleHybridSeparator   string // ┿
}

func (b Board) printBoard() {
	TopBox := strings.Repeat(b.HorizontalBoxSeparator+b.BottonColumnSeparator, 2) + b.HorizontalBoxSeparator + b.TopBoxSeparator
	TopBar := b.TopLeft + TopBox + b.TopRight + n1
	RowBox := strings.Repeat(b.HorizontalRowSeparator, 3)
	BoxHorizontal := b.LeftBoxSeparator + b.HorizontalBoxSeparator + b.MiddleBoxSeparator
	BottomBox := strings.Repeat(b.HorizontalBoxSeparator+b.TopColumnSeparator, 2) + b.HorizontalBoxSeparator + b.BottomBoxSeparator
	BottomBar := b.BottomLeft + BottomBox + b.BottomRight + n1
	fmt.Printf("%s", TopBar)
	for i := 0; i < rows; i++ {
		completerow := getNum(i) + n1
		rowSeparator := b.LeftRowSeparator + RowBox + b.VerticalBoxSeparator + n1
		if i == 3 {
			fmt.Printf(BoxHorizontal)
		}
		fmt.Printf(completerow)
		fmt.Printf(rowSeparator)
	}
	fmt.Printf("%s", BottomBar)
}

func (b Board) getNum(row int) string {
	var numbers []int
	i := row
	for j := 0; j < cols; j++ {
		num, err := state.game[i][j]
		numbers[j] = num
	}
	Box1 := " " + strconv.Itoa(numbers[0]) + " " + b.VerticalColumnSeparator + " " + strconv.Itoa(numbers[1]) + " " + b.VerticalColumnSeparator + " " + strconv.Itoa(numbers[2]) + " " + b.VerticalColumnSeparator
	Box2 := " " + strconv.Itoa(numbers[3]) + " " + b.VerticalColumnSeparator + " " + strconv.Itoa(numbers[4]) + " " + b.VerticalColumnSeparator + " " + strconv.Itoa(numbers[5]) + " " + b.VerticalColumnSeparator
	Box3 := " " + strconv.Itoa(numbers[6]) + " " + b.VerticalColumnSeparator + " " + strconv.Itoa(numbers[7]) + " " + b.VerticalColumnSeparator + " " + strconv.Itoa(numbers[8]) + " "
	completerow := b.VerticalBoxSeparator + Box1 + Box2 + Box3 + b.VerticalBoxSeparator
	return completerow
}
