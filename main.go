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

// UPDATE: I've decided that the 4x8 board representation is very unhelpful.
// The 2x2x8 is better, that is to say, the board is represented as two player views.
// Each player view is a 2x8 board.
// For calcuations of scores this can still be mapped to the 4x8 representation if need be,
// But for carrying out mutations this is a much more simple strategy.

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

func main() {
	fmt.Println("Instantiating a random board")
	newboard := random_board(202)
}
