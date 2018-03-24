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
	ruleset    rulesets.RuleSet // Decides what moves are allowed
	evaluator  evaluators.Evaluator
	Winner     int // which player has won, -1 = neither yet, 0 = player 0, 1 = player 1
	NextPlayer int
}

func (self gamecontroller) LastUserPosition() ds.Coord {
	return self.board.Player_1.LastPosition
}

func (self gamecontroller) LastComputerPosition() ds.Coord {
	return self.board.Player_2.LastPosition
}

func NewGameController(generator func() ds.Board, ai ai.AI, ruleset rulesets.RuleSet, evaluator evaluators.Evaluator) *gamecontroller {
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

func capture_possible(board ds.Board, player_number, row, column int) bool {
	//Returns whether landing in this pit would result in capturing oponent's seeds
	_, p_op := ds.PlayersFromName(player_number, &board)
	opponent_column := 7 - column
	opponent_row_0_seeds := p_op.Positions[0][opponent_column]
	opponent_row_1_seeds := p_op.Positions[1][opponent_column]
	return (opponent_row_0_seeds != 0 && opponent_row_1_seeds != 0 && row == 1)
}

func user_move() []ds.Move { // Takes user input as a string and returns a slice of Moves
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your move: ")
	text, _ := reader.ReadString('\n')
	t := string(text)

	var rows, columns []int
	var actions []string

	var moves []ds.Move

	for i := 1; i < len(t)/3+1; i++ {
		row, _ := strconv.Atoi(t[i*3-3 : i*3-2])
		column, _ := strconv.Atoi(t[i*3-2 : i*3-1])
		action := t[i*3-1 : i*3]
		rows = append(rows, row)
		columns = append(columns, column)
		actions = append(actions, action)

		moves = append(moves, ds.Move{row, column, action})

	}
	return moves
}

func (gc *gamecontroller) UserMove() {
	// Ask the user for a move, create an instruction from his reply, apply it to the board

	// STEP 1: Find out what the possible moves are

	validMove := false

	fmt.Println("Computing moves available to user...")
	availableMoves := gc.ruleset.AvailableMoves(gc.board, 0)

	var move int;
	for validMove == false {
		fmt.Println("Available moves are...")
		for i, move := range availableMoves {
			fmt.Println(i, move)
		}
		fmt.Println("<<<===>>>")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your move: ")
		text, _ := reader.ReadString('\n')
		t := string(text)

		fmt.Println("As a string, the selected move is", t)
		move, _ = strconv.Atoi(t[0 : len(t)-1])
		fmt.Println("You selected move: ", move)

		// check that the move is valid:
		validMove = (move > 0) && move < len(availableMoves)
		if validMove == false {
			fmt.Println("That move was not a valid move, please try again..")
		}
	}
	selectedUserMove := availableMoves[move][0] //I think there is a small fix here somewhere, since available moves should be only one move.
	gc.board = gc.board.ExecuteMove(selectedUserMove, 1) // board should be updated too, when executing move it should keep a record of what it is doing..
	fmt.Println("Board after the last move...")
	gc.board.Display()
}

func (gc *gamecontroller) ComputerMove() {
	// Step 1: Ask the AI for the best instruction

	moveSequence := gc.ai.BestInstruction(gc.board, gc.ruleset, gc.evaluator)
	fmt.Println("Computer's response:")
	fmt.Println(moveSequence)

	// Step 2: Apply the instruction

	gc.board = gc.board.ExecuteMoveSequence(moveSequence, 2)

	fmt.Println("Board after computer's response:")
	gc.board.Display()

	fmt.Println("Now computer performing a capture")

	fmt.Println("Board after capture performed:")
	gc.board.Display()

}
