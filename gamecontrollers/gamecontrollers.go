package gamecontrollers

import "gosoro/ds"
import "gosoro/ai"
import "gosoro/mutators"

import "fmt"
import "bufio"
import "os"
import "strconv"

type gamecontroller struct {
	board   ds.Board
	ai      ai.AI
	mutator mutators.Mutator
}

func NewGameController(generator func() ds.Board, ai ai.AI, mutator mutators.Mutator) *gamecontroller {
	board := generator()
	b := gamecontroller{board, ai, mutator}
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

func perform_capture(board ds.Board, player_number, row, column int) (ds.Board, []ds.Instruction) {
	p, p_op := ds.PlayersFromName(player_number, &board)
	opponent_column := 7 - column
	opponent_row_0_seeds := p_op.Positions[0][opponent_column]
	opponent_row_1_seeds := p_op.Positions[1][opponent_column]
	p_op.Positions[0][opponent_column] = 0
	p_op.Positions[1][opponent_column] = 0
	captured_seeds := opponent_row_0_seeds + opponent_row_1_seeds
	p.Positions[row][column] += captured_seeds

	var next_instructions []ds.Instruction

	i1 := ds.Instruction{row, column, "A", board}
	next_instructions = []ds.Instruction{i1}

	if board.Is_bidirectional(row, column) {
		i2 := ds.Instruction{row, column, "C", board}
		next_instructions = append(next_instructions, i2)
	}
	//fmt.Println("Capture has been evaluated, the following instructions are being returned:")
	//fmt.Println(next_instructions)
	return board, next_instructions
}

func user_move() []ds.Move { // Takes user input as a string and returns a slice of Moves
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your move: ")
	text, _ := reader.ReadString('\n')
	t := string(text)

	var rows, columns []int
	var directions []string

	var moves []ds.Move

	for i := 1; i < len(t)/3+1; i++ {
		row, _ := strconv.Atoi(t[i*3-3 : i*3-2])
		column, _ := strconv.Atoi(t[i*3-2 : i*3-1])
		direction := t[i*3-1 : i*3]
		rows = append(rows, row)
		columns = append(columns, column)
		directions = append(directions, direction)

		moves = append(moves, ds.Move{row, column, direction})

	}
	return moves
}

func (gc gamecontroller) UserMove() {
	// Ask the user for a move, create an instruction from his reply, apply it to the board

	moves := user_move()
	fmt.Println("User's moves:")
	fmt.Println(moves)

	fmt.Println("Board before user move is executed:")
	gc.board.Display()
	for _, move := range moves {
		gc.board = gc.mutator.ExecuteMove(gc.board, move, 1)
	}
	fmt.Println("Board after user move is executed:")
	gc.board.Display()

	// TODO: once the above is complete, need a method to get all possible moves for the
	// computer. This will involve certain things like the board's ability to remember the last move,
	// or it is possible that this logic should live in the mutator.
}
