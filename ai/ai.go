package ai

import "gosoro/ds"
import "gosoro/rulesets"

type AI struct {
}

func (AI) BestInstruction(board ds.Board, ruleset rulesets.RuleSet) ds.Move {
	// A very naive algorithm, pick a random pit which has more than 1 seed and move clockwise

	available_moves := ruleset.AvailableMoves(board, 2)

	return available_moves[0]
}
