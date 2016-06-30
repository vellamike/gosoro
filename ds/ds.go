package ds

import "fmt"

type Player struct {
	//Represents a player's territory in their frame of reference
	Positions [2][8]int
}

type Coord struct {
	row    int
	column int
	board  *Board
	player *Player
}

type Board struct {
	player_1 Player
	player_2 Player
}

type Instruction struct {
	row       int
	column    int
	direction string
	board     Board
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
	fmt.Println(reverse_array(this.player_2.Positions[0]))
	fmt.Println(reverse_array(this.player_2.Positions[1]))
	fmt.Println(this.player_1.Positions[1])
	fmt.Println(this.player_1.Positions[0])
}
