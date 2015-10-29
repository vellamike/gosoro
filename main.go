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
import "math/rand"
import "bufio"
import "os"
import "strconv"

type player struct {
	//Represents a player's territory in their frame of reference
	positions [2][8]int
}

type coord struct {
	row    int
	column int
	board  *Board
	player *player
}

type Board struct {
	player_1 player
	player_2 player
}

type Instruction struct {
	row       int
	column    int
	direction string
	board     Board
}

func (this Board) is_bidirectional(row, column int) bool {
	//Whether clockwise and counterclockwise moves are allowed from this position
	var bidir bool
	if column == 0 || column == 1 || column == 6 || column == 7 {
		bidir = true
	} else {
		bidir = false
	}
	return bidir
}

func (this Board) display() {
	//Display the board to the screen from the computer's perspective
	fmt.Println(reverse_array(this.player_2.positions[0]))
	fmt.Println(reverse_array(this.player_2.positions[1]))
	fmt.Println(this.player_1.positions[1])
	fmt.Println(this.player_1.positions[0])
}

func random_position(num_seeds int) player {
	//choose a random pit
	var p player

	for i := 0; i < num_seeds; i++ {
		row := rand.Intn(2)
		column := rand.Intn(8)
		p.positions[row][column] += 1
	}

	return p

}

