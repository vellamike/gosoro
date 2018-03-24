package main

// Rough outline of how this ought to work:

// 1. a board is prepared by a board generator and presented to the user. In future the user can create their own boards.
// 2. User is presented with a list of moves. The ruleset/mutator determines which moves are allowed.
// 3. The move is non-terminating, in which case Step 2 is repeated, or it is terminating, in which case it is the turn of the computer to play.

// Game controller should keep a copy of all boards and moves in order to provide the ruleset with all the information it may require.

// New architecture will look something like this:

// 1. A board - this is a data structure that we basically already have. It may also contain information such as which player is to play next, in the middle of a play, or whether any positions are 'hot'
// 2. An instruction - When applied to a board this produces a new board. Instruction is a string of characters (same format which the user uses). The board does not "decide" whether a move was valid or not
// 3. A Mutator - this takes a board and returns all the possible instruction objects. The mutator is specific to the game.
// 4. An AI - given a minimum of a board on which its turn is due and a mutator the player will decide what move it wants to play to improve its position. In the first instance it will most likely use minimax.
// 5. A game controller - this is responsible for receiving user input, handing the board between the AI and opponent. The mutator is instantiated with the following:
//    1. Board originator
//    2. AI
//    3. Mutator

import "gosoro/boardgenerators"
import "gosoro/gamecontrollers"
import "gosoro/ai"
import "gosoro/rulesets"
import "gosoro/evaluators"

import "fmt"


func main() {
	// instantiate a board generator
	board_generator := boardgenerators.Randomboard

	// instantiate an AI
	ai := ai.AI{}

	igisoro_ruleset := rulesets.RuleSet{}

	evaluator := evaluators.Evaluator{}

	// instantiate a game controller, composed of the board generator, AI and the mutator
	controller := gamecontrollers.NewGameController(
		board_generator,
		ai,
		igisoro_ruleset,
		evaluator,
	)

	// Display the board to the user
	fmt.Println("Game started, board as follows:")
	controller.DisplayBoard()

	for controller.Winner < 0 { // keep looping until there is a winner
		if controller.NextPlayer == 0 { // user to move
			controller.UserMove()
			fmt.Println("Last position of User:")
			fmt.Println(controller.LastUserPosition())
		}
		if controller.NextPlayer == 1 { // AI to move
			controller.UserMove()
			fmt.Println("Last position of User:")
			fmt.Println(controller.LastUserPosition())
		}
	}

	fmt.Println("The winner is player ", controller.Winner)
}
