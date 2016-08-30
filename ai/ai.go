package ai

import "gosoro/ds"
import "math/rand"

type AI struct {
}

func (AI) BestInstruction(board ds.Board) ds.Move {
	// A very naive algorithm, pick a random pit which has more than 1 seed and move clockwise
	player := board.Player_2

	var seed_number int
	var column int
	var row int

	for seed_number < 2{
		row = rand.Intn(2)
		column = rand.Intn(8)
		seed_number = player.Positions[row][column]
	}

	return ds.Move{row, column, "C"}
}
