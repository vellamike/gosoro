package main

// Need the following methods:
// 1. Board initialiser
// 2. Scoring function
// 3. Board mmutator

import "fmt"

type Board [4][8]int;

func uniform_board(seeds int) (Board) {
	newboard := Board{
		{0,0,0,0,0,0,0,0},
		{seeds,seeds,seeds,seeds,seeds,seeds,seeds,seeds},
		{seeds,seeds,seeds,seeds,seeds,seeds,seeds,seeds},
		{0,0,0,0,0,0,0,0}}
	return newboard
}


func main() {
	fmt.Println("Instantiating a uniform board")
	newboard := uniform_board(4)
	fmt.Println(newboard)
}
