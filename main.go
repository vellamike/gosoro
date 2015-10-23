package main

// Initial rules: No reversing, only victory is if other player can't move

// Need the following methods:
// 1. Board initialiser - this is a method which just returns a board.
// 2. Scoring function
// 3. Board mmutator
// 4. A string-based notation for the move.
// 5. A player-perspective on the board, so that confusions regarding clockwise/counterclockwise etc are easily resolved.

// Go Questions:
// Q 1. What is a slice?
//
// Q 2. How do you add a method to a type
// A func (this Type) func_name(func_param param_type) (return_type) {...}
//
// Q 3. What is an interface
// A If a type has all the correct function names, including signatures, then it satisfies an interface
//   And can be passed to another method which takes that interface.
//   Still not sure about the details of this but it sounds quite interesting.

// UPDATE: I've decided that the 4x8 board representation is very unhelpful.
// The 2x2x8 is better, that is to say, the board is represented as two player views.
// Each player view is a 2x8 board.
// For calcuations of scores this can still be mapped to the 4x8 representation if need be,
// But for carrying out mutations this is a much more simple strategy.

// Thoughts on mutators:
// In some positions a decision can be made whether to go clockwise or counterclockwise.
// Some sort of tree is going to be needed to keep track of decisions, OR we could not
// implement this aspect of it for now.

import "fmt"
import "math/rand"

type player struct {
	positions [2][8]int
}

type Board struct {
	player_1 player
	player_2 player
}

func row_sum(intarray [8]int) (int) {
	sum := 0
	for i:= range intarray{
		sum += intarray[i]
	}
	return sum
}


func (this player) naive_score() (score int) {
	score = row_sum(this.positions[0]) + row_sum(this.positions[1]);
	return
}

func random_position(num_seeds int) (player) {
	var p player

	for i:=0; i<num_seeds; i++ {
		row := rand.Intn(2)
		column := rand.Intn(8)
		p.positions[row][column] += 1
	}

	return p
	
}

func random_board(num_seeds int) (Board) {
	var newboard Board;

	newboard.player_1 = random_position(num_seeds)
	newboard.player_2 = random_position(num_seeds)

	return newboard
}


func move(board Board, player_number int, row, column int) (Board) {
	// Move counterclockwise.
	// For now don't capture or re-move
	var p *player
	var p_op *player

	if player_number == 1 {
		p = &board.player_1
		p_op = &board.player_2
	} else {
		p = &board.player_2
		p_op = &board.player_1
	}

	fmt.Println("Initial position of player:")
	fmt.Println(p)

	fmt.Println("Initial position of opponent:")
	fmt.Println(p_op)
		

	num_seeds := p.positions[row][column]

	//Now do the clockwise moving
	fmt.Println("Number of seeds player will move:")
	fmt.Println(num_seeds)

	current_row := row
	current_column := column
	p.positions[row][column] = 0 // empty the starting pit
	for i := 0; i < num_seeds; i++{
		// get the new current row and current column
		switch {
		case current_row == 0 && current_column < 7:
			current_column += 1
		case current_row == 0 && current_column == 7:
			current_row = 1
	        case current_row == 1 && current_column > 0:
			current_column -= 1
	        case current_row == 1 && current_column == 0:
			current_row = 0
		}
		p.positions[current_row][current_column] += 1
	}

	oponent_column := 7 - current_column // oponent's column
	fmt.Println("Opponent column targeted:")
	fmt.Println(oponent_column)
	oponent_row_0_seeds := p_op.positions[0][oponent_column]
	oponent_row_1_seeds := p_op.positions[1][oponent_column]
	if (oponent_row_0_seeds != 0 && oponent_row_1_seeds != 0 && current_row == 1) { // capture occurs
		fmt.Println("Capture!")
		p_op.positions[0][oponent_column] = 0
		p_op.positions[1][oponent_column] = 0
		captured_seeds := oponent_row_0_seeds + oponent_row_1_seeds
		fmt.Println("Captured seeds:")
		fmt.Println(captured_seeds)
		p.positions[current_row][current_column] += captured_seeds
		board = move(board, player_number, current_row, current_column)
	}

	return board
}

func main() {
	fmt.Println("Instantiating a random board")
	newboard := random_board(12)
	moved_board := move(newboard, 1, 1, 2)
	fmt.Println("Final position of player 1:")
	fmt.Println(moved_board.player_1)
	fmt.Println("Final position of opponent:")
	fmt.Println(moved_board.player_2)
}
