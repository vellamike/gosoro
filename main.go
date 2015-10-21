package main

// Need the following methods:
// 1. Board initialiser - this is a method which just returns a board.
// 2. Scoring function
// 3. Board mmutator

// Questions:
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

func main() {
	fmt.Println("Instantiating a uniform board")
	newboard := uniform_board(4)
	fmt.Println(newboard)
	add_seeds(&newboard,1,1,3)
	fmt.Println(newboard)
	fmt.Println(newboard.naive_scorers())
	fmt.Println(newboard.zero_zero())
}
