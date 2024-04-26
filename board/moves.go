package board

func (self *Board) MoveUp() (updated bool) {

	updated = false

	for i := range self.col {
		if self.shiftUpCol(i) {
			updated = true
		}
	}

	return updated
}

func (self *Board) MoveDown() (updated bool) {

	updated = false

	for i := range self.col {
		if self.shiftDownCol(i) {
			updated = true
		}
	}

	return updated
}

func (self *Board) MoveLeft() (updated bool) {

	updated = false

	for i := range self.row {
		if self.shiftLeftRow(i) {
			updated = true
		}
	}

	return updated
}

func (self *Board) MoveRight() (updated bool) {

	updated = false

	for i := range self.row {
		if self.shiftRightRow(i) {
			updated = true
		}
	}

	return updated
}

func (self *Board) shiftUpCol(colIndex int) (updated bool) {

	updated = false
	board := self.board

	for i := range self.row {

		if board[i][colIndex] == 0 {
			continue
		}

		for j := i + 1; j < self.row; j++ {

			if board[i][colIndex] == board[j][colIndex] {

				board[i][colIndex] *= 2
				board[j][colIndex] = 0

				updated = true

				break
			}

			if board[j][colIndex] != 0 {
				break
			}
		}
	}

	for i := range self.row {

		if board[i][colIndex] != 0 {
			continue
		}

		for j := i + 1; j < self.row; j++ {

			if board[j][colIndex] != 0 {

				board[i][colIndex] = board[j][colIndex]
				board[j][colIndex] = 0

				updated = true

				break
			}
		}
	}

	return updated
}

func (self *Board) shiftDownCol(colIndex int) (updated bool) {

	updated = false
	board := self.board

	for i := self.row - 1; i >= 0; i-- {

		if board[i][colIndex] == 0 {
			continue
		}

		for j := i - 1; j >= 0; j-- {

			if board[i][colIndex] == board[j][colIndex] {

				board[i][colIndex] *= 2
				board[j][colIndex] = 0

				updated = true

				break
			}

			if board[j][colIndex] != 0 {
				break
			}
		}
	}

	for i := self.row - 1; i >= 0; i-- {

		if board[i][colIndex] != 0 {
			continue
		}

		for j := i - 1; j >= 0; j-- {

			if board[j][colIndex] != 0 {

				board[i][colIndex] = board[j][colIndex]
				board[j][colIndex] = 0

				updated = true

				break
			}
		}
	}

	return updated

}

func (self *Board) shiftLeftRow(rowIndex int) (updated bool) {

	updated = false
	targetRow := self.board[rowIndex]

	for i := range targetRow {

		if targetRow[i] == 0 {
			continue
		}

		for j := i + 1; j < self.col; j++ {

			if targetRow[i] == targetRow[j] {

				targetRow[i] *= 2
				targetRow[j] = 0

				updated = true

				break
			}

			if targetRow[j] != 0 {
				break
			}
		}
	}

	for i := range targetRow {

		if targetRow[i] != 0 {
			continue
		}

		for j := i + 1; j < self.row; j++ {

			if targetRow[j] != 0 {
				targetRow[i] = targetRow[j]
				targetRow[j] = 0
				break
			}
		}
	}

	return updated
}

func (self *Board) shiftRightRow(rowIndex int) (updated bool) {

	updated = false
	targetRow := self.board[rowIndex]

	for i := self.col - 1; i >= 0; i-- {

		if targetRow[i] == 0 {
			continue
		}

		for j := i - 1; j >= 0; j-- {

			if targetRow[i] == targetRow[j] {

				targetRow[i] *= 2
				targetRow[j] = 0

				updated = true

				break
			}

			if targetRow[j] != 0 {
				break
			}
		}
	}

	for i := self.col - 1; i >= 0; i-- {

		if targetRow[i] != 0 {
			continue
		}

		for j := i - 1; j >= 0; j-- {

			if targetRow[j] != 0 {
				targetRow[i] = targetRow[j]
				targetRow[j] = 0
				break
			}
		}
	}

	return updated
}
