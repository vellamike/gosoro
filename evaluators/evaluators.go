package evaluators


import "gosoro/ds"
import "math/rand"


type Evaluator struct {

}

func (Evaluator) Score (board ds.Board) float32 {
	// returns score of player 1
	return rand.Float32()
}
