package gamecontrollers

import "gosoro/ds"
import "gosoro/ai"
import "gosoro/mutators"

type gamecontroller struct {
	board   ds.BoardInterface
	ai      ai.AI
	mutator mutators.Mutator
}

func NewGameController(generator func() ds.Board, ai ai.AI, mutator mutators.Mutator) *gamecontroller {
	board := generator()
	b := gamecontroller{board, ai, mutator}
	return &b

}

func (gc gamecontroller) DisplayBoard() {
	gc.board.Display()
}
