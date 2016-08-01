package gamecontrollers

import "gosoro/ds"
import "gosoro/ai"
import "gosoro/mutators"

import "fmt"
import "bufio"
import "os"
import "strconv"

type gamecontroller struct {
	board   ds.BoardInterface
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
	fmt.Println(t)
	fmt.Println(len(t))

	var rows, columns []int
	var directions []string

	var moves []ds.Move

	for i := 1; i < len(t)/3+1; i++ {

		fmt.Println("In loop ", i)
		row, _ := strconv.Atoi(t[i*3-3 : i*3-2])
		fmt.Println("row:", row)
		column, _ := strconv.Atoi(t[i*3-2 : i*3-1])
		fmt.Println("col:", column)
		direction := t[i*3-1 : i*3]
		fmt.Println("direction", direction)
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
	fmt.Println(gc.board)
	for _, move := range moves {
		gc.board = gc.mutator.ExecuteMove(gc.board, move)
	}
	fmt.Println("Board after user move is executed:")
	fmt.Println(gc.board)

	// TODO: once the above is complete, need a method to get all possible moves for the
	// computer. This will involve certain things like the board's ability to remember the last move,
	// or it is possible that this logic should live in the mutator.
}
