package main

// Initial rules: No reversing, only victory is if other player can't move

// Need the following methods:
// 1. Board initialiser - this is a method which just returns a board.
// 2. Scoring function
// 3. Board mmutator
// 4. A string-based notation for the move.
// 5. A player-perspective on the board, so that confusions regarding clockwise/counterclockwise etc are easily resolved.

// Go Questions:
// Q 1. What is a slice?
//
// Q 2. How do you add a method to a type
// A func (this Type) func_name(func_param param_type) (return_type) {...}
//
// Q 3. What is an interface
// A If a type has all the correct function names, including signatures, then it satisfies an interface
//   And can be passed to another method which takes that interface.
//   Still not sure about the details of this but it sounds quite interesting.
// Q how do maps work in go?
//

// UPDATE: I've decided that the 4x8 board representation is very unhelpful.
// The 2x2x8 is better, that is to say, the board is represented as two player views.
// Each player view is a 2x8 board.
// For calcuations of scores this can still be mapped to the 4x8 representation if need be,
// But for carrying out mutations this is a much more simple strategy.

// Thoughts on mutators:
// In some positions a decision can be made whether to go clockwise or counterclockwise.
// Some sort of tree is going to be needed to keep track of decisions, OR we could not
// implement this aspect of it for now.

// Instruction Format: RCD (row, column, direction)
// Example: 02C or 16A

import "fmt"
import "math/rand"
import "strconv"

type player struct {
	positions [2][8]int
}

type Board struct {
	player_1 player
	player_2 player
}

func (this Board) is_birectional(player, row, column int) bool {
	return false
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

func next_position(current_row, current_column int) (int, int) {
	switch {
	case current_row == 0 && current_column < 7:
		current_column += 1
	case current_row == 0 && current_column == 7:
		current_row = 1
	case current_row == 1 && current_column > 0:
		current_column -= 1
	case current_row == 1 && current_column == 0:
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
func move(row, column int, direction string, board Board, player_number int) (terminal_board bool, final_board Board, next_instructions []string) {
	p, p_op := players_from_name(player_number, &board)

	terminal_board = false // whether or not a leaf board has been reached
	var ins []string       // If not a leaf node, the instructions available from this board
	for (terminal_board == false) || len(ins) != 0 {
		num_seeds := p.positions[row][column]
		p.positions[row][column] = 0     // empty the starting pit
		for i := 0; i < num_seeds; i++ { //move the seeds, currently not using direction
			row, column = next_position(row, column)
			p.positions[row][column] += 1
		} // need a separate method to perform the capture (lines below)
		// and a new one to decide what moves are available from that
		// from the new position, which in any case will be useful separately
		oponent_column := 7 - column // oponent's column
		oponent_row_0_seeds := p_op.positions[0][oponent_column]
		oponent_row_1_seeds := p_op.positions[1][oponent_column]
		if oponent_row_0_seeds != 0 && oponent_row_1_seeds != 0 && row == 1 { // capture occurs
			p_op.positions[0][oponent_column] = 0
			p_op.positions[1][oponent_column] = 0
			captured_seeds := oponent_row_0_seeds + oponent_row_1_seeds
			p.positions[row][column] += captured_seeds
			//if a capture occurs, need to evaluate whether there is more than one possible decision
			if board.is_birectional(player_number, row, column) {
				//this is never evaluated fo rnow, but will need to append new instructions
			}
		} else {
			terminal_board = true
		}
	}
	return terminal_board, board, ins
}

// methods needed:

// 1. Return an array of all available layer-1 instructions
// 2. Consume that array, returning a list of all final boards and corresponding instructions,
//    including those for instructions which did not evaluate to leaf boards

func main() {
	fmt.Println("Instantiating a random board")
	newboard := random_board(12)
	fmt.Println(newboard)
	terminal_board, moved_board, instructions := move(1, 2, "A", newboard, 1)
	fmt.Println(terminal_board)
	fmt.Println(moved_board)
	fmt.Println(instructions)
	fmt.Println(strconv.convert(1))
}
