package main

// Thought on design:
// Adding new instructions to the stack shouldn't just be something to do when a bidirectional node occurs
// rather, once a move is completed the next instruction/s, if there are any, should be added to the stack
// along with the mutated board

import "fmt"
import "math/rand"

type player struct {
	positions [2][8]int
}

type Board struct {
	player_1 player
	player_2 player
}

type Instruction struct {
	row       int
	column    int
	direction string
}

func (this Board) is_bidirectional(player, row, column int) bool {
	var bidir bool
	//Omweso rules
	if column == 0 || column == 2 || column == 6 || column == 7 {
		bidir = true
	} else {
		bidir = false
	}
	return bidir
}

func random_position(num_seeds int) player {
	var p player

	for i := 0; i < num_seeds; i++ {
		row := rand.Intn(2)
		column := rand.Intn(8)
		p.positions[row][column] += 1
	}

	return p

}

func random_board(num_seeds int) Board {
	var newboard Board

	newboard.player_1 = random_position(num_seeds)
	newboard.player_2 = random_position(num_seeds)

	return newboard
}

func next_position(current_row, current_column int, direction string) (int, int) {
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

func players_from_name(player_number int, board *Board) (p, p_op *player) {
	if player_number == 1 {
		p = &board.player_1
		p_op = &board.player_2
	} else {
		p = &board.player_2
		p_op = &board.player_1
	}
	return p, p_op
}

func capture_possible(board Board, player_number, row, column int) bool {
	_, p_op := players_from_name(player_number, &board)
	opponent_column := 7 - column
	opponent_row_0_seeds := p_op.positions[0][opponent_column]
	opponent_row_1_seeds := p_op.positions[1][opponent_column]
	return (opponent_row_0_seeds != 0 && opponent_row_1_seeds != 0 && row == 1)
}

func perform_capture(board Board, player_number, row, column int) (Board, []Instruction) {
	p, p_op := players_from_name(player_number, &board)
	opponent_column := 7 - column
	p_op.positions[0][opponent_column] = 0
	p_op.positions[1][opponent_column] = 0
	opponent_row_0_seeds := p_op.positions[0][opponent_column]
	opponent_row_1_seeds := p_op.positions[1][opponent_column]
	captured_seeds := opponent_row_0_seeds + opponent_row_1_seeds
	p.positions[row][column] += captured_seeds

	var next_instructions []Instruction

	if board.is_bidirectional(player_number, row, column) {
		fmt.Println("Found a bidirectional board")
		i1 := Instruction{row, column, "C"}
		i2 := Instruction{row, column, "A"}
		next_instructions = []Instruction{i1, i2}
	}
	return board, next_instructions
}

func execute_instruction() {
	// this method should take as an argument a board, player number and an instruction
	// It should return the modified board and any subsequent instructions
	// A controller loop (e.g `move`) could then make calls to this method until all
	// leaf boards have been found
}

func move(instruction Instruction, board Board, player_number int) (final_board Board, next_instructions []Instruction) {
	p, _ := players_from_name(player_number, &board)

	leaf_board := false // whether or not a leaf board has been reached
	row := instruction.row
	column := instruction.column
	direction := instruction.direction

	for (leaf_board == false) && (len(next_instructions) == 0) {
		fmt.Println("Performing a move")
		num_seeds := p.positions[row][column]
		fmt.Println("Number of seeds:")
		fmt.Println(num_seeds)
		p.positions[row][column] = 0     // empty the starting pit
		for i := 0; i < num_seeds; i++ { //move the seeds, currently not using direction
			row, column = next_position(row, column, direction)
			p.positions[row][column] += 1
		}

		if capture_possible(board, player_number, row, column) {
			direction = "A" // If next_instructions has not been populated, the direction should be A
			board, next_instructions = perform_capture(board, player_number, row, column)
		} else {
			leaf_board = true
		}
	}
	return board, next_instructions
}

func main() {
	fmt.Println("Instantiating a random board")
	newboard := random_board(12)
	fmt.Println(newboard)
	new_instruction := Instruction{1, 2, "C"}
	fmt.Println(new_instruction)
	board, instructions := move(new_instruction, newboard, 1) // if there are no instructions, it is a terminal board
	fmt.Println(board)
	fmt.Println(instructions)
}
