package ds

import "fmt"

type BoardInterface interface {
	Display()
}

type Player struct {
	//Represents a player's territory in their frame of reference
	Positions [2][8]int
}

type Coord struct {
	Row    int
	Column int
	Board  *Board
	Player *Player
}

type Board struct {
	Player_1 Player
	Player_2 Player
}

type Instruction struct {
	Row       int
	Column    int
	Direction string
	Board     Board
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

func (this Board) Is_bidirectional(row, column int) bool {
	//Whether clockwise and counterclockwise moves are allowed from this position
	var bidir bool
	if column == 0 || column == 1 || column == 6 || column == 7 {
		bidir = true
	} else {
		bidir = false
	}
	return bidir
}

func (this Board) Display() {
	//Display the board to the screen from the computer's perspective
	fmt.Println(reverse_array(this.Player_2.Positions[0]))
	fmt.Println(reverse_array(this.Player_2.Positions[1]))
	fmt.Println(this.Player_1.Positions[1])
	fmt.Println(this.Player_1.Positions[0])
}


func PlayersFromName(player_number int, board *Board) (p, p_op *Player) {
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


// Execute instruction method probably needs to be rethought

//func ExecuteInstruction(instruction Instruction, player_number int) (Board, []Instruction) {
//	// Execute the instruction (no decison on whether it is legal or not) and return
//	// Possible following instructions if there are any. Following instructions
//	// Result from captures
//	board := instruction.Board
//	p, _ := PlayersFromName(player_number, &board)
//	current_row := instruction.Row
//	current_column := instruction.Column
//	current_direction := instruction.Direction
//	var next_instructions []Instruction
//
//	num_seeds := p.Positions[current_row][current_column] //How many seeds will be moved
//	p.Positions[current_row][current_column] = 0          // empty the starting pit
//	for i := 0; i < num_seeds; i++ {                      //move the seeds
//		current_row, current_column = next_position(current_row, current_column, current_direction)
//		p.Positions[current_row][current_column] += 1
//	}
//
//	// now for the capturing
//	if capture_possible(board, player_number, current_row, current_column) {
//		board, next_instructions = perform_capture(board, player_number, current_row, current_column)
//	}
//	return board, next_instructions
//}
