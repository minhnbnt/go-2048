package main

import (
	"bufio"
	"go-2048/board"
	"os"
)

func main() {

	stdin := bufio.NewReader(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	board := board.New(4, 4)
	board.SpawnNewElement()

	for {

		board.Print(stdout)

		input, _, _ := stdin.ReadRune()
		stdin.Discard(1)

		var updated bool

		switch input {
		case 'w':
			updated = board.MoveUp()
		case 'a':
			updated = board.MoveLeft()
		case 's':
			updated = board.MoveDown()
		case 'd':
			updated = board.MoveRight()
		}

		if updated {
			board.SpawnNewElement()
		}
	}
}
