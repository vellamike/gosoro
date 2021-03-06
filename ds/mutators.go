package ds

import "fmt"

type Mutator struct {
}

func (Mutator) ExecuteMove(board Board, move Move, player_number int) Board {

	// Algorithm description:

	// 1. set seeds_in_hand to current_coord
	// 2. set the seed count of current_coord to zero
	// 3. Set current position to the coord where the move started
	// while (seeds_in_hand > 0):
	// 	1. find the next_position from current_position
	// 	2. Add 1 to the next position
	// 	3. Set seeds_in_hand to seeds_in_hand -1
	//      4. Set current_position to next_position

	// TODO: The ability to understand captures should be added to this method,
	// It can then be added to the BoardInterface because it will then support
	// very primitive operations. The Board can provide the mutator with all the methods
	// it needs to understand what moves are possible according to the Mutator ruleset.

	var player, opponent *Player

	if player_number == 1 {
		player = &board.Player_1
		opponent = &board.Player_2
	} else {
		player = &board.Player_2
		opponent = &board.Player_1
	}

	action := move.Action

	if action == "A" || action == "C" {
		move_seeds(player, move)
	} else if action == "S" {
		capture(player, opponent, move)
	}

	return board
}

func capture(agressor, victim *Player, move Move) {

	var captured_seeds int
	column := move.Column

	victim_column := 7 - column
	captured_seeds += victim.Positions[0][victim_column]
	captured_seeds += victim.Positions[1][victim_column]

	victim.Positions[0][victim_column] = 0
	victim.Positions[1][victim_column] = 0

	agressor.Positions[1][column] += captured_seeds

}

func move_seeds(player *Player, move Move) {
	seeds_in_hand := player.Positions[move.Row][move.Column]
	player.Positions[move.Row][move.Column] = 0
	current_row := move.Row
	current_column := move.Column

	for seeds_in_hand > 0 {
		next_row, next_column := next_position(current_row, current_column, move.Action)
		player.Positions[next_row][next_column] += 1
		seeds_in_hand -= 1
		current_row = next_row
		current_column = next_column
	}

	player.LastPosition = Coord{current_row, current_column}
}

func (Mutator) Capture(board Board, capturing_player int, column int) Board {
	fmt.Println("Capturing!")

	var agressor *Player
	var victim *Player

	if capturing_player == 1 {
		agressor = &board.Player_1
		victim = &board.Player_2
	} else {
		agressor = &board.Player_2
		victim = &board.Player_1
	}

	var captured_seeds int

	//row 0
	opponent_column := 7 - column
	captured_seeds += victim.Positions[0][opponent_column]
	victim.Positions[0][opponent_column] = 0
	captured_seeds += victim.Positions[1][opponent_column]
	victim.Positions[1][opponent_column] = 0

	agressor.Positions[1][column] += captured_seeds

	return board
}

func next_position(current_row, current_column int, direction string) (int, int) {
	//Based on the current position and direction, identify the next position
	switch {
	case current_row == 0 && current_column < 7 && direction == "A":
		current_column += 1
	case current_row == 0 && current_column == 7 && direction == "A":
		current_row = 1
	case current_row == 1 && current_column > 0 && direction == "A":
		current_column -= 1
	case current_row == 1 && current_column == 0 && direction == "A":
		current_row = 0
	case current_row == 0 && current_column > 0 && direction == "C": // move left on bottom row
		current_column -= 1
	case current_row == 0 && current_column == 0 && direction == "C": // move up to top row
		current_row = 1
	case current_row == 1 && current_column < 7 && direction == "C": // move right on top row
		current_column += 1
	case current_row == 1 && current_column == 7 && direction == "C": // move down to bottom row
		current_row = 0
	}
	return current_row, current_column
}
