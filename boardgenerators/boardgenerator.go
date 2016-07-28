package boardgenerators

import "gosoro/utils"
import "gosoro/ds"

func Randomboard() ds.Board {
	return utils.Random_board(32)
}
