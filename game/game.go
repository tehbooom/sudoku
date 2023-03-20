package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	Size = 9
	sqrt = 3
)

type State struct {
	Board       [Size][Size]int
	SolvedBoard [Size][Size]int
}

var dictOriginalValues []int

func (s *State) Start(difficulty int) {
	s.fillDiagnol()
	s.fillRemaining(0, 3)
	s.SolvedBoard = s.Board
	s.removeCells(difficulty)
	s.setOriginalValue()
}

func generateNum() int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(9) + 1
	return num
}

func (s *State) fillDiagnol() {
	for i := 0; i < Size; i = i + 3 {
		s.fillBox(i, i)
	}
}

func (s *State) isValid(row, col, num int) bool {
	for i := 0; i < Size; i++ {
		if s.Board[row][i] == num {
			return false
		}
	}

	for i := 0; i < Size; i++ {
		if s.Board[i][col] == num {
			return false
		}
	}

	boxRow := row - row%3
	boxCol := col - col%3
	for i := boxRow; i < boxRow+3; i++ {
		for j := boxCol; j < boxCol+3; j++ {
			if s.Board[i][j] == num {
				return false
			}
		}
	}

	return true
}

func (s *State) fillBox(row, col int) {
	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			for true {
				num := generateNum()
				if s.isValid(row, col, num) {
					s.Board[i][j] = num
					break
				}
			}
		}
	}
}

func (s *State) fillRemaining(i, j int) bool {
	if i == Size-1 && j == Size {
		return true
	}

	if j == Size {
		i += 1
		j = 0
	}

	if s.Board[i][j] != 0 {
		return s.fillRemaining(i, j+1)
	}

	for num := 1; num <= Size; num++ {
		if s.isValid(i, j, num) {
			s.Board[i][j] = num
			if s.fillRemaining(i, j+1) {
				return true
			}
			s.Board[i][j] = 0
		}
	}
	return false
}

func (s *State) removeCells(value int) {
	count := 0
	for count != value {
		i := rand.Intn(9)
		j := rand.Intn(9)
		if s.Board[i][j] != 0 {
			s.Board[i][j] = 0
			count++
		}
	}
}

func (s *State) setOriginalValue() {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if s.Board[i][j] != 0 {

				value, err := strconv.Atoi(fmt.Sprintf("%d%d", i, j))
				if err != nil {
					panic(err)
				}
				dictOriginalValues = append(dictOriginalValues, value)
			}
		}
	}
}

func (s *State) CheckForWinner() string {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if s.Board[i][j] == 0 {
				return "incomplete"
			}
		}
	}
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if s.Board[i][j] != s.SolvedBoard[i][j] {
				return "completeWrong"
			}
		}
	}
	return "winner"
}