func random_board(num_seeds int) Board {
	//Initialize a random board
	var newboard Board

	newboard.player_1 = random_position(num_seeds)
	newboard.player_2 = random_position(num_seeds)

	return newboard
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

func players_from_name(player_number int, board *Board) (p, p_op *player) {
	//Retrun a player based on their identifier
	if player_number == 1 {
		p = &board.player_1
		p_op = &board.player_2
	} else {
		p = &board.player_2
		p_op = &board.player_1
	}
	return p, p_op
}

func capture_possible(board Board, player_number, row, column int) bool {
	//Returns whether landing in this pit would result in capturing oponent's seeds
	_, p_op := players_from_name(player_number, &board)
	opponent_column := 7 - column
	opponent_row_0_seeds := p_op.positions[0][opponent_column]
	opponent_row_1_seeds := p_op.positions[1][opponent_column]
	return (opponent_row_0_seeds != 0 && opponent_row_1_seeds != 0 && row == 1)
}

func perform_capture(board Board, player_number, row, column int) (Board, []Instruction) {
	p, p_op := players_from_name(player_number, &board)
	opponent_column := 7 - column
	opponent_row_0_seeds := p_op.positions[0][opponent_column]
	opponent_row_1_seeds := p_op.positions[1][opponent_column]
	p_op.positions[0][opponent_column] = 0
	p_op.positions[1][opponent_column] = 0
	captured_seeds := opponent_row_0_seeds + opponent_row_1_seeds
	p.positions[row][column] += captured_seeds

	var next_instructions []Instruction

	i1 := Instruction{row, column, "A", board}
	next_instructions = []Instruction{i1}

	if board.is_bidirectional(row, column) {
		i2 := Instruction{row, column, "C", board}
		next_instructions = append(next_instructions, i2)
	}
	//fmt.Println("Capture has been evaluated, the following instructions are being returned:")
	//fmt.Println(next_instructions)
	return board, next_instructions
}

func reverse_array(arr [8]int) [8]int {
	//Return the reverse of a size-8 array, handy for visualisation
	num_elements := len(arr)
	var reversed_array [8]int
	for i := 0; i < num_elements; i++ {
		reversed_array[i] = arr[num_elements-i-1]
	}
	return reversed_array
}

func execute_instruction(instruction Instruction, player_number int) (Board, []Instruction) {
	// Execute the instruction (no decison on whether it is legal or not) and return
	// Possible following instructions if there are any. Following instructions
	// Result from captures
	board := instruction.board
	p, _ := players_from_name(player_number, &board)
	current_row := instruction.row
	current_column := instruction.column
	current_direction := instruction.direction
	var next_instructions []Instruction

	num_seeds := p.positions[current_row][current_column] //How many seeds will be moved
	p.positions[current_row][current_column] = 0          // empty the starting pit
	for i := 0; i < num_seeds; i++ {                      //move the seeds
		current_row, current_column = next_position(current_row, current_column, current_direction)
		p.positions[current_row][current_column] += 1
	}

	// now for the capturing
	if capture_possible(board, player_number, current_row, current_column) {
		board, next_instructions = perform_capture(board, player_number, current_row, current_column)
	}
	return board, next_instructions
}

func pop_instruction_stack(stack *[][]Instruction) []Instruction {
	// An instruction stack is slice of slices of instructions.
	// This method returns a value off the stack, and removes that value
	// From the stack
	len_stack := len(*stack)
	val := (*stack)[len_stack-1]
	*stack = (*stack)[:len_stack-1]
	return val
}

func all_moves(board Board, player_number int) (boards []Board, instructions [][]Instruction) {
	// first find all rows and columns with more than two
	p, _ := players_from_name(player_number, &board)
	var coords []coord
	for row_index, row := range p.positions {
		for column_index, num_seeds := range row {
			if num_seeds > 1 {
				coord := coord{row_index, column_index, &board, p}
				coords = append(coords, coord)
			}
		}
	}

	if len(coords) <= 0 {
		fmt.Println("COMPUTER HAS BEEN DEFEATED")
	}

	var initial_instructions []Instruction // list of all initial instructions, not yet popualting the stack
	for _, c := range coords {
		instruction := Instruction{c.row, c.column, "A", board}          // Perform the counterclockwise move
		initial_instructions = append(initial_instructions, instruction) // Append instruction to initial instruction list
		if board.is_bidirectional(c.row, c.column) {                     // If we can move counterclockwise from this position
			instruction = Instruction{c.row, c.column, "C", board}
			initial_instructions = append(initial_instructions, instruction)
		}
	}

	var instructions_stack [][]Instruction // list of lists of instructions

	for _, i_instruction := range initial_instructions { // populate the instruction stack
		instructions_stack = append(instructions_stack, []Instruction{i_instruction})
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

func user_move() (row, column int, direction string) {
	// Now let's try and get a move from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your move: ")
	text, _ := reader.ReadString('\n')
	t := string(text)
	fmt.Println(t)
	row, _ = strconv.Atoi(t[:1])
	column, _ = strconv.Atoi(t[1:2])
	direction = t[2:3]
	return
}

func execute_user_move(board Board) Board {
	row, column, direction := user_move()
	instruction := Instruction{row, column, direction, board}
	b, _ := execute_instruction(instruction, 2) // need to figure out what to do if user has a choice
	return b
}

func score(board Board) int {
	//returns the score
	positions := board.player_1.positions
	total := 0
	for _, p := range positions {
		for _, i := range p {
			total += i
		}
	}
	return total
}

func computer_move(board Board) (Board, []Instruction) {
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
	boards[max_index].display()
	for _, instruc := range instruction_sets[max_index] {
		fmt.Println(instruc.row, instruc.column, instruc.direction)
		instruc.board.display()
		fmt.Println("---")
	}
	return boards[max_index], instruction_sets[max_index] // return final board and corresponding instruction sets
}

func main() {
	fmt.Println("Instantiating a random board")
	newboard := random_board(32)
	newboard.display()
	fmt.Println("Play now begins...")
	for true {
		board1, instructions := computer_move(newboard)
		fmt.Println("Final board:")
		board1.display()
		fmt.Println(instructions)
		newboard = execute_user_move(board1)
		newboard.display()
	}

	//	fmt.Println(next_position(1,1,"A"))
	//	fmt.Println(next_position(1,0,"A"))
	//	fmt.Println(next_position(0,0,"A"))
}
