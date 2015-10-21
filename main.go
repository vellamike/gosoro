package main

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

import "fmt"
import "math/rand"

type Board struct {
	positions [4][8]int
}

func (this Board) zero_zero() (int){
	return this.positions[0][0]
}

func uniform_board(seeds int) (Board) {
	newboard := Board{}
	for i := range newboard.positions {
		for j := range newboard.positions[i] {
			newboard.positions[i][j] = seeds
		}
	}
	return newboard
}

func add_seeds(board *Board, row int, column int, num_seeds int){
	// Add a seed to a specific row and column
	board.positions[row][column] += num_seeds
}

func row_sum(intarray [8]int) (int) {
	sum := 0
	for i:= range intarray{
		sum += intarray[i]
	}
	return sum
}

func (this Board) naive_scorers() (int,int) {

	player_1_score := row_sum(this.positions[0]) + row_sum(this.positions[1]);
	player_2_score := row_sum(this.positions[2]) + row_sum(this.positions[3]);

	return player_1_score, player_2_score
}

func random_board(num_seeds int) (Board) {
	// TODO: remove the repetition from here, I expect there are smarter ways to get
	// a uniformly-distributed number
	var newboard Board;
	// do the first player first
	for i:=0; i<num_seeds; i++ {
		row := rand.Intn(2)
		column := rand.Intn(8)
		newboard.positions[row][column] += 1
	}

	//then the second
	for i:=0; i<num_seeds; i++ {
		row := rand.Intn(2) + 2
		column := rand.Intn(8)
		newboard.positions[row][column] += 1
	}
	return newboard
}

func crazy_mutator(instruction string, board *Board) {
	// Given the instruction, the (in RCD format) this mutator
	// Will move one piece in that direction.
	// For the columnn moved to, whichever player has the highest
	// number of seeds gets to keep their seeds, while all seeds in
	// the opposing column are destroyed.
	
	row := int(instruction)[0]
	return board
}

func main() {
	fmt.Println("Instantiating a random board")
	newboard := random_board(20)
	fmt.Println(newboard)
}
