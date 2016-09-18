package utils

import "math/rand"
import "time"
import "gosoro/ds"

func random_position(num_seeds int) ds.Player {
	//choose a random pit
	var p ds.Player

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < num_seeds; i++ {
		row := rand.Intn(2)
		column := rand.Intn(8)
		p.Positions[row][column] += 1
	}

	return p

}

func Random_board(num_seeds int) ds.Board {
	//Initialize a random board
	player1 := random_position(num_seeds)
	player2 := random_position(num_seeds)

	return ds.NewBoard(player1, player2)
}
