package main

// Rough outline of how this ought to work:

// 1. a board is prepared by a board generator and presented to the user. In future the user can create their own boards.
// 2. User is presented with a list of moves. The ruleset/mutator determines which moves are allowed.
// 3. The move is non-terminating, in which case Step 2 is repeated, or it is terminating, in which case it is the turn of the computer to play.

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

	// Get a ruleset for the game we are playing, which is Igisoro
	igisoro_ruleset := rulesets.IgisoroRuleSet{}

	// Evaluator will be used by the AI to figure out what good and bad moves are
	evaluator := evaluators.Evaluator{}

	// instantiate a game controller, composed of the board generator, AI and the mutator
	controller := gamecontrollers.NewGameController(
		board_generator,
		ai,
		igisoro_ruleset,
		evaluator,
	)

	// Display the board to the user
	fmt.Println("===> Game started, board as follows:")
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

	fmt.Println("===> The winner is player ", controller.Winner)
}
