package mutators

import "gosoro/ds"
import "fmt"

type Mutator struct {
}

func pop_instruction_stack(stack *[][]ds.Instruction) []ds.Instruction {
	// An instruction stack is slice of slices of instructions.
	// This method returns a value off the stack, and removes that value
	// From the stack
	len_stack := len(*stack)
	val := (*stack)[len_stack-1]
	*stack = (*stack)[:len_stack-1]
	return val
}

func (Mutator) ExecuteMove(board ds.Board, move ds.Move, player_number int) ds.Board {
	fmt.Println("Updating board")

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

	var player *ds.Player

	if player_number == 1 {
		player = &board.Player_1
	} else {
		player = &board.Player_2
	}

	seeds_in_hand := player.Positions[move.Row][move.Column]
	player.Positions[move.Row][move.Column] = 0
	current_row := move.Row
	current_column := move.Column

	for seeds_in_hand > 0 {
		next_row, next_column := next_position(current_row, current_column, move.Direction)
		player.Positions[next_row][next_column] += 1
		seeds_in_hand -= 1
		current_row = next_row
		current_column = next_column
	}

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

// This method needs rethinking because it relies on the ExecuteInstruction method
//func all_moves(board ds.Board, player_number int) (boards []ds.Board, instructions [][]ds.Instruction) {
//	// first find all rows and columns with more than two
//	p, _ := ds.PlayersFromName(player_number, &board)
//	var coords []ds.Coord
//	for row_index, row := range p.Positions {
//		for column_index, num_seeds := range row {
//			if num_seeds > 1 {
//				coord := ds.Coord{row_index, column_index, &board, p}
//				coords = append(coords, coord)
//			}
//		}
//	}
//
//	if len(coords) <= 0 {
//		fmt.Println("COMPUTER HAS BEEN DEFEATED")
//	}
//
//	var initial_instructions []ds.Instruction // list of all initial instructions, not yet popualting the stack
//	for _, c := range coords {
//		instruction := ds.Instruction{c.Row, c.Column, "A", board}       // Perform the counterclockwise move
//		initial_instructions = append(initial_instructions, instruction) // Append instruction to initial instruction list
//		if board.Is_bidirectional(c.Row, c.Column) {                     // If we can move counterclockwise from this position
//			instruction = ds.Instruction{c.Row, c.Column, "C", board}
//			initial_instructions = append(initial_instructions, instruction)
//		}
//	}
//
//	var instructions_stack [][]ds.Instruction // list of lists of instructions
//
//	for _, i_instruction := range initial_instructions { // populate the instruction stack
//		instructions_stack = append(instructions_stack, []ds.Instruction{i_instruction})
//	}
//
//	for len(instructions_stack) > 0 { // While the stack is not emtpy
//		instruction_set := (pop_instruction_stack(&instructions_stack)) // Get a list of instructions
//
//		instruction := instruction_set[len(instruction_set)-1]  // take the last instruction
//		b, i := ds.ExecuteInstruction(instruction, player_number) // Get the resulting instruction
//		if len(i) > 0 {
//			for _, ins := range i { // If more instructions, populate the instructions stack
//				new_instruction_set := append(instruction_set, ins)
//				instructions_stack = append(instructions_stack, new_instruction_set)
//			}
//		} else {
//			boards = append(boards, b)
//			instructions = append(instructions, instruction_set)
//		}
//	}
//
//	return boards, instructions // return the boards and full corresponding instruction sets
//}
