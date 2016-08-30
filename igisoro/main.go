package main

// TO DO:x
// When computer and user move, if they capture, they should be able to move again
// If capture does not finish on a bidirectional pit, no further instructions are
// being returned

// Computer is sometimes making illegal moves, these are listed as legal moves but then not correctly
// Implemented - example 13A  which should lead to a capture (i.e it is a logical move) actually
// being effected as 13C. The reason for this is not clear yet, possibly the boards and relevant
// instructions are getting confused
// idea: perhaps each instruction should contain the initial board, instruction, and final board
// this would be a much more safe and OO approach.

// New architecture:

// Now that I'm a bit more confident with go, I'd like to set up the application with the following objects:
// 1. A board - this is a data structure that we basically already have. It may also contain information such as which player is to play next, in the middle of a play, or whether any positions are 'hot'
// 2. An instruction - When applied to a board this produces a new board. Instruction is a string of characters (same format which the user uses). The board does not "decide" whether a move was valid or not
// 3. A Mutator - this takes a board and returns all the possible instruction objects. The mutator is specific to the game.
// 4. An AI - given a minimum of a board on which its turn is due and a mutator the player will decide what move it wants to play to improve its position. In the first instance it will most likely use minimax.
// 5. A game controller - this is responsible for receiving user input, handing the board between the AI and opponent. The mutator is instantiated with the following:
//    1. Board originator
//    2. AI
//    3. Mutator

import "gosoro/ds"
import "gosoro/boardgenerators"
import "gosoro/mutators"
import "gosoro/gamecontrollers"
import "gosoro/ai"

func score(board ds.Board) int {
	//returns the score
	positions := board.Player_1.Positions
	total := 0
	for _, p := range positions {
		for _, i := range p {
			total += i
		}
	}
	return total
}

//func computer_move(board ds.Board) (ds.Board, []ds.Instruction) {
//	// Need to update so that the computer choses a move based on some optimality criterion
//	boards, instruction_sets := all_moves(board, 1)
//	max_index := 0
//	current_best_score := 0
//	for board_index, b := range boards {
//		s := score(b)
//		if s > current_best_score {
//			current_best_score = s
//			max_index = board_index
//		}
//	}
//
//	fmt.Println("I, the computer, chose the moves which finally evaluate to:")
//	boards[max_index].Display()
//	for _, instruc := range instruction_sets[max_index] {
//		fmt.Println(instruc.Row, instruc.Column, instruc.Direction)
//		instruc.Board.Display()
//		fmt.Println("---")
//	}
//	return boards[max_index], instruction_sets[max_index] // return final board and corresponding instruction sets
//}

func main() {

	// instantiate a board generator
	board_generator := boardgenerators.Randomboard

	// instantiate a mutator
	mutator := mutators.Mutator{}

	// instantiate an AI
	ai := ai.AI{}

	// instantiate a game controller, composed of the board generator, AI and the mutator
	controller := gamecontrollers.NewGameController(board_generator, ai, mutator)

	// Display the board to the user
	controller.DisplayBoard()

	for controller.Winner() < 1 {
		// Ask the user for their move
		controller.UserMove()
		// Computer plays its move
		controller.ComputerMove()
	}
}
