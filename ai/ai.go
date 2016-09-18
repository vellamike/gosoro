package ai

import "gosoro/ds"
import "gosoro/rulesets"
import "gosoro/evaluators"

type AI struct {
}

func (AI) BestInstruction(board ds.Board, ruleset rulesets.RuleSet, evaluator evaluators.Evaluator) ds.Move {

	// A very naive algorithm, get all the available moves and pick the one which evaluates to the highest score without considering the opponent's response

	available_moves := ruleset.AvailableMoves(board, 2)

	var score float32
	var index int

	for i, move := range available_moves {
		current_score := evaluator.Score(board.ExecuteMove(move, 2))
		if current_score > score {
			score = current_score
			index = i
		}
	}

	return available_moves[index]
}
