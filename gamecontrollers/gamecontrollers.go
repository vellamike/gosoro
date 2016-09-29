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
	board     ds.Board
	ai        ai.AI
	ruleset   rulesets.RuleSet
	evaluator evaluators.Evaluator
}

func (self gamecontroller) Winner() int {
	return 0
}

func (self gamecontroller) LastUserPosition() ds.Coord {
	return self.board.Player_1.LastPosition
}

func (self gamecontroller) LastComputerPosition() ds.Coord {
	return self.board.Player_2.LastPosition
}

func NewGameController(generator func() ds.Board, ai ai.AI, ruleset rulesets.RuleSet, evaluator evaluators.Evaluator) *gamecontroller {
	board := generator()
	b := gamecontroller{board, ai, ruleset, evaluator}
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
	moves := user_move()
	fmt.Println("User's moves:")
	fmt.Println(moves)

	fmt.Println("Board before user move is executed:")
	gc.board.Display()
	for _, move := range moves {
		gc.board = gc.board.ExecuteMove(move, 1)
	}
	fmt.Println("Board after user move is executed:")
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
