package main

// TO DO:
// When computer and user move, if they capture, they should be able to move again
// If capture does not finish on a bidirectional pit, no further instructions are
// being returned

// Computer is sometimes making illegal moves, these are listed as legal moves but then not correctly
// Implemented - example 13A  which should lead to a capture (i.e it is a logical move) actually
// being effected as 13C. The reason for this is not clear yet, possibly the boards and relevant
// instructions are getting confused
// idea: perhaps each instruction should contain the initial board, instruction, and final board
// this would be a much more safe and OO approach.

import "fmt"
import "bufio"
import "os"
import "strconv"
import "gosoro/ds"
import "gosoro/utils"

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

func players_from_name(player_number int, board *ds.Board) (p, p_op *ds.Player) {
	//Retrun a player based on their identifier
	if player_number == 1 {
		p = &board.Player_1
		p_op = &board.Player_2
	} else {
		p = &board.Player_2
		p_op = &board.Player_1
	}
	return p, p_op
}

func capture_possible(board ds.Board, player_number, row, column int) bool {
	//Returns whether landing in this pit would result in capturing oponent's seeds
	_, p_op := players_from_name(player_number, &board)
	opponent_column := 7 - column
	opponent_row_0_seeds := p_op.Positions[0][opponent_column]
	opponent_row_1_seeds := p_op.Positions[1][opponent_column]
	return (opponent_row_0_seeds != 0 && opponent_row_1_seeds != 0 && row == 1)
}

func perform_capture(board ds.Board, player_number, row, column int) (ds.Board, []ds.Instruction) {
	p, p_op := players_from_name(player_number, &board)
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



func execute_instruction(instruction ds.Instruction, player_number int) (ds.Board, []ds.Instruction) {
	// Execute the instruction (no decison on whether it is legal or not) and return
	// Possible following instructions if there are any. Following instructions
	// Result from captures
	board := instruction.Board
	p, _ := players_from_name(player_number, &board)
	current_row := instruction.Row
	current_column := instruction.Column
	current_direction := instruction.Direction
	var next_instructions []ds.Instruction

	num_seeds := p.Positions[current_row][current_column] //How many seeds will be moved
	p.Positions[current_row][current_column] = 0          // empty the starting pit
	for i := 0; i < num_seeds; i++ {                      //move the seeds
		current_row, current_column = next_position(current_row, current_column, current_direction)
		p.Positions[current_row][current_column] += 1
	}

	// now for the capturing
	if capture_possible(board, player_number, current_row, current_column) {
		board, next_instructions = perform_capture(board, player_number, current_row, current_column)
	}
	return board, next_instructions
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

func all_moves(board ds.Board, player_number int) (boards []ds.Board, instructions [][]ds.Instruction) {
	// first find all rows and columns with more than two
	p, _ := players_from_name(player_number, &board)
	var coords []ds.Coord
	for row_index, row := range p.Positions {
		for column_index, num_seeds := range row {
			if num_seeds > 1 {
				coord := ds.Coord{row_index, column_index, &board, p}
				coords = append(coords, coord)
			}
		}
	}

	if len(coords) <= 0 {
		fmt.Println("COMPUTER HAS BEEN DEFEATED")
	}

	var initial_instructions []ds.Instruction // list of all initial instructions, not yet popualting the stack
	for _, c := range coords {
		instruction := ds.Instruction{c.Row, c.Column, "A", board}          // Perform the counterclockwise move
		initial_instructions = append(initial_instructions, instruction) // Append instruction to initial instruction list
		if board.Is_bidirectional(c.Row, c.Column) {                     // If we can move counterclockwise from this position
			instruction = ds.Instruction{c.Row, c.Column, "C", board}
			initial_instructions = append(initial_instructions, instruction)
		}
	}

	var instructions_stack [][]ds.Instruction // list of lists of instructions

	for _, i_instruction := range initial_instructions { // populate the instruction stack
		instructions_stack = append(instructions_stack, []ds.Instruction{i_instruction})
	}

	for len(instructions_stack) > 0 { // While the stack is not emtpy
		instruction_set := (pop_instruction_stack(&instructions_stack)) // Get a list of instructions

		instruction := instruction_set[len(instruction_set)-1]  // take the last instruction
		b, i := execute_instruction(instruction, player_number) // Get the resulting instruction
		if len(i) > 0 {
			for _, ins := range i { // If more instructions, populate the instructions stack
				new_instruction_set := append(instruction_set, ins)
				instructions_stack = append(instructions_stack, new_instruction_set)
			}
		} else {
			boards = append(boards, b)
			instructions = append(instructions, instruction_set)
		}
	}

	return boards, instructions // return the boards and full corresponding instruction sets
}

func user_move() (rows, columns []int, directions []string) { // returns a slice of instructions
	// Now let's try and get a move from the user
	// TODO: Allow capture commands
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your move: ")
	text, _ := reader.ReadString('\n')
	t := string(text)
	fmt.Println(t)
	fmt.Println(len(t))
	for i := 1; i< len(t) / 3 + 1; i++{

		fmt.Println("In loop ", i)
		row ,_ := strconv.Atoi(t[i * 3 - 3: i * 3 - 2])
		fmt.Println("row:", row)
		column ,_ := strconv.Atoi(t[i * 3 - 2 : i * 3 - 1])
		fmt.Println("col:", column)
		direction := t[i * 3 - 1 : i * 3]
		fmt.Println("direction", direction)
		rows = append(rows, row)
		columns = append(columns, column)
		directions = append(directions, direction)
	}
	return
}

func execute_user_move(board ds.Board) ds.Board {
	row, column, direction := user_move()

	for i := range row {
		instruction := ds.Instruction{row[i], column[i], direction[i], board}
		board, _ = execute_instruction(instruction, 2) // need to figure out what to do if user has a choice
	}

	return board
}

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

func computer_move(board ds.Board) (ds.Board, []ds.Instruction) {
	// Need to update so that the computer choses a move based on some optimality criterion
	boards, instruction_sets := all_moves(board, 1)
	max_index := 0
	current_best_score := 0
	for board_index, b := range boards {
		s := score(b)
		if s > current_best_score {
			current_best_score = s
			max_index = board_index
		}
	}

	fmt.Println("I, the computer, chose the moves which finally evaluate to:")
	boards[max_index].Display()
	for _, instruc := range instruction_sets[max_index] {
		fmt.Println(instruc.Row, instruc.Column, instruc.Direction)
		instruc.Board.Display()
		fmt.Println("---")
	}
	return boards[max_index], instruction_sets[max_index] // return final board and corresponding instruction sets
}

func main() {
	fmt.Println("Instantiating a random board")
	newboard := utils.Random_board(32)
	newboard.Display()
	fmt.Println("Play now begins...")
	for true {
		board1, instructions := computer_move(newboard)
		fmt.Println("Final board:")
		board1.Display()
		fmt.Println(instructions)
		newboard = execute_user_move(board1)
		newboard.Display()
	}
	//      17A14A is a good move
	//	fmt.Println(next_position(1,1,"A"))
	//	fmt.Println(next_position(1,0,"A"))
	//	fmt.Println(next_position(0,0,"A"))
}
