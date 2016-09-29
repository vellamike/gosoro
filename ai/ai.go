package ai

import "gosoro/ds"
import "gosoro/rulesets"
import "gosoro/evaluators"
import "fmt"

type AI struct {
}

func (AI) BestInstruction(board ds.Board, ruleset rulesets.RuleSet, evaluator evaluators.Evaluator) []ds.Move {

	// A very naive algorithm, get all the available moves and pick the one which evaluates to the highest score without considering the opponent's response

	available_moves := ruleset.AvailableMoves(board, 2)

	fmt.Println("Considering available moves from:")
	fmt.Println(available_moves)

	index := 0
	initial_move := available_moves[index]
	score := evaluator.Score(board.ExecuteMoveSequence(initial_move, 2)) //initialize score to first position
	fmt.Println("Current best score is", score)
	for i, move_sequence := range available_moves[1:] {
		current_score := evaluator.Score(board.ExecuteMoveSequence(move_sequence, 2))
		fmt.Println("Score of move ", move_sequence, "is ", current_score)
		if current_score < score { // lower scores are better
			score = current_score
			index = i + 1
		}
	}
	fmt.Println("AI suggests move: ", available_moves[index])
	return available_moves[index]
}
