package board

import (
	"strings"
)

const (
	Cols = 9
	Rows = 9
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
	LeftBoxSeparator        string // ┠
	RightBoxSeparator       string // ┫
	MiddleBoxSeparator      string // ╋
	MiddleColumnSeparator   string // "┼"
	MiddleHybridSeparator   string // ┿
}

func (b Board) toString() {
	TopBox := strings.Repeat(b.HorizontalBoxSeparator+b.BottonColumnSeparator, 2) + b.HorizontalBoxSeparator + b.TopBoxSeparator
	TopBar := b.TopLeft + TopBox + b.TopRight
	BottomBox := strings.Repeat(b.HorizontalBoxSeparator+b.TopColumnSeparator, 2) + b.HorizontalBoxSeparator + b.BottomBoxSeparator
	BottomBar := b.BottomLeft + BottomBox + b.BottomRight

	// run for loop on NumberBox to input xy of matrix if xy is null then ""
	// Number := "" + "xy" + ""

	// NumberRow := b.VerticalBoxSeparator + NumberBox  + b.VerticalBoxSeparator
	// SeparatorRow :=
}

// for i := 0; i < boardSize; i++ {
// 	for j := 0; j < boardSize; j++ {
// 		num := state.board[i][j]
// 		if num == 0 {
// 			fmt.Print("    ")
// 		} else {
// 			if enableColor {
// 				fmt.Printf("\033[3%sm%4d\033[0m", colorPalette[int(math.Log2(float64(num)))], num)
// 			} else {
// 				fmt.Printf("%4d", num)
// 			}
// 		}
// 		if j != boardSize-1 {
// 			fmt.Print("|")
// 		}
// 	}
// 	fmt.Print("\n")
// 	if i != boardSize-1 {
// 		for j := 0; j < boardSize; j++ {
// 			if j != boardSize-1 {
// 				fmt.Print("-----")
// 			} else {
// 				fmt.Print("----")
// 			}
// 		}
// 		fmt.Print("\n")
// 	}
// }
