package board

import (
	"bufio"
	"fmt"
	"math/rand/v2"
)

type Board struct {
	board    [][]uint64
	row, col int
}

func New(row, col int) Board {

	values := make([][]uint64, row)

	for i := range values {
		values[i] = make([]uint64, col)
	}

	return Board{
		board: values,
		row:   row, col: col,
	}
}

func (self *Board) Print(w *bufio.Writer) {

	defer w.Flush()

	for _, row := range self.board {

		for _, value := range row {
			fmt.Fprintf(w, "%5d ", value)
		}

		w.WriteRune('\n')
	}
}

func (self *Board) SpawnNewElement() {

	board := self.board

	for {

		rowIndex := rand.N(self.row)
		colIndex := rand.N(self.col)

		if board[rowIndex][colIndex] != 0 {
			continue
		}

		updatedValue := uint64(rand.N(1)+1) * 2
		board[rowIndex][colIndex] = updatedValue

		return
	}
}
