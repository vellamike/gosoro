package gamecontrollers

import "gosoro/ds"
import "gosoro/ai"
import "gosoro/rulesets"
import "gosoro/evaluators"

import "fmt"
import "bufio"
import "os"
import "strconv"

type gamecontroller struct {
	board      ds.Board //DS which contains the board position, next player to move etc...
	ai         ai.AI
	ruleset    rulesets.IgisoroRuleSet // Decides what moves are allowed
	evaluator  evaluators.Evaluator
	Winner     int // which player has won, -1 = neither yet, 0 = player 0, 1 = player 1
	NextPlayer int
}

func NewGameController(generator func() ds.Board, ai ai.AI, ruleset rulesets.IgisoroRuleSet, evaluator evaluators.Evaluator) *gamecontroller {
	board := generator()
	b := gamecontroller{board,
		ai,
		ruleset,
		evaluator,
		-1, // No winner yet
		0}  // Human to start
	return &b

}

func (gc gamecontroller) DisplayBoard() {
	gc.board.Display()
}

// TODO: This should be handled by the ruleset?
func capture_possible(board ds.Board, player_number, row, column int) bool {
	//Returns whether landing in this pit would result in capturing oponent's seeds
	_, p_op := ds.PlayersFromName(player_number, &board)
	opponent_column := 7 - column
	opponent_row_0_seeds := p_op.Positions[0][opponent_column]
	opponent_row_1_seeds := p_op.Positions[1][opponent_column]
	return (opponent_row_0_seeds != 0 && opponent_row_1_seeds != 0 && row == 1)
}

func (gc *gamecontroller) UserMove() {
	// Ask the user for a move, create an instruction from their reply
	// apply it to the board, update the controller:
	// 0. Controller history will require updating.
	// 1. Has the user won the game?
	// 2. Whose move is next?

	// Step 1: Find out what the possible moves are

	fmt.Println("Computing moves available to user...")
	availableMoves := gc.ruleset.AvailableMoves(gc.board, 0)

	validMove := false

	var chosenMove int

	for validMove == false {
		fmt.Println("Available moves are...")
		for i, move := range availableMoves {
			fmt.Println(i, move)
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your move: ")
		text, _ := reader.ReadString('\n')
		t := string(text)

		c, err := strconv.Atoi(t[0 : len(t) - 1])
		chosenMove = c
		// check that the move is valid, if not then the user will be asked again
		validMove = (err == nil) && (chosenMove >= 0) && chosenMove < len(availableMoves)
		if validMove == false {
			fmt.Println("That move was not a valid move, please try again..")
		}
	}

	//TODO: Available moves should only contain one move
	selectedUserMove := availableMoves[chosenMove]

	fmt.Println("Applying the move ", chosenMove, " :", selectedUserMove)

	gc.ApplyMove(selectedUserMove, 0)
}


func (gc *gamecontroller) ApplyMove(move ds.Move, user int){
	// this function has the following responsibilities:
	// 1. Add the move and old boards to the controller's record
	// 2. Update the board to show the latest move
	// 3. Check if either player has won the game
	// 4. If there isn't a winner, decide who the next player is (same player as applied the move
	//    In the case of a branch move and same player otherwise.

}

func (gc * gamecontroller) ComputerMove(){

}
