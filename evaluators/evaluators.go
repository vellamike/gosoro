package evaluators

import "gosoro/ds"

type Evaluator struct {
}

func (Evaluator) Score(board ds.Board) int {
	// returns score of player 1
	var total_seeds int

	for _, i := range []int{1, 0} {
		for _, s := range board.Player_1.Positions[i] {
			total_seeds += s
		}
	}

	return total_seeds
}
