package rulesets

import "gosoro/ds"
import "fmt"

type RuleSet struct {
}

func (RuleSet) AvailableMoves (board ds.Board, player_number int) []ds.Move{

	var player ds.Player
	
	if player_number == 2 {
		player = board.Player_2
	} else {
		player = board.Player_1
	}
	
	var moves []ds.Move

	// outer loop: row
	for r := 0; r < 2; r++ {
		for c := 0; c<8; c++ {
			if player.Positions[r][c] > 1 {
				moves = append(moves, ds.Move{r, c, "C"})
			}
		}

	}
	fmt.Println("All available moves to computer:")
	fmt.Println(moves)
	return moves
}
