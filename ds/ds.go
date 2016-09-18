package ds

import "fmt"

type BoardInterface interface {
	Display()
}

type Player struct {
	//Represents a player's territory in their frame of reference
	Positions    [2][8]int
	Moved        bool
	LastPosition Coord
}

type Coord struct {
	Row    int
	Column int
}

type Board struct {
	Player_1 Player
	Player_2 Player
	mutator  Mutator
}

func NewBoard(player1, player2 Player) Board {
	return Board{player1, player2, Mutator{}}
}

type Instruction struct {
	Row       int
	Column    int
	Direction string
	Board     Board
}

type Move struct {
	Row    int
	Column int
	Action string
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

func (this Board) ExecuteMove(move Move, player_number int) Board {
	mutator := Mutator{}
	return mutator.ExecuteMove(this, move, player_number)
}
