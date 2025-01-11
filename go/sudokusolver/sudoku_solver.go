package main

import (
	"fmt"
	"reflect"
)

const boardSize = 9

type board struct {
	board [boardSize][boardSize]int

	usedInRow, usedInCol, usedInSq [boardSize][boardSize + 1]bool
}

func newBoard(b [][]byte) board {
	var r board
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b[i][j] != '.' {
				v := int(b[i][j] - '0')
				r.set(i, j, v)
			}
		}
	}
	return r
}

func squareIndex(r, c int) int {
	return r/3*3 + c/3
}

func (b *board) canUse(r, c int, v int) bool {
	return !b.usedInRow[r][v] && !b.usedInCol[c][v] && !b.usedInSq[squareIndex(r, c)][v]
}

func (b *board) set(r, c int, v int) {
	b.board[r][c] = v
	b.usedInRow[r][v] = true
	b.usedInCol[c][v] = true
	b.usedInSq[squareIndex(r, c)][v] = true
}

func (b *board) reset(r, c int) {
	v := b.board[r][c]
	b.board[r][c] = 0
	b.usedInRow[r][v] = false
	b.usedInCol[c][v] = false
	b.usedInSq[squareIndex(r, c)][v] = false
}

func (b *board) isSet(r, c int) bool {
	return b.board[r][c] != 0
}

func (b *board) save(raw [][]byte) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			raw[i][j] = byte(b.board[i][j]) + '0'
		}
	}
}

func solve(b *board, r, c int) bool {
	if r == boardSize {
		return true
	}

	if c == boardSize {
		return solve(b, r+1, 0)
	}

	if b.isSet(r, c) {
		return solve(b, r, c+1)
	}

	for i := 1; i <= 9; i++ {
		if b.canUse(r, c, i) {
			b.set(r, c, i)
			if solve(b, r, c+1) {
				return true
			}
			b.reset(r, c)
		}
	}

	return false
}

func solveSudoku(rawBoard [][]byte) {
	board := newBoard(rawBoard)
	solve(&board, 0, 0)
	board.save(rawBoard)
}

func main() {
	inputs := [][][]byte{
		{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		},
	}
	outputs := [][][]byte{
		{
			{'5', '3', '4', '6', '7', '8', '9', '1', '2'}, {'6', '7', '2', '1', '9', '5', '3', '4', '8'}, {'1', '9', '8', '3', '4', '2', '5', '6', '7'},
			{'8', '5', '9', '7', '6', '1', '4', '2', '3'}, {'4', '2', '6', '8', '5', '3', '7', '9', '1'}, {'7', '1', '3', '9', '2', '4', '8', '5', '6'},
			{'9', '6', '1', '5', '3', '7', '2', '8', '4'}, {'2', '8', '7', '4', '1', '9', '6', '3', '5'}, {'3', '4', '5', '2', '8', '6', '1', '7', '9'},
		},
	}
	for i := 0; i < len(inputs); i++ {
		solveSudoku(inputs[i])
		if reflect.DeepEqual(inputs[i], outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, inputs[i], outputs[i])
		}
	}
}
