package rulesets

import "gosoro/ds"
import "fmt"

type RuleSet struct {
}

func (this RuleSet) AvailableMoves(board ds.Board, player_number int) [][]ds.Move {
	// Still too simple - captures and conditions of moves (clockwise/counterclockwise
	// depending on landing place and pit) need to be taken into account.
	// This is an important step to assist the AI with making a decision by exploring the rulespace

	fmt.Println("Computing all available moves")
	var player ds.Player

	if player_number == 2 {
		player = board.Player_2
	} else {
		player = board.Player_1
	}

	var moves_buffer [][]ds.Move // here we store the moves which may either be terminal or continued

	// First need to get all the stub moves
	for r := 0; r < 2; r++ {
		for c := 0; c < 8; c++ {
			if player.Positions[r][c] > 1 {
				stub := []ds.Move{ds.Move{r, c, "C"}}
				moves_buffer = append(moves_buffer, stub)
			}
		}

	}

	fmt.Println("Stub moves computed:")
	fmt.Println(moves_buffer)

	// now treat moves_buffer as a stack

	var finalised_moves [][]ds.Move

	for len(moves_buffer) > 0 {
		fmt.Println("Moves buffer len != 0, len == ", len(moves_buffer))
		move_index := len(moves_buffer) - 1
		moves_under_consideration := moves_buffer[move_index]
		fmt.Println("Move sequence under consideration: ", moves_under_consideration)
		fmt.Println("Board before move sequence:")
		temporaryBoard := board
		fmt.Println("About to execute move sequence", moves_under_consideration)
		fmt.Println("Board before:")
		temporaryBoard.Display()
		for _, move := range moves_under_consideration {
			temporaryBoard = temporaryBoard.ExecuteMove(move, player_number)
		}
		fmt.Println("Board after:")
		temporaryBoard.Display()
		fmt.Println("Board after move sequence:")
		temporaryBoard.Display()

		last_position := temporaryBoard.PlayerFromNumber(player_number).LastPosition
		fmt.Println("Last position was:")
		fmt.Println(last_position)
		if this.CapturePossible(temporaryBoard, player_number, last_position) && (len(moves_under_consideration) < 10) {
			fmt.Println("Capture is possible")
			new_move_sequence := append(moves_under_consideration, ds.Move{last_position.Row,
				last_position.Column, "S"}) // the seize move
			new_move_sequence = append(new_move_sequence, ds.Move{last_position.Row,
				last_position.Column, "C"}) // the followup move
			moves_buffer[move_index] = new_move_sequence
		} else {
			fmt.Println("Terminal move")
			finalised_moves = append(finalised_moves, moves_under_consideration)
			moves_buffer = moves_buffer[:move_index]
		}
	}
	return finalised_moves
}

func (RuleSet) CapturePossible(board ds.Board, player_number int, pit ds.Coord) bool {

	var player ds.Player
	var opponent ds.Player

	if player_number == 2 {
		player = board.Player_2
		opponent = board.Player_1
	} else {
		player = board.Player_1
		opponent = board.Player_2
	}

	correct_row := pit.Row != 0
	sufficient_seeds := (player.Positions[0][pit.Column] != 0) && (player.Positions[1][pit.Column] != 0)
	opponent_has_seeds_for_capture := (opponent.Positions[0][7-pit.Column] + opponent.Positions[1][7-pit.Column]) != 0
	fmt.Println(correct_row, sufficient_seeds, opponent_has_seeds_for_capture)
	return correct_row && sufficient_seeds && opponent_has_seeds_for_capture
}
