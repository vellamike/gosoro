package main

// TO DO:
// 1. When computer and user move, if they capture, they can move again
// 2. Computer should pick a move based on some goal, not just randomly
import "fmt"
import "math/rand"
import "bufio"
import "os"
import "strconv"

type player struct {
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
	var bidir bool
	//Omweso rules
	if column == 0 || column == 1 || column == 6 || column == 7 {
		bidir = true
	} else {
		bidir = false
	}
	return bidir
}

func (this Board) display() {
	fmt.Println(reverse_array(this.player_2.positions[0]))
	fmt.Println(reverse_array(this.player_2.positions[1]))
	fmt.Println(this.player_1.positions[1])
	fmt.Println(this.player_1.positions[0])
}

func random_position(num_seeds int) player {
	var p player

	for i := 0; i < num_seeds; i++ {
		row := rand.Intn(2)
		column := rand.Intn(8)
		p.positions[row][column] += 1
	}

	return p

}

func random_board(num_seeds int) Board {
	var newboard Board

	newboard.player_1 = random_position(num_seeds)
	newboard.player_2 = random_position(num_seeds)

	return newboard
}

func next_position(current_row, current_column int, direction string) (int, int) {
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

	if board.is_bidirectional(row, column) {
		i1 := Instruction{row, column, "C", board}
		i2 := Instruction{row, column, "A", board}
		next_instructions = []Instruction{i1, i2}
	}
	return board, next_instructions
}

func execute_instruction(instruction Instruction, player_number int) (Board, []Instruction) {
	board := instruction.board
	p, _ := players_from_name(player_number, &board)
	current_row := instruction.row
	current_column := instruction.column
	current_direction := instruction.direction
	var next_instructions []Instruction

	// no while loops, just execute the instruction, including a capture, and return the board along with any next moves

	num_seeds := p.positions[current_row][current_column]
	p.positions[current_row][current_column] = 0 // empty the starting pit
	for i := 0; i < num_seeds; i++ {             //move the seeds, currently not using direction
		current_row, current_column = next_position(current_row, current_column, current_direction)
		p.positions[current_row][current_column] += 1
	}

	// now for the capturing
	if capture_possible(board, player_number, current_row, current_column) {
		board, next_instructions = perform_capture(board, player_number, current_row, current_column)
	}
	return board, next_instructions
}

func reverse_array(arr [8]int) [8]int {
	num_elements := len(arr)
	var reversed_array [8]int
	for i := 0; i < num_elements; i++ {
		reversed_array[i] = arr[num_elements-i-1]
	}
	return reversed_array
}

func pop_instruction_stack(stack *[][]Instruction) []Instruction {
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

	var initial_instructions []Instruction
	for _, c := range coords {
		instruction := Instruction{c.row, c.column, "A", board}
		initial_instructions = append(initial_instructions, instruction)
		if board.is_bidirectional(c.row, c.column) {
			instruction = Instruction{c.row, c.column, "C", board}
			initial_instructions = append(initial_instructions, instruction)
		}
	}

	var instructions_stack [][]Instruction

	for _, i_instruction := range initial_instructions {
		instructions_stack = append(instructions_stack, []Instruction{i_instruction})
	}

	for len(instructions_stack) > 0 {
		instruction_set := (pop_instruction_stack(&instructions_stack))

		instruction := instruction_set[len(instruction_set)-1]
		b, i := execute_instruction(instruction, player_number)
		if len(i) > 0 {
			for _, ins := range i {
				new_instruction_set := append(instruction_set, ins)
				instructions_stack = append(instructions_stack, new_instruction_set)
			}
		} else {
			boards = append(boards, b)
			instructions = append(instructions, instruction_set)
		}
	}

	return boards, instructions
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

func score(board Board) (int){
	//returns the score
	positions := board.player_1.positions
	total := 0
	for _,p := range positions{
		for _,i := range p{
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
	for board_index, b := range boards{
		s := score(b)
		if s > current_best_score{
			current_best_score = s
			max_index = board_index
		}
	}

	fmt.Println("I, the computer, chose this move:")
	boards[max_index].display()
	return boards[max_index], instruction_sets[max_index]
}

func main() {
	fmt.Println("Instantiating a random board")
	newboard := random_board(16)
	newboard.display()
	fmt.Println("Play now begins...")
	for true {
		board1, instructions := computer_move(newboard)
		fmt.Println(instructions[0].row, instructions[0].column, instructions[0].direction)

		newboard = execute_user_move(board1)
		newboard.display()
	}
}
